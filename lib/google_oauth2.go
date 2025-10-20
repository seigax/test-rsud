package lib

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

type GoogleOauth2 struct {
	*oauth2.Config
}

type GoogleOauth2Callback struct {
	Sub     string
	Email   string
	Name    string
	Picture string
}

func (g *GoogleOauth2) GetLoginURL() string {
	return g.AuthCodeURL("")
}

func (g *GoogleOauth2) GetByCode(code string) (resp GoogleOauth2Callback, err error) {
	tok, err := g.Exchange(context.Background(), code)
	if err != nil {
		return
	}

	client := g.Client(context.Background(), tok)
	result, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return
	}
	defer result.Body.Close()
	data, _ := io.ReadAll(result.Body)

	if result.StatusCode == http.StatusOK {
		if err != nil {
			err = Oauth2Error
			return
		} else {
			m := make(map[string]interface{})
			err = json.Unmarshal([]byte(string(data)), &m)
			if err != nil {
				err = Oauth2Error
				return
			}
			resp.Sub = m["sub"].(string)
			resp.Email = m["email"].(string)
			resp.Name = m["name"].(string)
			resp.Picture = m["picture"].(string)
			return
		}
	} else {
		err = Oauth2Error
		return
	}

}
