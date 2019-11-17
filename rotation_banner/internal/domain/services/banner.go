/*
	Данный сервис реализует интерфейс Service.
	Здесь происходит передача данных на уровень баз данных, а так же логгирование событий.
*/
package services

import (
	"context"

	"github.com/egor1344/banner/rotation_banner/internal/domain/interfaces"
	"go.uber.org/zap"
)

type Banner struct {
	Database interfaces.Database // Интерфейс для работы с БД
	Log      *zap.SugaredLogger  // Логгер
}

// AddBanner - Добавить баннер
func (b *Banner) AddBanner(ctx context.Context, idBanner int64, idSlot int64) error {
	b.Log.Info("add banner")
	err := b.Database.AddBanner(ctx, idBanner, idSlot)
	return err
}

// DelBanner - Удалить баннер
func (b *Banner) DelBanner(ctx context.Context, idBanner int64) error {
	b.Log.Info("Del Banner")
	err := b.Database.DelBanner(ctx, idBanner)
	return err
}

// CountTransition - Засчитать переход
func (b *Banner) CountTransition(ctx context.Context, idBanner int64, idSocDemGroup int64) error {
	b.Log.Info("Count Transition")
	err := b.Database.CountTransition(ctx, idBanner, idSocDemGroup)
	return err
}

// GetBanner - Выбрать баннер для показа
func (b *Banner) GetBanner(ctx context.Context, idSlot int64, idSocDemGroup int64) (int64, error) {
	b.Log.Info("Get Banner")
	idBanner, err := b.Database.GetBanner(ctx, idSlot, idSocDemGroup)
	return idBanner, err
}
