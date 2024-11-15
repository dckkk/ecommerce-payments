package main

import (
	"ecommerce_payments/cmd"
	"ecommerce_payments/helpers"
)

func main() {
	// load config
	helpers.SetupConfig()

	// load log
	helpers.SetupLogger()

	// load db
	helpers.SetupPostgreSQL()

	// load redis
	// helpers.SetupRedis()

	// run kafka consumer
	go cmd.ServeKafkaConsumerPaymentInit()
	go cmd.ServeKafkaConsumerRefund()

	// run http
	cmd.ServeHTTP()
}
