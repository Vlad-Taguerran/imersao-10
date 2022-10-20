package kafka

import (
	"encoding/json"
	route2 "github.com/Vladmir-Taguerran/simulatorGo/application/route"
	"github.com/Vladmir-Taguerran/simulatorGo/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
	"time"
)

// Produce is responsible to publish the positions of each request
// Example of a json request:
// {"clientId":"1","routeId":"1"}
// {"clientId":"2","routeId":"2"}
// {"clientId":"3","routeId":"3"}
func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route2.NewRoute()
	err := json.Unmarshal(msg.Value, &route)
	if err != nil {
		return
	}
	err = route.LoadPositions()
	if err != nil {
		return
	}
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}
	for _, p := range positions {
		err := kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		if err != nil {
			return
		}
		time.Sleep(time.Millisecond * 500)
	}
}
