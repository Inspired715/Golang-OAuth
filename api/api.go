package api

import (
	"net/http"
	"net/url"
	"projects/modules/app"
	"projects/modules/e"
	"projects/omodels"

	"github.com/gin-gonic/gin"
)

func RedirectGoogle(c *gin.Context) {
	const rootURl = "https://accounts.google.com/o/oauth2/auth"

	values := url.Values{}
	values.Add("response_type", "token") // Use "token" for Implicit Grant flow
	values.Add("client_id", "130711294860-g7463001bnr2cpon6loujjci0k5toouu.apps.googleusercontent.com")
	values.Add("redirect_uri", "http://localhost/index.php")
	values.Add("scope", "https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email")

	query := values.Encode()
	redirectURL := rootURl + "?" + query
	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}

func GetList(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})

	results, err := omodels.GetBlogList()

	if err != nil {
		return
	}

	data["results"] = results

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
