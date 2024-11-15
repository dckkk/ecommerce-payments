package cmd

import (
	"ecommerce_payments/helpers"
	"fmt"
	"strconv"
	"strings"

	"github.com/IBM/sarama"
)

func ServeKafkaConsumerPaymentInit() {
	brokers := strings.Split(helpers.GetEnv("KAFKA_BROKERS", "localhost:9092"), ",")
	topic := helpers.GetEnv("KAFKA_TOPIC_PAYMENT_INITIATE", "example-topic")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		helpers.Logger.Error("Failed to connect with kafka as consumer", err)
		return
	}

	partitionNumberStr := helpers.GetEnv("KAFKA_TOPIC_PAYMENT_INITIATE_PARTITION", "3")
	partitionNumber, _ := strconv.Atoi(partitionNumberStr)
	for i := int32(0); i < int32(partitionNumber); i++ {
		go func() {
			partitionConsumer, err := consumer.ConsumePartition(topic, i, sarama.OffsetOldest)
			if err != nil {
				helpers.Logger.Error("Failed to create consume partition 0", err)
				return
			}

			for msg := range partitionConsumer.Messages() {
				fmt.Printf("Received message: %s from partition %d\n", string(msg.Value), msg.Partition)
			}
		}()
	}

}

func ServeKafkaConsumerRefund() {
	brokers := strings.Split(helpers.GetEnv("KAFKA_BROKERS", "localhost:9092"), ",")
	topic := helpers.GetEnv("KAFKA_TOPIC_REFUND", "example-topic")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		helpers.Logger.Error("Failed to connect with kafka as consumer", err)
		return
	}

	partitionNumberStr := helpers.GetEnv("KAFKA_TOPIC_REFUND_PARTITION", "3")
	partitionNumber, _ := strconv.Atoi(partitionNumberStr)
	for i := int32(0); i < int32(partitionNumber); i++ {
		go func() {
			partitionConsumer, err := consumer.ConsumePartition(topic, i, sarama.OffsetOldest)
			if err != nil {
				helpers.Logger.Error("Failed to create consume partition 0", err)
				return
			}

			for msg := range partitionConsumer.Messages() {
				fmt.Printf("Received message: %s from partition %d\n", string(msg.Value), msg.Partition)
			}
		}()
	}
}
