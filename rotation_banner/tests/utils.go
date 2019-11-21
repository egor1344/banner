package tests

import (
	"testing"

	"github.com/jmoiron/sqlx"
)

// TruncateDb - очистка таблиц для адекватной проверки фукционала
func TruncateDb(t *testing.T, DB *sqlx.DB) {
	tables := []string{"banners", "rotations", "slot", "soc_dem_group", "statistic"}
	for _, table := range tables {
		_, err := DB.Query("truncate table " + table + " restart identity cascade;")
		if err != nil {
			t.Fatal(err)
		}
	}
}
