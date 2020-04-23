package shopify

import (
	"context"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Client interface {
	OAuthRequest(Request, OAuthRequest) (OAuthResponse, error)
	RecurringApplicationChargeCreate(details Request, request RecurringApplicationCharge) (RecurringApplicationCharge, error)
	RecurringApplicationChargeActivate(Request, RecurringApplicationCharge) (RecurringApplicationCharge, error)
	ShopGet(Request) (Shop, error)
	WebhookCreate(Request, Webhook) (Webhook, error)
	WebhookDelete(Request, Webhook) error
	WebhookList(Request, WebHookRequestOptions) ([]Webhook, error)
	ProductList(Request, ProductRequestOptions) ([]Product, error)
	CollectList(Request, CollectRequestOptions) ([]Collect, error)
	RecurringApplicationChargeList(Request, RecurringApplicationChargeOptons) ([]RecurringApplicationCharge, error)
	ScriptTagCreate(Request, ScriptTag) (ScriptTag, error)
}

type RestAdminClient struct {
	Http    *http.Client
	Logger  *log.Logger
	Version ApiVersion
}

type Request struct {
	ShopName    string
	AccessToken string
	Ctx         context.Context
	Method string
	Url url.URL
	Headers map[string]string
	Resource string
	Body io.Reader
}

func (r *RestAdminClient) Request(request Request) (result []byte, err error) {
	req, err := http.NewRequestWithContext(request.Ctx, request.Method, request.Url.String(), request.Body)
	if err != nil {
		err = errors.WithMessagef(err, "unable to create request with input %s", request)
		return
	}

	req.Header.Set("X-Shopify-Access-Token", request.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	for k, v := range request.Headers {
		req.Header.Set(k, v)
	}

	resp, err := r.Http.Do(req)
	if err != nil {
		err = errors.WithMessagef(err, "request failed %s", request)
		return
	}

	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.WithMessage(err, "failed to read response body")
		return
	}

	r.Logger.Println("received response: ", string(result))

	if resp.StatusCode >= 300 {
		err = errors.New(string(result))
		err = errors.WithMessagef(err, "received %v response", resp.StatusCode)
		return
	}

	return
}