package counters

import (
	"context"
	"database/sql"
	"fmt"

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
	sql := `
		SELECT count 
  		  FROM view_count 
         WHERE category = ?
	       AND slug = ?`
	results, err := v.db.QueryContext(context.Background(), sql, category, slug)
	if err != nil {
		return handler.ViewCount{}, fmt.Errorf("viewCounter.GetCount: %w", err)
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

func (v *ViewCounter) Update(category, slug string) (handler.ViewCount, error) {
	sql := `
		INSERT INTO view_count 
			 VALUES (?, ?, 1)
		ON CONFLICT (category, slug) DO
  	         UPDATE SET count=count + 1
	  	  RETURNING count`
	results, err := v.db.QueryContext(context.Background(), sql, category, slug)
	if err != nil {
		return handler.ViewCount{}, fmt.Errorf("viewCounter.Update: %w", err)
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
