package main

import (
	"fmt"
	"github.com/streadway/amqp"
)
import "log"

func errorFaic(err error, msg string)  {
	if err != nil{
		log.Fatal("%s: %s", msg, err)
	}
}

func main()  {
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

	var list = [5]string{"1","2","3","4","5,9"}
	for _, v := range list {
		var msg = fmt.Sprintf("this is go message %s", v)
		err = channel.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body: []byte(msg),
			})
	}

}

