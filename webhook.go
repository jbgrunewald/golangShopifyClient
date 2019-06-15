package shopify

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	ProductCreate  = "products/create"
	ProductUpdate  = "products/update"
	ProductDelete  = "products/delete"
	AppUninstalled = "app/uninstalled"

	XShopifyShopDomain = "X-Shopify-Shop-Domain"
	XShopifyHmacSha256 = "X-Shopify-Hmac-Sha256"
)

type Webhook struct {
	Address             string   `json:"address"`
	Created_at          string   `json:"created_at, omitempty"`
	Fields              []string `json:"fields omitempty"`
	Format              string   `json:"format"`
	Id                  int      `json:"id, omitempty"`
	MetafieldNamespaces []string `json:"metafield_namespaces, omitempty"`
	Topic               string   `json:"topic"`
	UpdatedAt           string   `json:"updated_at, omitempty"`
}

type WebhookWrapper struct {
	Webhook Webhook `json:"webhook"`
}

type WebhooksWrapper struct {
	Webhooks []Webhook `json:"webhooks"`
	Errors   string    `json:"errors"`
}

type WebHookRequestOptions struct {
	Address      string   `url:"address,omitempty"`
	CreatedAtMax string   `url:"created_at_max,omitempty"`
	CreatedAtMin string   `url:"created_at_min,omitempty"`
	Fields       []string `url:"fields,omitempty"`
	Limit        int      `url:"limit,omitempty"`
	SinceId      int      `url:"since_id,omitempty"`
	Topic        string   `url:"topic,omitempty"`
	UpdatedAtMin string   `url:"updated_at_min,omitempty"`
	UpdatedAtMax string   `url:"updated_at_max,omitempty"`
}

func (client *ShopifyApiImpl) CreateWebhook(details ShopifyRequestDetails, request Webhook) (result Webhook, err error) {
	requestUrl := "https://" + details.ShopName + "/admin/webhooks.json"

	log.Printf("Requesting to create webhook for topic %s for shop %s with URL %s", request.Topic, details.ShopName, requestUrl)

	requestStr, err := json.Marshal(WebhookWrapper{request})
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

	resp, err := client.Http.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return
	}

	buf, err := ioutil.ReadAll(resp.Body)
	log.Println("This is the response from the webhook create request: ", string(buf))
	wrapper := WebhookWrapper{}
	if err != nil {
		return
	}

	if resp.StatusCode != 201 {
		err = errors.New("The webhook request was unsuccesful")
		return
	}

	err = json.Unmarshal(buf, &wrapper)
	if err != nil {
		return
	}

	result = wrapper.Webhook

	if result.Id == 0 {
		err = errors.New("The request returned but the webhook has no Id, which implies it did not succeed")
		return
	}

	return
}

func (client *ShopifyApiImpl) DeleteWebhook(details ShopifyRequestDetails, request Webhook) (err error) {
	requestUrl := "https://" + details.ShopName + "/admin/api/2019-04/webhooks/" + string(request.Id) + ".json"

	log.Printf("Requesting to delete webhook for topic %s for shop %s with URL %s", request.Topic, details.ShopName, requestUrl)

	req, err := http.NewRequest("DELETE", requestUrl, nil)
	if err != nil {
		return
	}

	req.WithContext(details.Ctx)
	req.Header.Add("X-Shopify-Access-Token", details.AccessToken)

	resp, err := client.Http.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return
	}

	buf, err := ioutil.ReadAll(resp.Body)
	log.Println("This is the response from the webhook delete request: ", string(buf))
	if err != nil {
		return
	}

	return
}

func (client ShopifyApiImpl) GetWebhooks(details ShopifyRequestDetails, options WebHookRequestOptions) (webhooks []Webhook, err error) {
	v, err := query.Values(options)
	if err != nil {
		log.Println("there's an issue setting up the query params in the get webhooks request")
		return
	}
	requestUrl := "https://" + details.ShopName + "/admin/api/2019-04/webhooks.json?" + v.Encode()
	log.Println(requestUrl)

	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return
	}

	req.Header.Add("X-Shopify-Access-Token", details.AccessToken)

	resp, err := client.Http.Do(req)
	if err != nil {
		return
	}

	buf, _ := ioutil.ReadAll(resp.Body)
	log.Println("This is the response for the webhooks: ", string(buf))
	wrapper := WebhooksWrapper{}

	err = json.Unmarshal(buf, &wrapper)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(wrapper.Errors)
		return
	}

	webhooks = wrapper.Webhooks

	if len(webhooks) < options.Limit {
		return
	}

	options.SinceId = webhooks[len(webhooks)-1].Id
	nextResult, err := client.GetWebhooks(details, options)

	webhooks = append(webhooks, nextResult...)

	return
}