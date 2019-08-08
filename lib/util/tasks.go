package util

import (
	"context"
	"github.com/heroku/hypemeter-core/lib/data"
	"github.com/jasonlvhit/gocron"
	"go.mongodb.org/mongo-driver/bson"
)

func SetupCronJobs() {
	gocron.Every(3).Minutes().Do(invalidateExpiredSessions)
}

func invalidateExpiredSessions() {
	cursor := data.List("sessions", bson.D{{}})

	for {
		var session data.Session
		attemptError := cursor.Decode(&session)
		if attemptError == nil {
			break
		}
		if session.Expired() {
			_, err := data.Delete("sessions", bson.D{{Key: "session_id", Value: session.SessionID}})
			if err != nil {
				print(err)
			}
		}
		if !cursor.Next(context.TODO()) {
			break
		}
	}
}
