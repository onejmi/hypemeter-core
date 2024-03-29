package util

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"google.golang.org/api/youtube/v3"
	"net/http"
	"strings"
)

type YoutubeProfile struct {
	Name        string
	PictureLink string
}

func GrabYoutubeProfile(email string, accessToken string) (profile YoutubeProfile) {

	var name string
	var pictureLink string

	client := genClient(accessToken)
	youtubeService, err := youtube.New(client)

	if err != nil {
		print(err)
	}

	ownChannelRequest, _ := youtubeService.Channels.List("snippet,statistics").Mine(true).Do()
	items := ownChannelRequest.Items

	var topSubCount uint64 = 0
	for _, item := range items {
		if item.Statistics.SubscriberCount >= topSubCount {
			topSubCount = item.Statistics.SubscriberCount
			name = item.Snippet.Title
			pictureLink = item.Snippet.Thumbnails.Default.Url
		}
	}

	if name == "" {
		cutIndex := strings.IndexRune(email, '@')
		name = email[:cutIndex]
		fmt.Println(name)
		pictureLink = ""
	}

	return YoutubeProfile{
		Name:        name,
		PictureLink: pictureLink,
	}
}

func genClient(accessToken string) *http.Client {
	return oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken}))
}
