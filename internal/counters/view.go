package counters

import (
	"context"
	"database/sql"

	"github.com/mkauppila/web-page-stats/internal/handler"
)

type ViewCounter struct {
	db *sql.DB
}

func CreateViewCounter(db *sql.DB) *ViewCounter {
	return &ViewCounter{
		db: db,
	}
}

func (v *ViewCounter) GetCount(category, slug string) (handler.ViewCount, error) {
	return handler.ViewCount{}, nil
}

func (v *ViewCounter) Update(category, slug string) (handler.ViewCount, error) {
	// INSERT INTO Book (ID, Name)
	//  VALUES (1001, 'SQLite')
	//  ON CONFLICT (ID) DO
	//  UPDATE SET Name=excluded.Name;
	sql := `
		INSERT INTO view_count 
			 VALUES (?, ?, 0)
		ON CONFLICT (category, slug) DO
  	         UPDATE SET count=count + 1
	  	  RETURNING count`
	results, err := v.db.QueryContext(context.Background(), sql, category, slug)
	if err != nil {
		panic(err)
	}
	defer results.Close()

	var counts int
	for results.Next() {
		results.Scan(&counts)
	}

	return handler.ViewCount{
		Count: counts,
	}, nil
}
