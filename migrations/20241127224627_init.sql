-- +goose Up
-- +goose StatementBegin
CREATE TABLE count (
  category text,
  slug text,
  type text,
  PRIMARY KEY(category, slug)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE count;
-- +goose StatementEnd
