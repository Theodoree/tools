package main

import (
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		os.Exit(0)
	}
}

func dial() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		conn.Close()
	}
	failOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

func send() {

	conn := dial()
	defer conn.Close()

	/*
	   -----------------------------------
	   | type|channelId|payLoad|frame-end|
	   |---------------------------------|
	    通信协议   1byte 类型  2byte channelId  size 4byte  payload  frame-end
	    每个client都会有一个位图,来维护已使用的channelId
	*/
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	/*
	   整体拼接逻辑抽象为三个层面
	   消息本身的编码(消息本身的编码)
	   Frame类型的编码(对payload附加一些参数)
	   最终的消息组装成符合通信协议的
	*/

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   //   持久化
		false,   // 未使用删除
		false,   // 独占
		false,   // 异步、同步调用
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	/*组装二进制协议发送*/

	for i := 0; i < 10; i++ {
		body := "Hello World!" + strconv.Itoa(i)
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		log.Printf(" [x] Sent %s", body)
		failOnError(err, "Failed to publish a message")
	}

}
func recv(sleep time.Duration, consumeId uint64) {

	conn := dial()
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	/* 确保queue存在 */
	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	/*
	   该函数,会将消费者加入chanMap,再conn的循环read中,会一层一层的分发给这些channel下面的consumer。
	*/

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // 消费者tag
		false,  // 自动Ack
		false,  // 独占
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	ch.Qos(1, 0, false)
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s 消费者 %d", d.Body, consumeId)
			time.Sleep(sleep)
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

// Work Queues
func newTask() {

	conn := dial()
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // 持久化
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := bodyFrom([]string{"hello", "world"})
	/*组装二进制协议发送*/

	for i := 0; i < 20; i++ {
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(body + " " + strconv.Itoa(i)),
			})
		failOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s", body)
	}

}
func worker(sleep time.Duration, consumeId uint64, nack bool) {
	conn := dial()
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	/*声明队列*/

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	/*
	   实际上是设置单批次发送数量(byte),减少RTT的支出
	   预计获取数量
	   预计获取byte
	*/
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // 如果为true,那么会应用于当前连接的所有channel和未来的channel
	)
	failOnError(err, "Failed to set QoS")

	/*
	   初始化订阅的消费者
	*/
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			if nack {
				log.Printf("消费者 %d Received a message: %s  消费失败", consumeId, d.Body)
				time.Sleep(sleep)
				d.Nack(false, true)
				nack = false
			} else {
				log.Printf("消费者 %d Received a message: %s", consumeId, d.Body)
				time.Sleep(sleep)
				d.Ack(false)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

/* Publish/Subscribe
发布订阅模式
*/
func receiveLogs() {
	conn := dial()
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// exchange声明
	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // 持久
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	// queue声明
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // 独占
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// 将queue绑定到exchange
	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"logs", // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	// 声明消费者
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
func emitLog() {
	conn := dial()
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// exchange更像topic,分发给queue
	// 声明exchange
	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	body := bodyFrom([]string{"hello", "world"})
	err = ch.Publish(
		"logs", // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)

}

/* Routing
路由模式
*/
func receiveLogsDirect() {
	conn := dial()
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// exchange声明
	err = ch.ExchangeDeclare(
		"logs_direct", // name
		"direct",      // type
		true,          // 持久化
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	// queue声明
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // 独占
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	for _, s := range []string{"info", "warn"} {
		log.Printf("Binding queue %s to exchange %s with routing key %s", q.Name, "logs_direct", s)
		// 绑定exchange和queue_routing key
		err = ch.QueueBind(
			q.Name,        // queue name
			s,             // routing key
			"logs_direct", // exchange
			false,
			nil)
		failOnError(err, "Failed to bind a queue")
	}

	// 初始化消费者
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
func emitLogDirect() {
	conn := dial()
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_direct", // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	body := bodyFrom([]string{"warn", "a warning"})
	err = ch.Publish(
		"logs_direct", // exchange
		severityFrom([]string{"warn", "a warning"}), // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}

// Topics
func receiveLogsTopics() {
	conn := dial()
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// exchange初始化
	err = ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	// queue初始化
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	for _, s := range []string{"kern.*", "*.critical"} {
		log.Printf("Binding queue %s to exchange %s with routing key %s", q.Name, "logs_topic", s)
		// 绑定routing key
		err = ch.QueueBind(
			q.Name,       // queue name
			s,            // routing key
			"logs_topic", // exchange
			false,
			nil)
		failOnError(err, "Failed to bind a queue")
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
func emitLogTopics() {
	conn := dial()
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	body := bodyFrom([]string{"kern.critical", "A critical kernel error"})
	err = ch.Publish(
		"logs_topic", // exchange
		severityFrom([]string{"kern.critical", "A critical kernel error"}), // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}

// RPC
func rpcServer() {
	conn := dial()
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rpc_queue", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	//round trip time
	// 优化RTT
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	// 声明消费者
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			n, err := strconv.Atoi(string(d.Body))
			failOnError(err, "Failed to convert body to integer")

			log.Printf(" [.] fib(%d)", n)
			response := fib(n)

			// 写回
			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(strconv.Itoa(response)),
				})
			failOnError(err, "Failed to publish a message")

			// 已消费
			if err != nil {
				d.Nack(false, true)
			} else {
				d.Ack(false)
			}
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}
func rpcClient(n int) {
	rand.Seed(time.Now().UTC().UnixNano())

	log.Printf(" [x] Requesting fib(%d)", n)
	res, err := fibonacciRPC(n)
	failOnError(err, "Failed to handle RPC request")

	log.Printf(" [.] Got %d", res)
	os.Exit(0)
}
func fibonacciRPC(n int) (res int, err error) {
	conn := dial()
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 初始化一个未命名queue(其实就是随机名)
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// 初始化消费者,绑定queue
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	corrId := randomString(32)

	// 往rpc_queue发送一个rpc调用请求。
	err = ch.Publish(
		"",          // exchange
		"rpc_queue", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          []byte(strconv.Itoa(n)),
		})
	failOnError(err, "Failed to publish a message")

	// 遍历queue,等待调用方消费者写回
	for d := range msgs {
		if corrId == d.CorrelationId {
			res, err = strconv.Atoi(string(d.Body))
			failOnError(err, "Failed to convert body to integer")
			d.Ack(false)
			break
		} else {
			d.Nack(false, true)
		}
	}

	return
}

func bodyFrom(args []string) string {
	if len(args) == 0 {
		return "hello"
	}
	return strings.Join(args, " ")
}
func severityFrom(args []string) string {
	if len(args) == 0 {
		return "info"
	}
	return args[0]
}
func randomString(l int) string {
	buf := make([]byte, l)
	for i := 0; i < l; i++ {
		buf[i] = byte(randInt(65, 90))
	}
	return string(buf)
}
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
func fib(n int) int {
	var f = make([]int, n+1)

	switch n {
	case 0, 1:
		return 1
	}

	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
