package data

import "time"

const lifeDuration int64 = 3600 * 1000

type Session struct {
	SessionID    string `json:"session_id" bson:"session_id"`
	UserID       string `json:"user_id" bson:"user_id"`
	CreationTime int64  `json:"creation_time" bson:"creation_time"`
}

func (s Session) Expired() bool {
	var currTime = time.Now().Unix()
	return s.CreationTime < (currTime - lifeDuration)
}
