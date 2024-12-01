-- +goose Up
-- +goose StatementBegin
CREATE TABLE view_count (
  category text CHECK(category IN ('articles','snippets')) NOT NULL,
  slug text NOT NULL,
  count int CHECK(count >= 0) NOT NULL,
  PRIMARY KEY(category, slug)
);

CREATE TABLE reaction_count (
  category text CHECK(category IN ('articles','snippets')) NOT NULL,
  slug text NOT NULL,
  love integer CHECK(love >= 0) NOT NULL,
  like integer CHECK(like >= 0) NOT NULL,
  mindblown integer CHECK(mindblown >= 0) NOT NULL,
  puzzling integer CHECK(puzzling >= 0) NOT NULL,
  PRIMARY KEY(category, slug)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE view_count;
DROP TABLE reaction_count;
-- +goose StatementEnd
