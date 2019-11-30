package models

// Rotation - модель обьекта ротации
type Rotation struct {
	ID       int64 `db:"id" json:"-"`
	IDBanner int64 `db:"id_banner" json:"id_banner"`
	IDSlot   int64 `db:"id_slot" json:"id_slot"`
}
