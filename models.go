package main

import (
	"time"

	"github.com/Hrugved/rssagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"` 
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databasFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.ID,
	}
}

func databasFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _,dbFeed := range dbFeeds {
		feeds = append(feeds, databasFeedToFeed(dbFeed))
	}
	return feeds
}
 
type FeedFollow struct {
	ID        uuid.UUID `json:"id"` 
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databasFeedFollowToFeedFollow(dbFeedFollow database.FeedsFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID:    dbFeedFollow.ID,
		FeedID: dbFeedFollow.FeedID,
	}
}

func databasFeedsFollowsToFeedsFollows(dbFeedsFollows []database.FeedsFollow) []FeedFollow {
	feedsFollows := []FeedFollow{}
	for _,dbFeedFollow := range dbFeedsFollows {
		feedsFollows = append(feedsFollows, databasFeedFollowToFeedFollow(dbFeedFollow))
	}
	return feedsFollows
}