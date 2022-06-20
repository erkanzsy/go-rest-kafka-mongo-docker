package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gopkg.in/mgo.v2"
	"time"
)

const (
	hosts      = "localhost:27017"
	database   = "db"
	username   = ""
	password   = ""
	collection = "flights"
)

type Flight struct {
	From    string `json:"from"`
	To 	   	string `json:"to"`
	Airline string `json:"airline"`
	Price   string `json:"price"`
}

type MongoStore struct {
	session *mgo.Session
}

var mongoStore = MongoStore{}

func main() {

	//Create MongoDB session
	session := initialiseMongo()
	mongoStore.session = session

	receiveFromKafka()

}

func initialiseMongo() (session *mgo.Session) {

	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	return

}

func receiveFromKafka() {

	fmt.Println("Start receiving from Kafka")
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "group-id-1",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"flights-topic1"}, nil)

	for {
		msg, err := c.ReadMessage(-1)

		if err == nil {
			fmt.Printf("Received from Kafka %s: %s\n", msg.TopicPartition, string(msg.Value))
			flight := string(msg.Value)
			saveFlightToMongo(flight)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	c.Close()

}

func saveFlightToMongo(flightString string) {
	fmt.Println("saveFlightToMongo")

	col := mongoStore.session.DB(database).C(collection)

	var _flight Flight
	b := []byte(flightString)
	err := json.Unmarshal(b, &_flight)
	if err != nil {
		panic(err)
	}

	errMongo := col.Insert(_flight)
	if errMongo != nil {
		panic(errMongo)
	}

	fmt.Printf("Saved to MongoDB : %s\n", flightString)
}