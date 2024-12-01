-- +goose Up
-- +goose StatementBegin
CREATE TABLE view_count (
  path text CHECK(length(trim(path)) > 0) NOT NULL,
  count int CHECK(count >= 0) NOT NULL,
  PRIMARY KEY(path)
);

CREATE TABLE reaction_count (
  path text CHECK(length(trim(path)) > 0) NOT NULL,
  love integer CHECK(love >= 0) NOT NULL,
  like integer CHECK(like >= 0) NOT NULL,
  mindblown integer CHECK(mindblown >= 0) NOT NULL,
  puzzling integer CHECK(puzzling >= 0) NOT NULL,
  PRIMARY KEY(path)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE view_count;
DROP TABLE reaction_count;
-- +goose StatementEnd
