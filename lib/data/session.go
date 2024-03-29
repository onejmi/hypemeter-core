package data

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

const lifeDuration int64 = 3600

type Session struct {
	SessionID    string `json:"session_id" bson:"session_id"`
	UserID       string `json:"user_id" bson:"user_id"`
	CreationTime int64  `json:"creation_time" bson:"creation_time"`
}

func (s Session) Expired() bool {
	var currTime = time.Now().Unix()
	var timeElapsed = currTime - s.CreationTime
	return timeElapsed > lifeDuration
}

func (s Session) Remove() error {
	_, err := Delete("sessions", bson.D{{Key: "session_id", Value: s.SessionID}})
	return err
}

func Authorize(c *gin.Context) {
	sessionID := c.Request.Header.Get("Session-ID")
	sessionConstraint := bson.D{{Key: "session_id", Value: sessionID}}
	if Exists("sessions", sessionConstraint) {
		var session Session
		_ = GetOne("sessions", sessionConstraint, &session)
		if session.Expired() {
			_ = session.Remove()
			_ = c.AbortWithError(http.StatusUnauthorized, ResponseErr{Status: "Invalid Session ID"})
			c.JSON(http.StatusUnauthorized, c.Errors.Last().JSON())
			return
		} else {
			c.Next()
			return
		}
	} else {
		_ = c.AbortWithError(http.StatusUnauthorized, ResponseErr{Status: "Invalid Session ID"})
		c.JSON(http.StatusUnauthorized, c.Errors.Last().JSON())
		return
	}
}

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "Your current session is valid",
	})
}
