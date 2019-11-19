package interfaces

import "context"

// Database - Интерфейс реализирующий работу с БД
type Database interface {
	// Добавить баннер
	AddBanner(ctx context.Context, idBanner int64, idSlot int64) error
	// Удалить баннер
	DelBanner(ctx context.Context, idBanner int64) error
	// Засчитать переход
	CountTransition(ctx context.Context, idBanner int64, idSocDemGroup int64, idSlot int64) error
	// Выбрать баннер для показа
	GetBanner(ctx context.Context, idSlot int64, idSocDemGroup int64) (int64, error)
}
