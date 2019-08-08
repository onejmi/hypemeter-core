package component

import (
	"encoding/hex"
	"encoding/json"
	googleAuthIDTokenVerifier "github.com/futurenda/google-auth-id-token-verifier"
	"github.com/gin-gonic/gin"
	"github.com/heroku/hypemeter-core/lib/data"
	"github.com/heroku/hypemeter-core/lib/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"io/ioutil"
	"net/http"
	"time"
)

var audiences = []string{
	"897244323047-g8j3g0ntsvdsvd4rudi58kti8cp1h6jc.apps.googleusercontent.com",
	"897244323047-nqvfn4vpio92gcmd90l39mt711jmdfcp.apps.googleusercontent.com",
}

type AuthenticationRequest struct {
	IdToken     string `json:"id_token"`
	AccessToken string `json:"access_token"`
}

func HandleLogin(c *gin.Context) {
	rawBody, _ := ioutil.ReadAll(c.Request.Body)
	var auth AuthenticationRequest
	err := json.Unmarshal(rawBody, &auth)

	if err == nil {
		verifier := googleAuthIDTokenVerifier.Verifier{}
		verr := verifier.VerifyIDToken(auth.IdToken, audiences)
		if verr == nil {
			googleProfile, derr := googleAuthIDTokenVerifier.Decode(auth.IdToken)
			if derr == nil {
				if !data.Exists("profiles", bson.D{{Key: "id", Value: googleProfile.Sub}}) {
					data.Insert("profiles", data.Profile{
						Id:       googleProfile.Sub,
						Email:    googleProfile.Email,
						Username: util.GrabYoutubeName(googleProfile.Email, auth.AccessToken),
						Picture:  googleProfile.Picture,
						Tier:     0,
						GoogleAuth: data.OAuth{
							AccessToken: auth.AccessToken,
						},
						CreationTime: time.Now().Unix(),
					})
				}
				var session data.Session
				if !data.Exists("sessions", bson.D{{Key: "user_id", Value: googleProfile.Sub}}) {
					sessionUUID, _ := uuid.New()
					data.Insert("sessions", data.Session{
						SessionID:    hex.EncodeToString(sessionUUID[:]),
						UserID:       googleProfile.Sub,
						CreationTime: time.Now().Unix(),
					})
				}
				_ = data.GetOne("sessions", bson.D{{Key: "user_id", Value: googleProfile.Sub}}, &session)
				c.JSON(http.StatusOK, session)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": "Failed to decode ID Token",
				})
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "Invalid ID Token",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Is your request body empty or missing fields?",
		})
	}

}
