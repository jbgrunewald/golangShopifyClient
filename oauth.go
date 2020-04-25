package shopify

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

type OAuthResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
}

type OAuthRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
}

func (c *RestAdminClient) OAuthRequest(details ShopifyContext, request OAuthRequest) (result OAuthResponse, err error) {

	accessTokenRequestUrl := "https://" + details.ShopName + "/admin/oauth/access_token"

	requestStr, err := json.Marshal(request)
	if err != nil {
		return
	}

	resp, err := c.http.Post(accessTokenRequestUrl, "application/json", bytes.NewBuffer(requestStr))
	if err != nil {
		return
	}

	defer resp.Body.Close()

	buf, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(buf, &result)
	if err != nil {
		return
	}

	return
}
