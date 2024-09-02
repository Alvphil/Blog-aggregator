-- name: FollowFeed :one
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
VALUES($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteFollowFeed :exec
DELETE from feed_follows WHERE id = $1 AND user_id = $2;

-- name: GetAllFeedFollows :many
SELECT * from feed_follows WHERE user_id = $1;