package models

import (
	"time"

	"github.com/kysal/go-feed/db"
	"github.com/kysal/go-feed/feed"
)

type Post struct {
	Id        int64
	Content   string
	CreatedAt time.Time
	FeedId    int64
}

func (p *Post) Publish() error {
	result, err := db.DB.Exec("INSERT INTO posts(content, created_at, feed_id) VALUES (?, ?, ?)", p.Content, time.Now().Unix(), p.FeedId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.Id = id

	return nil

}

func GetAllPosts() ([]Post, error) {
	rows, err := db.DB.Query("SELECT * FROM posts WHERE feed_id = ? LIMIT 100", feed.FeedId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var currentPost Post
		err = rows.Scan(&currentPost.Id, &currentPost.Content, &currentPost.CreatedAt, &currentPost.FeedId)
		if err != nil {
			return posts, err
		}
		posts = append(posts, currentPost)
	}

	err = rows.Err()
	if err != nil {
		return posts, err
	}

	return posts, nil
}
