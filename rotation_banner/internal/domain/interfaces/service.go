package interfaces

import (
	"context"
)

// Service - Интерфейс реализирующий "прослойку" между API(в данном случае GRPC, REST API) и БД
type Service interface {
	// Добавить баннер
	AddBanner(ctx context.Context, idBanner int64, idSlot int64) error
	// Удалить баннер
	DelBanner(ctx context.Context, idBanner int64) error
	// Засчитать переход
	CountTransition(ctx context.Context, idBanner int64, idSocDemGroup int64) error
	// Выбрать баннер для показа
	GetBanner(ctx context.Context, idSlot int64, idSocDemGroup int64) (int64, error)
}
