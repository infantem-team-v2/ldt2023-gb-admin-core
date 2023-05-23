package dconfig

type BrokerConfig struct {
	RabbitMQ struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
	} `json:"RabbitMQ"`
	Kafka struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"Kafka"`
}
