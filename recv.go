package main

import (
	"github.com/streadway/amqp"
	"log"
)

func errorFaic(err error, msg string)  {
	if err != nil{
		log.Fatal("%s: %s", msg, err)
	}
}

func main(){
	conn, err := amqp.Dial("amqp://admin:123456@127.0.0.1:5672")
	errorFaic(err, "amqp connect error")
	defer conn.Close()

	channel, err := conn.Channel()
	errorFaic(err, "channel open error")
	defer channel.Close()

	q, err := channel.QueueDeclare("python-test",
		true,
		false,
		false,
		false,
		nil,
	)
	errorFaic(err, "queue declare error")

	msgs, err := channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
		)
	errorFaic(err, "queue consume error")

	forever := make(chan bool)
	go func() {
		for d := range msgs{
			log.Printf("this id %s", string(d.Body))
		}
	}()

	<- forever
}