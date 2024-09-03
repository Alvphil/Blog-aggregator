package main

import (
	"time"

	"github.com/Alvphil/Blog-aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAT time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAT: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAT time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(dbfeed database.Feed) Feed {
	return Feed{
		ID:        dbfeed.ID,
		CreatedAT: dbfeed.CreatedAt,
		UpdatedAt: dbfeed.UpdatedAt,
		Name:      dbfeed.Name,
		Url:       dbfeed.Url,
		UserID:    dbfeed.UserID,
	}
}

type FollowFeed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAT time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFollowFeed(dbfeed database.FeedFollow) FollowFeed {
	return FollowFeed{
		ID:        dbfeed.ID,
		CreatedAT: dbfeed.CreatedAt,
		UpdatedAt: dbfeed.UpdatedAt,
		UserID:    dbfeed.UserID,
		FeedID:    dbfeed.FeedID,
	}
}

type CreateAndFollow struct {
	Feed       Feed       `json:"feed"`
	FeedFollow FollowFeed `json:"feed_follow"`
}

func databaseCreateAndFollow(feed Feed, feedFollow FollowFeed) CreateAndFollow {
	return CreateAndFollow{
		Feed:       feed,
		FeedFollow: feedFollow,
	}
}
