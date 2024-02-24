package main

import (
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 连接到RabbitMQ服务器
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 创建一个通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明一个队列
	q, err := ch.QueueDeclare(
		"hello", // 队列名称
		false,   // 是否持久化
		false,   // 是否自动删除
		false,   // 是否排他性
		false,   // 是否阻塞
		nil,     // 其他属性
	)
	failOnError(err, "Failed to declare a queue")

	// 接收消息
	msgs, err := ch.Consume(
		q.Name, // 队列名称
		"",     // 消费者标签
		true,   // 是否自动应答
		false,  // 是否独占
		false,  // 是否阻塞
		false,  // 额外属性
		nil,    // 其他参数
	)
	failOnError(err, "Failed to register a consumer")

	// 处理接收到的消息
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Println("Waiting for messages...")
	<-forever
}
