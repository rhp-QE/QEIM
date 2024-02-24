package reqnet

import (
	"github.com/streadway/amqp"
	"log"
)

type reqnetImpl struct {
	receiveMessageCallBack func(uid uint64, msg []byte)
	mqAddr                 string

	sendConn    *amqp.Connection
	receiveConn *amqp.Connection
}

// / 发送消息
func (reqnet *reqnetImpl) SendMessage(uid uint64, msg []byte) {
	// 创建一个通道
	ch, err := reqnet.sendConn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明一个队列
	q, err := ch.QueueDeclare(
		string(uid), // 队列名称
		false,       // 是否持久化
		false,       // 是否自动删除
		false,       // 是否排他性
		false,       // 是否阻塞
		nil,         // 其他属性
	)
	failOnError(err, "Failed to declare a queue")

	// 发送消息到队列
	body := "Hello RabbitMQ!"
	err = ch.Publish(
		"",     // 交换机名称
		q.Name, // 队列名称
		false,  // 是否强制
		false,  // 是否立即
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	log.Println("Message sent:", body)
}

// / 注册接受消息回调
func (reqnet *reqnetImpl) InjectReceiveMessageHandle(callBack func(uid uint64, msg []byte)) {
	// 创建一个通道
	ch, err := reqnet.receiveConn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 接收消息
	msgs, err := ch.Consume(
		string(uid), // 队列名称
		"",          // 消费者标签
		true,        // 是否自动应答
		false,       // 是否独占
		false,       // 是否阻塞
		false,       // 额外属性
		nil,         // 其他参数
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

func (reqnet *reqnetImpl) Start() {
	var err error
	reqnet.sendConn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "sendconnection error")

	reqnet.receiveConn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "receiveConn error")
}

func (reqnet *reqnetImpl) Stop() {
	reqnet.receiveConn.Close()
	reqnet.receiveConn.Close()
}

//#pramark private

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
