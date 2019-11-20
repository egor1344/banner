package postgres

import (
	"context"

	"github.com/egor1344/banner/rotation_banner/pkg/ucb1"
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
	//pgbs.Log.Info("exists banner")
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
	//pgbs.Log.Info("exists slot")
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

// existsSocDemGroup - Проверка на существование соц. дем. группы
func (pgbs *PgBannerStorage) existsSocDemGroup(ctx context.Context, idSocDemGroup int64, create bool) (bool, error) {
	//pgbs.Log.Info("exists slot")
	var count int64
	err := pgbs.db.GetContext(ctx, &count, "select count(*) from soc_dem_group where id=$1;", idSocDemGroup)
	if err != nil {
		pgbs.Log.Error("databases error ", err)
	}
	if count == 0 {
		if create {
			_, err = pgbs.db.ExecContext(ctx, "INSERT into soc_dem_group values ($1);", idSocDemGroup)
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

// existsStatistic - Проверка на существование записи о статистике с данными параметрами
func (pgbs *PgBannerStorage) existsStatistic(ctx context.Context, idBanner, idSocDemGroup, idSlot int64, create bool) (bool, error) {
	//pgbs.Log.Info("exists slot")
	var count int64
	err := pgbs.db.GetContext(ctx, &count, "select count(*) from statistic where id_banner=$1 and id_soc_dem=$2 and id_slot=$3", idBanner, idSocDemGroup, idSlot)
	if err != nil {
		pgbs.Log.Error("databases error ", err)
	}
	if count == 0 {
		if create {
			_, err = pgbs.db.ExecContext(ctx, "INSERT into statistic(id_banner, id_soc_dem, count_click, count_views, id_slot) values ($1, $2, $3, $4, $5);", idBanner, idSocDemGroup, 1, 1, idSlot)
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
func (pgbs *PgBannerStorage) AddBanner(ctx context.Context, idBanner, idSlot int64) error {
	//pgbs.Log.Info("bd add banner")
	existsBanner, err := pgbs.existsBanner(ctx, idBanner, true)
	if err != nil {
		pgbs.Log.Error("error exists banner ", err)
	}
	if !existsBanner {
		pgbs.Log.Error("banner not exists")
	}
	existsSlot, err := pgbs.existsSlot(ctx, idSlot, true)
	if err != nil {
		pgbs.Log.Error("error exists banner ", err)
	}
	if !existsSlot {
		pgbs.Log.Error("slot not exists")
	}
	_, err = pgbs.db.ExecContext(ctx, "insert into rotations(id_banner, id_slot) values ($1, $2);", idBanner, idSlot)
	if err != nil {
		pgbs.Log.Error(err)
		return err
	}
	return nil
}

// DelBanner - Удалить баннер из ротаций
func (pgbs *PgBannerStorage) DelBanner(ctx context.Context, idBanner int64) error {
	//pgbs.Log.Info("bd del banner")
	_, err := pgbs.db.ExecContext(ctx, "delete from rotations where id_banner=$1;", idBanner)
	if err != nil {
		pgbs.Log.Error(err)
		return err
	}
	return nil
}

// CountTransition - Засчитать переход
func (pgbs *PgBannerStorage) CountTransition(ctx context.Context, idBanner, idSocDemGroup int64, idSlot int64) error {
	//pgbs.Log.Info("bd count transition")
	existsSocDemGroup, err := pgbs.existsSocDemGroup(ctx, idSocDemGroup, true)
	if err != nil {
		pgbs.Log.Error("error exists socDemGroup ", err)
	}
	if !existsSocDemGroup {
		pgbs.Log.Error("socDemGroup not exists")
	}
	existsBanner, err := pgbs.existsBanner(ctx, idBanner, true)
	if err != nil {
		pgbs.Log.Error("error exists banner ", err)
	}
	if !existsBanner {
		pgbs.Log.Error("banner not exists")
	}
	existsSlot, err := pgbs.existsSlot(ctx, idSlot, true)
	if err != nil {
		pgbs.Log.Error("error exists banner ", err)
	}
	if !existsSlot {
		pgbs.Log.Error("slot not exists")
	}
	existsStatistic, err := pgbs.existsStatistic(ctx, idBanner, idSocDemGroup, idSlot, true)
	if err != nil {
		pgbs.Log.Error("error exists banner ", err)
	}
	if !existsStatistic {
		pgbs.Log.Error("slot not exists")
	}
	_, err = pgbs.db.ExecContext(ctx, "update statistic set count_click = count_click+1"+
		"where id_banner=$1 and id_soc_dem=$2 and id_slot=$3", idBanner, idSocDemGroup, idSlot)
	if err != nil {
		pgbs.Log.Error(err)
		return err
	}
	return nil
}

// incCountView - Увеличение количества показов в статистике
func (pgbs *PgBannerStorage) incCountView(ctx context.Context, idBanner, idSlot, idSocDemGroup int64) error {
	pgbs.Log.Info("bd get banner")
	_, err := pgbs.db.ExecContext(ctx, `update statistic set count_views = count_views + 1 where id_slot = $1 and id_soc_dem = $2 and id_banner=$3;`, idSlot, idSocDemGroup, idBanner)
	if err != nil {
		pgbs.Log.Error("error incCountView ", err)
	}
	return nil
}

// GetBanner - Выбрать баннер для показа
func (pgbs *PgBannerStorage) GetBanner(ctx context.Context, idSlot, idSocDemGroup int64) (int64, error) {
	//pgbs.Log.Info("bd get banner")
	rows, err := pgbs.db.QueryxContext(ctx, `select id_banner,
														 count_click,
                                                         count_views,
														 SUM(count_views) OVER
															  (PARTITION BY id_slot) AS all_count_views
													from statistic
														 where id_slot = $1
														 and id_soc_dem = $2;`, idSlot, idSocDemGroup)
	if err != nil {
		pgbs.Log.Error("error get banner ", err)
	}
	var s struct {
		IdBanner      int64 `db:"id_banner"`
		CountClick    int64 `db:"count_click"`
		CountView     int64 `db:"count_views"`
		AllCountViews int64 `db:"all_count_views"`
	}
	var lbs ucb1.ListBannerStatistic
	for rows.Next() {
		err = rows.StructScan(&s)
		lbs.Objects = append(lbs.Objects, &ucb1.BannerStatistic{CountClick: s.CountClick, ID: s.IdBanner, CountDisplay: s.CountView})
	}
	lbs.AllCountDisplay = s.AllCountViews
	id, err := lbs.GetRelevantObject()
	if err != nil {
		pgbs.Log.Error("error in GetRelevantObject ", err)
	}
	err = pgbs.incCountView(ctx, id, idSlot, idSocDemGroup)
	if err != nil {
		pgbs.Log.Error("error in GetRelevantObject ", err)
	}
	return id, nil
}
