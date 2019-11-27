package amqp

import (
	"context"
	"encoding/json"

	"go.uber.org/zap"

	"github.com/streadway/amqp"
)

type Event struct {
	typeEvent                       string
	idBanner, idSocDemGroup, idSlot int64
}

type Rabbit struct {
	Log                *zap.SugaredLogger // Логгер
	connection         *amqp.Connection
	channel            *amqp.Channel
	QueueName, AMQPDSN string
}

// Init - инициализация подключения к rabbitmq
func (r *Rabbit) Init() error {
	conn, err := amqp.Dial(r.AMQPDSN)
	if err != nil {
		r.Log.Error(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		r.Log.Error(err)
	}
	_, err = ch.QueueDeclare(
		r.QueueName, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		r.Log.Error(err)
	}
	r.channel = ch
	r.connection = conn
	return nil
}

// Close - инициализация подключения к rabbitmq
func (r *Rabbit) Close() {
	r.Log.Info("Closing connection")
	r.channel.Close()
	r.connection.Close()
}

func (r *Rabbit) AddEvent(ctx context.Context, typeEvent string, idBanner, idSocDemGroup, idSlot int64) error {
	body := &Event{typeEvent: typeEvent, idSocDemGroup: idSocDemGroup, idSlot: idSlot, idBanner: idBanner}
	b, err := json.Marshal(body)
	if err != nil {
		r.Log.Error(err)
	}
	r.Log.Info(b, string(b))
	err = r.channel.Publish("", r.QueueName, false, false, amqp.Publishing{ContentType: "json", Body: b})
	if err != nil {
		r.Log.Error(err)
	}
	return nil
}
