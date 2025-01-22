-- +goose Up
-- +goose StatementBegin
CREATE UNIQUE INDEX unique_idx_messages_content ON messages ((content::text));
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE messages;
-- +goose StatementEnd
