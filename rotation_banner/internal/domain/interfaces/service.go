package interfaces

import (
	"context"
)

// Service - Интерфейс реализирующий "прослойку" между API(в данном случае GRPC, REST API) и БД
type Service interface {
	// Добавить баннер
	AddBanner(ctx context.Context, idBanner, idSlot int64) error
	// Удалить баннер
	DelBanner(ctx context.Context, idBanner int64) error
	// Засчитать переход
	CountTransition(ctx context.Context, idBanner, idSocDemGroup, idSlot int64) error
	// Выбрать баннер для показа
	GetBanner(ctx context.Context, idSlot, idSocDemGroup int64) (int64, error)
	// Закрытие открытых соединений при сбое (БД, AMQP прочее)
	CloseConnection()
}
