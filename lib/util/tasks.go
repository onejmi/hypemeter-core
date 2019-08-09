package util

import (
	"context"
	"fmt"
	"github.com/heroku/hypemeter-core/lib/data"
	"github.com/jasonlvhit/gocron"
	"go.mongodb.org/mongo-driver/bson"
)

func SetupCronJobs() {
	gocron.Every(2).Minutes().Do(invalidateExpiredSessions)
	gocron.Start()
}

func invalidateExpiredSessions() {
	cursor := data.List("sessions", bson.D{{}})

	for {
		var session data.Session
		attemptError := cursor.Decode(&session)
		fmt.Println("wait for it...")
		if attemptError != nil {
			fmt.Println(attemptError.Error())
			break
		}
		fmt.Println("wait what")
		if session.Expired() {
			fmt.Println("um")
			_, err := data.Delete("sessions", bson.D{{Key: "session_id", Value: session.SessionID}})
			if err != nil {
				fmt.Println(err)
			}
		}
		if !cursor.Next(context.TODO()) {
			fmt.Println("huh")
			break
		}
	}
}
