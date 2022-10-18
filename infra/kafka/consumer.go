package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

type KafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

func (k *KafkaConsumer) consume() {
	//configMap := &ckafka.ConfigMap{}
}
