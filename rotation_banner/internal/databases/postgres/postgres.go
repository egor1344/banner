package postgres

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

// InitPgBannerStorage - Инициализация соединения с БД
func InitPgBannerStorage(dsn string) (*PgBannerStorage, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PgBannerStorage{db: db}, nil
}

// existsBanner - Проверка на существование баннера
func (pgbs *PgBannerStorage) existsBanner(ctx context.Context, idBanner int64, create bool) (bool, error) {
	pgbs.Log.Info("exists banner")
	var count int64
	err := pgbs.db.GetContext(ctx, &count, "select count(*) from banners where id=$1;", idBanner)
	if err != nil {
		pgbs.Log.Error("databases error ", err)
	}
	if count == 0 {
		if create {
			_, err = pgbs.db.ExecContext(ctx, "INSERT into banners values ($1);", idBanner)
			if err != nil {
				pgbs.Log.Error(err)
			}
			return true, nil
		} else {
			return false, nil
		}
	}
	return true, nil
}

// existsSlot - Проверка на существование баннера
func (pgbs *PgBannerStorage) existsSlot(ctx context.Context, idSlot int64, create bool) (bool, error) {
	pgbs.Log.Info("exists slot")
	var count int64
	err := pgbs.db.GetContext(ctx, &count, "select count(*) from slot where id=$1;", idSlot)
	if err != nil {
		pgbs.Log.Error("databases error ", err)
	}
	if count == 0 {
		if create {
			_, err = pgbs.db.ExecContext(ctx, "INSERT into slot values ($1);", idSlot)
			if err != nil {
				pgbs.Log.Error(err)
			}
			return true, nil
		} else {
			return false, nil
		}
	}
	return true, nil
}

// AddBanner - Добавить баннер в ротацию
func (pgbs *PgBannerStorage) AddBanner(ctx context.Context, idBanner int64, idSlot int64) error {
	pgbs.Log.Info("bd add banner")
	existsBanner, err := pgbs.existsBanner(ctx, idBanner, true)
	if err != nil {
		pgbs.Log.Error("error exists banner ", err)
	}
	if !existsBanner {
		pgbs.Log.Error("banner not exists ", err)
	}
	existsSlot, err := pgbs.existsSlot(ctx, idSlot, true)
	if err != nil {
		pgbs.Log.Error("error exists banner ", err)
	}
	if !existsSlot {
		pgbs.Log.Error("slot not exists ", err)
	}
	_, err = pgbs.db.ExecContext(ctx, "insert into rotations(id_banner, id_slot) values ($1, $2);", idBanner, idSlot)
	if err != nil {
		pgbs.Log.Error(err)
	}
	return nil
}

// DelBanner - Удалить баннер из ротаций
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
