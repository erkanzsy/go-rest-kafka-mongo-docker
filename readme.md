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
