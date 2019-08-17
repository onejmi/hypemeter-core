package component

import (
	"github.com/gin-gonic/gin"
	"github.com/heroku/hypemeter-core/lib/data"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type PublicProfile struct {
	Username     string   `json:"username"`
	Picture      string   `json:"picture"`
	Tier         int      `json:"tier"`
	FollowCount  int      `json:"follow_count"`
	Following    []string `json:"following"`
	CreationTime int64    `json:"creation_time"`
}

func ViewProfile(c *gin.Context) {
	username := c.Query("username")
	var publicProfile PublicProfile

	if username != "" {
		if (data.Exists("profiles", bson.D{{Key: "username", Value: username}})) {
			_ = data.GetOne("profiles", bson.D{{Key: "username", Value: username}}, &publicProfile)
			c.JSON(http.StatusOK, publicProfile)
			return
		}
	} else {
		var currentSession data.Session
		err := data.GetOne("sessions", bson.D{{Key: "session_id", Value: c.Request.Header.Get("Session-ID")}}, &currentSession)
		if err == nil {
			userErr := data.GetOne("profiles", bson.D{{Key: "id", Value: currentSession.UserID}}, &publicProfile)
			if userErr == nil {
				c.JSON(http.StatusOK, publicProfile)
				return
			}
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"status": "That user could not be found",
	})
}
