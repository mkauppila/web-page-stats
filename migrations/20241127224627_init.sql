-- +goose Up
-- +goose StatementBegin
CREATE TABLE count (
  category text NOT NULL,
  slug text NOT NULL,
  type text CHECK(type IN ('articles','snippets')) NOT NULL,
  count integer CHECK(count >0) NOT NULL,
  PRIMARY KEY(category, slug)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE count;
-- +goose StatementEnd
