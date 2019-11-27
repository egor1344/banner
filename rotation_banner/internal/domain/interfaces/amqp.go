package interfaces

import "context"

// AMQP - Интерфейс реализирующий работу с очередями сообщений
type AMQP interface {
	// Добавить событие
	AddEvent(ctx context.Context, typeEvent string, idBanner, idSocDemGroup, idSlot int64) error
	Close()
}
