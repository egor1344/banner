package models

type Statistic struct {
	Id            int64 `db:"id"`
	IdBanner      int64 `db:"id_banner"`
	IdSlot        int64 `db:"id_slot"`
	IdSocDemGroup int64 `db:"id_soc_dem"`
	CountClick    int64 `db:"count_click"`
	CountViews    int64 `db:"count_views"`
}
