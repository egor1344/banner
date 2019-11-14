package databases

import (
	"context"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// PgBannerStorage - реализует работу с БД. Реализует интерфейс database
type PgBannerStorage struct {
	db  *sqlx.DB
	Log *zap.SugaredLogger
}

// AddBanner - Добавить баннер
func (pgbs *PgBannerStorage) AddBanner(ctx context.Context, idBanner int64, idSlot int64) (string, error) {
	pgbs.Log.Info("bd add banner")
	return "", nil
}

// DelBanner - Удалить баннер
func (pgbs *PgBannerStorage) DelBanner(ctx context.Context, idBanner int64) error {
	pgbs.Log.Info("bd del banner")
	return nil
}

// CountTransition - Засчитать переход
func (pgbs *PgBannerStorage) CountTransition(ctx context.Context, idBanner int64, idSocDemGroup int64) error {
	pgbs.Log.Info("bd count transition")
	return nil
}

// GetBanner - Выбрать баннер для показа
func (pgbs *PgBannerStorage) GetBanner(ctx context.Context, idSlot int64, idSocDemGroup int64) (int64, error) {
	pgbs.Log.Info("bd get banner")
	return 0, nil
}
