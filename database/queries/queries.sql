-- name: Add :one
INSERT INTO messages (content) VALUES (@content) RETURNING id;

-- name: Find :one
SELECT * FROM messages WHERE content = (@content::jsonb);

-- name: List :many
SELECT * FROM messages; 
