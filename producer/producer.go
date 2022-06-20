package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Flight struct {
	From    string `json:"from"`
	To 	   	string `json:"to"`
	Airline string `json:"airline"`
	Price   string `json:"price"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/flight-history", FlightsPostHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":9090", router))
}

func FlightsPostHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}

	var _flight Flight
	err = json.Unmarshal(b, &_flight)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	saveFlightToKafka(_flight)

	jsonString, err := json.Marshal(_flight)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("content-type", "application/json")

	w.Write(jsonString)
}

func saveFlightToKafka(Flight Flight) {
	fmt.Println("saveFlightToKafka")

	jsonString, err := json.Marshal(Flight)

	jobString := string(jsonString)

	fmt.Println(jobString)

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}

	topic := "flights-topic1"
	for _, word := range []string{string(jobString)} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}
}