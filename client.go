package shopify

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Client interface {
	OAuthRequest(Request, OAuthRequest) (OAuthResponse, error)

	RecurringApplicationChargeCreate(details Request, request RecurringApplicationCharge) (RecurringApplicationCharge, error)
	RecurringApplicationChargeActivate(Request, RecurringApplicationCharge) (RecurringApplicationCharge, error)

	ShopGet(Request) (Shop, error)

	WebhookCreate(ShopifyContext, Webhook) (Webhook, error)
	WebhookDelete(ShopifyContext, int) error
	WebhookList(ShopifyContext, WebHookRequestOptions) ([]Webhook, error)

	ProductList(Request, ProductRequestOptions) ([]Product, error)

	CollectList(Request, CollectRequestOptions) ([]Collect, error)

	RecurringApplicationChargeList(Request, RecurringApplicationChargeOptons) ([]RecurringApplicationCharge, error)

	ScriptTagCreate(Request, ScriptTag) (ScriptTag, error)
}

type QueryParamStringer interface {
	UrlOptionsString() (queryParams string, err error)
}

type Lister interface {
	GetResourceName() string
}

type Creator interface {
	BuildCreateUrl(Request) string
	GetResourceName() string
}

type Getter interface {
	GetId() int
	GetResourceName() string
	BuildGetUrl(Request) string
}

type RestAdminClient struct {
	Http    *http.Client
	Logger  *log.Logger
	Version ApiVersion
}

type ShopifyContext struct {
	ShopName    string
	AccessToken string
	Ctx         context.Context
}

type Request struct {
	Context ShopifyContext
	Method  string
	Url     string
	Headers map[string]string
	Body    []byte
	Version ApiVersion
}

func (r *RestAdminClient) Request(request Request) (result io.Reader, err error) {
	req, err := http.NewRequestWithContext(request.Context.Ctx, request.Method, request.Url, bytes.NewBuffer(request.Body))
	if err != nil {
		err = errors.WithMessagef(err, "unable to create request with input %+v", request)
		return
	}

	req.Header.Set("X-Shopify-Access-Token", request.Context.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	for k, v := range request.Headers {
		req.Header.Set(k, v)
	}

	resp, err := r.Http.Do(req)
	if err != nil {
		err = errors.WithMessagef(err, "request failed %s", request)
		return
	}

	result = resp.Body

	if resp.StatusCode >= 300 {
		response, _ := ioutil.ReadAll(resp.Body)
		err = errors.New(string(response))
		err = errors.WithMessagef(err, "received %v response", resp.StatusCode)
		return
	}

	return
}

func (r *RestAdminClient) List(context ShopifyContext, options QueryParamStringer, resource Lister) (err error) {
	var request = Request{
		Context: context,
		Method:  "GET",
		Version: r.Version,
	}
	optionString, err := options.UrlOptionsString()
	if err != nil {
		return
	}
	request.Url = BuildSimpleUrl(request, resource.GetResourceName()) + "?" + optionString
	r.Logger.Printf("sending request for the %s with url %s", resource.GetResourceName(), request.Url)

	buf, err := r.Request(request)

	decoder := json.NewDecoder(buf)
	err = decoder.Decode(resource)
	if err != nil {
		return

	}

	return
}

func (r *RestAdminClient) Get(context ShopifyContext, resource Getter) (err error) {
	var request = Request{
		Context: context,
		Method:  "GET",
		Version: r.Version,
	}

	request.Url = resource.BuildGetUrl(request)

	buf, err := r.Request(request)

	decoder := json.NewDecoder(buf)
	err = decoder.Decode(resource)
	if err != nil {
		return

	}

	return
}

func (r *RestAdminClient) Create(context ShopifyContext, returnResource Creator, originalResource Creator) (err error) {
	var request = Request{
		Context: context,
		Method:  "POST",
		Version: r.Version,
	}
	request.Body, err = json.Marshal(originalResource)
	if err != nil {
		err = errors.WithMessage(err, "failure while marshaling the request data")
		return
	}
	request.Url = returnResource.BuildCreateUrl(request)
	buf, err := r.Request(request)
	if err != nil {
		err = errors.WithMessage(err, "there was error while making the request")
		return
	}

	decoder := json.NewDecoder(buf)
	err = decoder.Decode(&returnResource)

	return
}
