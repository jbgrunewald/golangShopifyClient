package shopify

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
)

type Client interface {
	OAuthRequest(RequestDetails, OAuthRequest) (OAuthResponse, error)
	RecurringApplicationChargeCreate(details RequestDetails, request RecurringApplicationCharge) (RecurringApplicationCharge, error)
	RecurringApplicationChargeActivate(RequestDetails, RecurringApplicationCharge) (RecurringApplicationCharge, error)
	ShopGet(RequestDetails) (Shop, error)
	WebhookCreate(RequestDetails, Webhook) (Webhook, error)
	WebhookDelete(RequestDetails, Webhook) error
	WebhookList(RequestDetails, WebHookRequestOptions) ([]Webhook, error)
	ProductList(RequestDetails, ProductRequestOptions) ([]Product, error)
	CollectList(RequestDetails, CollectRequestOptions) ([]Collect, error)
	RecurringApplicationChargeList(RequestDetails, RecurringApplicationChargeOptons) ([]RecurringApplicationCharge, error)
	ScriptTagCreate(RequestDetails, ScriptTag) (ScriptTag, error)
}

type RestAdminClient struct {
	Http    *http.Client
	Logger  *log.Logger
	Version ApiVersion
}

type RequestDetails struct {
	ShopName    string
	AccessToken string
	Ctx         context.Context
}

func (r RequestDetails) GetBaseUrl() string {
	return r.ShopName
}

func (r RequestDetails) GetAuthToken() string {
	return r.AccessToken
}

type RequestBuilder interface {
	GetBaseUrl() string
	GetAuthToken() string
}

func (r RestAdminClient) get(details RequestBuilder, resource string) (buf []byte, err error) {
	requestUrl := "https://" + details.GetBaseUrl() + "/admin/" + r.Version.String() + resource + ".json"

	r.Logger.Printf("Requesting resource %s for shop %s using URL %s\n", resource, r.Version.String(), requestUrl)

	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return
	}

	req.Header.Add("X-Shopify-Access-Token", details.GetAuthToken())

	resp, err := r.Http.Do(req)
	if err != nil {
		return
	}

	buf, err = ioutil.ReadAll(resp.Body)
	r.Logger.Printf("response for %s received: %s", requestUrl, string(buf))

	return
}
