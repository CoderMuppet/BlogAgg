// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: feed_follow.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedsFollow = `-- name: CreateFeedsFollow :one
    WITH inserted_feed_follows AS (
        INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING id, created_at, updated_at, user_id, feed_id
)
SELECT 
    inserted_feed_follows.id, inserted_feed_follows.created_at, inserted_feed_follows.updated_at, inserted_feed_follows.user_id, inserted_feed_follows.feed_id,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follows
INNER JOIN feeds ON inserted_feed_follows.feed_id = feeds.id
INNER JOIN users ON inserted_feed_follows.user_id = users.id
`

type CreateFeedsFollowParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
}

type CreateFeedsFollowRow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
	FeedName  string
	UserName  string
}

func (q *Queries) CreateFeedsFollow(ctx context.Context, arg CreateFeedsFollowParams) (CreateFeedsFollowRow, error) {
	row := q.db.QueryRowContext(ctx, createFeedsFollow,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.FeedID,
	)
	var i CreateFeedsFollowRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
		&i.FeedName,
		&i.UserName,
	)
	return i, err
}

const getFeedFollowsForUser = `-- name: GetFeedFollowsForUser :many

SELECT feed_follows.id, feed_follows.created_at, feed_follows.updated_at, feed_follows.user_id, feed_follows.feed_id, feeds.name AS feed_name, users.name AS user_name
FROM feed_follows
INNER JOIN feeds ON feed_follows.feed_id = feeds.id
INNER JOIN users ON feed_follows.user_id = users.id
WHERE feed_follows.user_id = $1
`

type GetFeedFollowsForUserRow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
	FeedName  string
	UserName  string
}

func (q *Queries) GetFeedFollowsForUser(ctx context.Context, userID uuid.UUID) ([]GetFeedFollowsForUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollowsForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFeedFollowsForUserRow
	for rows.Next() {
		var i GetFeedFollowsForUserRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.FeedID,
			&i.FeedName,
			&i.UserName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const unfollowFeed = `-- name: UnfollowFeed :exec

DELETE FROM feed_follows WHERE user_id = $1 AND feed_id = $2
`

type UnfollowFeedParams struct {
	UserID uuid.UUID
	FeedID uuid.UUID
}

func (q *Queries) UnfollowFeed(ctx context.Context, arg UnfollowFeedParams) error {
	_, err := q.db.ExecContext(ctx, unfollowFeed, arg.UserID, arg.FeedID)
	return err
}
