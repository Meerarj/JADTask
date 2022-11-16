package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type EnvVars struct {
	number1 int `envconfig:"NUMBER_ONE" default:"1"`
	Number2 int `envconfig:"NUMBER_TWO" required:"true"`
}
type myAppEnvVars struct {
	MyNumber1 int    `envconfig:"MY_NUMBER_ONE" default:"1"`
	MyNumber2 int    `envconfig:"MY_NUMBER_TWO" required:"true"`
	User      string `envconfig:"USER"`
}
// type KafkaConfig struct {
	Address    string `envconfig:"KAFKA_ADDRESS"`
	Port       int    `envconfig:"KAFKA_PORT" default:"9092"`
	WriteTopic string `envconfig:"KAFKA_WRITE_TOPIC" default:"topic1"`
	ReadTopic  string `split_words:"true"`
}

func main() {
	// e := EnvVars{}
	env := &EnvVars{}
	var tEnv *EnvVars
	fmt.Printf("tEnv: %v", tEnv)
	var pEnv *EnvVars = env
	// os.Setenv("KAFKA_WRITE_TOPIC", "write-topic")
	err := envconfig.Process("", pEnv)
	if err != nil {
		fmt.Printf("env not found: %v", err)
	}
	fmt.Printf("Num1: %d, Num2: %d\n", env.number1, env.Number2)
	printEnvVars(pEnv)
	printEnvVars(env)
	myEnv := &myAppEnvVars{}
	err = envconfig.Process("my", myEnv)
	if err != nil {
		fmt.Printf("env not found: %v", err)
	}

	fmt.Printf("Num1: %d, Num2: %d, user: %s\n", myEnv.MyNumber1, myEnv.MyNumber2, myEnv.User)
	kafkaConfig := &KafkaConfig{}
	err = envconfig.Process("KAFKA", kafkaConfig)
	if err != nil {
		fmt.Printf("kafka config parameter error: %v", err)
	}
	fmt.Printf("kafka config: %+v\n", kafkaConfig)
}
func printEnvVars(e *EnvVars) {
	fmt.Printf("e: %v\n", e)
}
