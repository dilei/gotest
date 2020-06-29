package main

// test 1
// 2个Producer 1个Consumer
// produce1() 发布publish "x","y" 到 topic "test"
// produce2() 发布publish "z" 到 topic "test"
// consumer1() 订阅subscribe channel "sensor01" of topic "test"

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

type myMessageHandler struct{}

// HandleMessage implements the Handler interface.
func (h *myMessageHandler) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
	return nil
}

func main() {
	var role = flag.String("role", "", "producer | consumer")
	flag.Parse()

	if *role == "producer" {
		produce1()
	} else {
		consumer1()
		for {
			time.Sleep(time.Second * 10)
		}
	}
}

// go-nsq包做客户端，使用tcp进行message的收发
func produce1() {
	// Instantiate a producer.
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("10.255.209.167:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	messageBody := []byte("hello")
	topicName := "test"

	// Synchronously publish a single message to the specified topic.
	// Messages can also be sent asynchronously and/or in batches.
	err = producer.Publish(topicName, messageBody)
	if err != nil {
		log.Fatal(err)
	}

	// Gracefully stop the producer.
	producer.Stop()
}

func consumer1() {
	// Instantiate a consumer that will subscribe to the provided channel.
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test", "ch01", config)
	if err != nil {
		log.Fatal(err)
	}

	// Set the Handler for messages received by this Consumer. Can be called multiple times.
	// See also AddConcurrentHandlers.
	consumer.AddHandler(&myMessageHandler{})

	// c.ConnectToNSQD(addrs) // 直连nsqd方式

	// Use nsqlookupd to discover nsqd instances.
	// See also ConnectToNSQD, ConnectToNSQDs, ConnectToNSQLookupds.
	err = consumer.ConnectToNSQLookupd("10.255.209.167:4161")
	if err != nil {
		log.Fatal(err)
	}

	// Gracefully stop the consumer.
	// consumer.Stop()
}
