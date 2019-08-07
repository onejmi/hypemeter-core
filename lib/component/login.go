package component

import (
	"encoding/json"
	googleAuthIDTokenVerifier "github.com/futurenda/google-auth-id-token-verifier"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
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

	if err != nil {
		verifier := googleAuthIDTokenVerifier.Verifier{}
		verr := verifier.VerifyIDToken(auth.IdToken, audiences)
		if verr == nil {
			claimSet, derr := googleAuthIDTokenVerifier.Decode(auth.IdToken)
			if derr == nil {
				c.JSON(http.StatusOK, claimSet)
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
			"status": "Please provide an ID and Access Token",
		})
	}

}
