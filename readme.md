# go-rest-kafka-mongo-docker

### This project is about:

- Golang
- Kafka
- Mux
- MongoDB

## Run this project

`make setup`

`make produce`

````
curl --request POST \
  --url http://localhost:9090/flight-history \
  --header 'Content-Type: application/json' \
  --data '{
	"from": "IST",
	"to": "AMS",
	"airline": "PC",
	"price": "3000"
}'
````
`make consume`


Check RockMongo [here](http://localhost:8080/ "rockmongo link")


## Screen

![Screen Shot 2022-06-20 at 22 10 48](https://user-images.githubusercontent.com/22520257/174665102-44c5bcc5-503a-4b77-94e2-02813c4fd994.png)