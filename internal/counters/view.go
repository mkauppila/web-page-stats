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

func (v *ViewCounter) GetCount(ctx context.Context, path string) (handler.ViewCount, error) {
	sql := `
		SELECT count 
  		  FROM view_count 
         WHERE path = ?`

	var counts int
	err := v.db.
		QueryRowContext(ctx, sql, path).
		Scan(&counts)
	if err != nil {
		return handler.ViewCount{}, fmt.Errorf("viewCounter.GetCount: %w", err)
	}

	return handler.ViewCount{
		Count: counts,
	}, nil
}

func (v *ViewCounter) Update(ctx context.Context, path string) (handler.ViewCount, error) {
	sql := `
		INSERT INTO view_count 
			 VALUES (?, 1)
		ON CONFLICT (path) DO
  	         UPDATE SET count=count + 1
	  	  RETURNING count`

	var counts int
	err := v.db.
		QueryRowContext(ctx, sql, path).
		Scan(&counts)
	if err != nil {
		return handler.ViewCount{}, fmt.Errorf("viewCounter.Update: %w", err)
	}

	return handler.ViewCount{
		Count: counts,
	}, nil
}
