package feed

import (
	"errors"
	"os"
	"strconv"
)

var FeedId int64

func ParseFeed() error {
	args := os.Args
	if len(args) < 2 {
		return errors.New("No feed id was provided. Please provide an integer value for the feed you would like to use.")
	}

	feedId := args[1]
	id, err := strconv.ParseInt(feedId, 10, 64)
	if err != nil {
		return errors.New("Unable to parse feed id. Please provide an integer value for the feed id.")
	}
	FeedId = id
	return nil
}
