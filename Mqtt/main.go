package main

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func main() {
	var broker = "mqtt.antares.id"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	// opts.SetUsername("emqx")
	// opts.SetPassword("public")
	opts.SetKeepAlive(60 * 2 * time.Second)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for true {
		sub(client)
		time.Sleep(time.Second * 5)
	}
	//sub(client)
	//publish(client)

	client.Disconnect(250)
}

// func publish(client mqtt.Client) {
// 	num := 10
// 	for i := 0; i < num; i++ {
// 		text := fmt.Sprintf("Message %d", i)
// 		token := client.Publish("topic/test", 0, false, text)
// 		token.Wait()
// 		time.Sleep(time.Second)
// 	}
// }

func sub(client mqtt.Client) {
	topic := "/oneM2M/resp/antares-cse/0b6b7781f24344d5:b0ec30902bc42bdf/json"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	//fmt.Printf("Subscribed to topic: %s\n", topic)
}
