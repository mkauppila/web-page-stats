package counters

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mkauppila/web-page-stats/internal/handler"
)

// try this:
// err := db.QueryRowContext(ctx, "SELECT username, created_at FROM users WHERE id=?", id).Scan(&username, &created)

type ReactionCounter struct {
	db *sql.DB
}

func CreateReactionCounter(db *sql.DB) *ReactionCounter {
	return &ReactionCounter{
		db: db,
	}
}

func (v *ReactionCounter) GetCount(ctx context.Context, category, slug string) (handler.ReactionCounts, error) {
	sql := `
		SELECT love, like, mindblown, puzzling
  		  FROM reaction_count
         WHERE category = ?
	       AND slug = ?`

	var love, like, mindblown, puzzling int
	err := v.db.
		QueryRowContext(ctx, sql, category, slug).
		Scan(&love, &like, &mindblown, &puzzling)
	if err != nil {
		return handler.ReactionCounts{}, fmt.Errorf("ReactionCounter.GetCount: %w", err)
	}

	return handler.ReactionCounts{
		Like:      like,
		Love:      love,
		Mindblown: mindblown,
		Puzzling:  puzzling,
	}, nil
}

func (v *ReactionCounter) Update(ctx context.Context, category, slug, reaction string) (handler.ReactionCounts, error) {
	var sql string
	switch reaction {
	case "love":
		sql = loveQuery()
	case "like":
		sql = likeQuery()
	case "mindblown":
		sql = mindblownQuery()
	case "puzzling":
		sql = puzzlingQuery()
	}

	var love, like, mindblown, puzzling int
	err := v.db.
		QueryRowContext(ctx, sql, category, slug).
		Scan(&love, &like, &mindblown, &puzzling)

	if err != nil {
		return handler.ReactionCounts{}, fmt.Errorf("viewCounter.Update: %w", err)
	}

	return handler.ReactionCounts{
		Like:      like,
		Love:      love,
		Mindblown: mindblown,
		Puzzling:  puzzling,
	}, nil
}

func loveQuery() string {
	return `INSERT INTO reaction_count 
   		    VALUES (?, ?, 1, 0, 0, 0)
   	   ON CONFLICT (category, slug) DO
	    UPDATE SET love=love + 1
	     RETURNING love, like, mindblown, puzzling`
}

func likeQuery() string {
	return `INSERT INTO reaction_count 
   		    VALUES (?, ?, 0, 1, 0, 0)
   	   ON CONFLICT (category, slug) DO
	    UPDATE SET like=like + 1
	     RETURNING love, like, mindblown, puzzling`
}

func mindblownQuery() string {
	return `INSERT INTO reaction_count 
   		    VALUES (?, ?, 0, 0, 1, 0)
   	   ON CONFLICT (category, slug) DO
	    UPDATE SET mindblown=mindblown + 1
	     RETURNING love, like, mindblown, puzzling`
}

func puzzlingQuery() string {
	return `INSERT INTO reaction_count 
   		    VALUES (?, ?, 0, 0, 0, 1)
   	   ON CONFLICT (category, slug) DO
	    UPDATE SET puzzling=puzzling + 1
	     RETURNING love, like, mindblown, puzzling`
}
