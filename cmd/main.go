package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Abhi-Harsha/kakfka-consumer/config"
	"github.com/Abhi-Harsha/kakfka-consumer/consumer"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("consumer")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AutomaticEnv()
	viper.BindEnv("KafkaBrokerUrl", "KAFKA_BROKER_URL")
	viper.BindEnv("KafkaBrokerUserName", "KAFKA_BROKER_USER_NAME")
	viper.BindEnv("KafkaBrokerPassword", "KAFKA_BROKER_PASSWORD")
	var config config.KafkaConsumer
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("conf not found")
		} else {
			// Config file was found but another error was produced
		}
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("error unmarshalling: ", err)
	}

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.KafkaBrokerUrl,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
		"security.protocol": "SASL_SSL",
		"sasl.mechanism":    "PLAIN",
		"sasl.username":     config.KafkaBrokerUserName,
		"sasl.password":     config.KafkaBrokerPassword,
	})
	fmt.Println("connecting to consumer...")
	if err != nil {
		panic(err)
	}
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go consumer.Read(c)

	<-interrupt
	fmt.Println("interrupt!!!")
	c.Close()
}
