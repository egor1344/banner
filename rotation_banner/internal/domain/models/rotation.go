package models

type Rotation struct {
	Id       int64 `db:"id"json:"-"`
	IdBanner int64 `db:"id_banner"json:"id_banner"`
	IdSlot   int64 `db:"id_slot"json:"id_slot"`
}
