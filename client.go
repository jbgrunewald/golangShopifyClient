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
	"strings"
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
	NewWrapper() WebhooksWrapper
	GetCount() int
	GetResourceName() string
}

type RestAdminClient struct {
	Http    *http.Client
	Logger  *log.Logger
	Version ApiVersion
}

type ShopifyContext struct {
	ShopName     string
	AccessToken  string
	Ctx          context.Context
	AutoPaginate bool
}

type Request struct {
	Context ShopifyContext
	Method  string
	Url     string
	Headers map[string]string
	Body    []byte
}

func (r *RestAdminClient) Request(request Request) (result io.Reader, err error) {
	req, err := http.NewRequestWithContext(request.Context.Ctx, request.Method, request.Url, bytes.NewBuffer(request.Body))
	if err != nil {
		err = errors.WithMessagef(err, "unable to create request with input %s", request)
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

/*
A helper for building the base url for a shopify request
*/
func (r *RestAdminClient) BuildBaseUrl(request Request) (url string) {
	pathBuilder := strings.Builder{}
	pathBuilder.WriteString("https://")
	pathBuilder.WriteString(request.Context.ShopName)
	pathBuilder.WriteString("/admin/")
	pathBuilder.WriteString(r.Version.String())
	pathBuilder.WriteString("/")
	url = pathBuilder.String()
	return
}

/*
A Helper function to build a request url with only a resource component.
*/
func (r *RestAdminClient) BuildSimpleUrl(request Request, resource string) (url string) {
	pathBuilder := strings.Builder{}
	pathBuilder.WriteString(r.BuildBaseUrl(request))
	pathBuilder.WriteString("/")
	pathBuilder.WriteString(resource)
	pathBuilder.WriteString(".json")
	url = pathBuilder.String()
	return
}

/*
A helper for building a url that has an id
*/
func (r *RestAdminClient) BuildIdUrl(request Request, resource string, id int) (url string) {
	pathBuilder := strings.Builder{}
	pathBuilder.WriteString(r.BuildBaseUrl(request))
	pathBuilder.WriteString("/")
	pathBuilder.WriteString(resource)
	pathBuilder.WriteString(".json")
	url = pathBuilder.String()
	return
}

func (r *RestAdminClient) List(context ShopifyContext, options QueryParamStringer, resource Lister) (err error) {
	var request = Request{
		Context: context,
		Method:  "GET",
	}
	optionString, err := options.UrlOptionsString()
	if err != nil {
		return
	}
	request.Url = r.BuildSimpleUrl(request, resource.GetResourceName()) + "?" + optionString
	r.Logger.Println("sending request for the webhooks", request.Url)

	buf, err := r.Request(request)

	decoder := json.NewDecoder(buf)
	err = decoder.Decode(resource)
	if err != nil {
		return

	}

	return
}