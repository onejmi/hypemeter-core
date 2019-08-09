package util

import (
	"context"
	"fmt"
	"github.com/heroku/hypemeter-core/lib/data"
	"github.com/jasonlvhit/gocron"
	"go.mongodb.org/mongo-driver/bson"
)

func SetupCronJobs() {
	gocron.Every(4).Hours().Do(invalidateExpiredSessions)
	gocron.Start()
}

func invalidateExpiredSessions() {
	cursor := data.List("sessions", bson.D{})
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var session data.Session
		attemptError := cursor.Decode(&session)
		if attemptError != nil {
			fmt.Println(attemptError.Error())
			break
		}
		if session.Expired() {
			_, err := data.Delete("sessions", bson.D{{Key: "session_id", Value: session.SessionID}})
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
