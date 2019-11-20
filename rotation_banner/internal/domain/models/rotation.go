package models

type Rotation struct {
	Id       int64 `db:"id"`
	IdBanner int64 `db:"id_banner"`
	IdSlot   int64 `db:"id_slot"`
}
