-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE todo (
  id         INT          NOT NULL AUTO_INCREMENT,
  name       VARCHAR(255) NOT NULL,
  complete  BOOLEAN      NOT NULL DEFAULT FALSE,
  deadline   DATETIME     NOT NULL,
  updated_at DATETIME     NOT NULL,
  created_at DATETIME     NOT NULL,
  PRIMARY KEY(id)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE todo;
-- +goose StatementEnd