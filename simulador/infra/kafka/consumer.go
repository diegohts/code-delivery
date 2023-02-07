package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaConsumer contém toda a lógica e configurações do consumidor das conexões do Apache Kafka/
// Também possui um canal de mensagens que é um canal onde as mensagens serão enviadas
type KafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

// NewKafkaConsumer cria uma nova estrutura KafkaConsumer com seu canal de mensagem como dependência
func NewKafkaConsumer(msgChan chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MsgChan: msgChan,
	}
}

// Consumir consome todas as mensagens extraídas do apache kafka e enviadas para o canal de mensagens
func (k *KafkaConsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}
	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("Erro ao consumir mensagem kafka:" + err.Error())
	}
	topics := []string{os.Getenv("KafkaReadTopic")}
	c.SubscribeTopics(topics, nil)
	fmt.Println("O consumidor Kafka foi iniciado")
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			k.MsgChan <- msg
		}
	}
}
