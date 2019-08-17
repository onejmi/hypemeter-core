package data

type Profile struct {
	Id           string   `json:"id" bson:"id"`
	Username     string   `json:"username" bson:"username"`
	Email        string   `json:"email" bson:"email"`
	Picture      string   `json:"picture" bson:"picture"`
	Tier         int      `json:"tier" bson:"tier"`
	FollowCount  int      `json:"follow_count" bson:"follow_count"`
	Following    []string `json:"following" bson:"following"`
	GoogleAuth   OAuth    `json:"google_auth" bson:"google_auth"`
	CreationTime int64    `json:"creation_time" bson:"creation_time"`
}

type OAuth struct {
	AccessToken string `json:"access_token" bson:"access_token"`
}
