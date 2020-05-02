package shopify

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type ScriptTag struct {
	CreatedAt    string `json:"created_at,omitempty"`
	Event        string `json:"event,omitempty"`
	Id           int    `json:"id,omitempty"`
	Src          string `json:"src,omitempty"`
	DisplayScope string `json:"display_scope,omitempty"`
	UpdateAt     string `json:"update_at,omitempty"`
}

type ScriptTageWrapper struct {
	ScriptTag ScriptTag `json:"script_tag"`
}

func (c *RestAdminClient) ScriptTagCreate(details Ctx, request ScriptTag) (result ScriptTag, err error) {
	requestUrl := "https://" + details.ShopName + "/admin/" + c.Version.String() + "script_tags.json"

	c.Logger.Printf("Making the script tag request for shop %s using URL %s\n", details.ShopName, requestUrl)

	requestStr, err := json.Marshal(ScriptTageWrapper{request})
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(requestStr))
	if err != nil {
		return
	}

	req.WithContext(details.Ctx)

	req.Header.Add("X-Shopify-Access-Token", details.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Http.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return
	}

	if resp.StatusCode != 201 {
		c.Logger.Println("The billing request response status code is: ", resp.StatusCode)
		return result, errors.New("Received non 201 response code from the server")
	}

	buf, _ := ioutil.ReadAll(resp.Body)
	c.Logger.Println("The response for the recurring billing request is: ", string(buf))
	wrapper := ScriptTageWrapper{}
	err = json.Unmarshal(buf, &wrapper)
	if err != nil {
		c.Logger.Println("Error unmarshaling the script tag response", err.Error())
	}

	c.Logger.Println("The result of the unmarshaling: ", wrapper)
	result = wrapper.ScriptTag

	return
}
