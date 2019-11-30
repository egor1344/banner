package models

// Statistic - модель обьекта статистики
type Statistic struct {
	ID            int64 `db:"id" json:"id,omitempty"`
	IDBanner      int64 `db:"id_banner" json:"id_banner,omitempty"`
	IDSlot        int64 `db:"id_slot" json:"id_slot,omitempty"`
	IDSocDemGroup int64 `db:"id_soc_dem" json:"id_soc_dem,omitempty"`
	CountClick    int64 `db:"count_click" json:"count_click,omitempty"`
	CountViews    int64 `db:"count_views" json:"count_views,omitempty"`
}
