package models

type Statistic struct {
	Id            int64 `db:"id"json:"id,omitempty"`
	IdBanner      int64 `db:"id_banner"json:"id_banner,omitempty"`
	IdSlot        int64 `db:"id_slot"json:"id_slot,omitempty"`
	IdSocDemGroup int64 `db:"id_soc_dem"json:"id_soc_dem,omitempty"`
	CountClick    int64 `db:"count_click"json:"count_click,omitempty"`
	CountViews    int64 `db:"count_views"json:"count_views,omitempty"`
}
