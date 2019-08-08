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
	print("logging")
	cursor := data.List("sessions", bson.D{{}})
	print("double logz")

	for {
		print("ok somewhere")
		var session data.Session
		attemptError := cursor.Decode(&session)
		print("wait for it...")
		if attemptError != nil {
			break
		}
		print("wait what")
		if session.Expired() {
			print("um")
			_, err := data.Delete("sessions", bson.D{{Key: "session_id", Value: session.SessionID}})
			if err != nil {
				print(err)
			}
		}
		if !cursor.Next(context.TODO()) {
			print("huh")
			break
		}
	}
}
