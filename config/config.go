package config

type KafkaConsumer struct {
	KafkaBrokerUrl      string `mapstructure:KAFKA_BROKER_URL`
	KafkaBrokerUserName string `mapstructure:KAFKA_BROKER_USER_NAME`
	KafkaBrokerPassword string `mapstructure:KAFKA_BROKER_PASSWORD`
}
