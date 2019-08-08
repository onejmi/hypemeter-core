package util

import (
	"context"
	"golang.org/x/oauth2"
	"google.golang.org/api/youtube/v3"
	"net/http"
	"strings"
)

func GrabYoutubeName(email string, accessToken string) (name string) {
	client := genClient(accessToken)
	youtubeService, err := youtube.New(client)

	if err != nil {
		print(err)
	}

	ownChannelRequest, _ := youtubeService.Channels.List("snippet,statistics").Mine(true).Do()
	items := ownChannelRequest.Items

	var topSubCount uint64 = -1
	for _, item := range items {
		if item.Statistics.SubscriberCount > topSubCount {
			topSubCount = item.Statistics.SubscriberCount
			name = item.Snippet.Title
		}
	}

	if name == "" {
		cutIndex := strings.IndexRune(email, '@')
		name = email[:cutIndex]
	}

	return name
}

func genClient(accessToken string) *http.Client {
	return oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken}))
}
