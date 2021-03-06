package sqlstore

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

func TestDB(t *testing.T, databaseUrl string) (*sql.DB, func(...string)) {
	//t.Helper()
	//
	//config := NewConfig()
	//config.DatabaseUrl = databaseUrl
	//s := New(config)
	//if err := s.Open(); err != nil {
	//	t.Fatal(err)
	//}
	//
	//return s, func(tables ...string) {
	//	if len(tables) > 0 {
	//		if  _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
	//			t.Fatal(err)
	//		}
	//	}
	//
	//	s.Close()
	//}

	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
			if len(tables) > 0 {
				db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
				db.Close()
			}
	}
}
