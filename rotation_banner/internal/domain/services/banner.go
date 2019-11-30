package services

/*
	Данный сервис реализует интерфейс Service.
	Здесь происходит передача данных на уровень баз данных, а так же логгирование событий.
*/
import (
	"context"

	"github.com/egor1344/banner/rotation_banner/pkg/metrics"

	"github.com/egor1344/banner/rotation_banner/internal/domain/interfaces"
	"go.uber.org/zap"
)

// Banner - структура которая опредляет внутренние актрибуты сервиса
type Banner struct {
	Database interfaces.Database // Интерфейс для работы с БД
	AMQP     interfaces.AMQP     // Очередь сообщений
	Log      *zap.SugaredLogger  // Логгер
}

// AddBanner - Добавить баннер
func (b *Banner) AddBanner(ctx context.Context, idBanner int64, idSlot int64) error {
	b.Log.Info("add banner")
	metrics.AddBannerCounter.Inc()
	err := b.Database.AddBanner(ctx, idBanner, idSlot)
	if err != nil {
		return err
	}
	return nil
}

// DelBanner - Удалить баннер
func (b *Banner) DelBanner(ctx context.Context, idBanner int64) error {
	b.Log.Info("Del Banner")
	metrics.DelBannerCounter.Inc()
	err := b.Database.DelBanner(ctx, idBanner)
	if err != nil {
		return err
	}
	return nil
}

// CountTransition - Засчитать переход
func (b *Banner) CountTransition(ctx context.Context, idBanner, idSocDemGroup, idSlot int64) error {
	b.Log.Info("Count Transition")
	metrics.CountTransitionCounter.Inc()
	err := b.Database.CountTransition(ctx, idBanner, idSocDemGroup, idSlot)
	if err != nil {
		return err
	}
	err = b.AMQP.AddEvent(ctx, "count_transition", idBanner, idSocDemGroup, idSlot)
	if err != nil {
		return err
	}
	return nil
}

// GetBanner - Выбрать баннер для показа
func (b *Banner) GetBanner(ctx context.Context, idSlot int64, idSocDemGroup int64) (int64, error) {
	b.Log.Info("Get Banner")
	metrics.GetBannerCounter.Inc()
	idBanner, err := b.Database.GetBanner(ctx, idSlot, idSocDemGroup)
	if err != nil {
		return 0, err
	}
	err = b.AMQP.AddEvent(ctx, "get_banner", idBanner, idSocDemGroup, idSlot)
	if err != nil {
		return 0, err
	}
	return idBanner, nil
}

// CloseConnection - Закрытие открытых коннектов
func (b *Banner) CloseConnection() {
	b.Log.Info("CloseConnection")
	b.AMQP.Close()
	b.Database.Close()
}
