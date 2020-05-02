package shopify

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
)

type Client interface {
	OAuthRequest(Request, OAuthRequest) (OAuthResponse, error)

	RecurringApplicationChargeCreate(details Request, request RecurringApplicationCharge) (RecurringApplicationCharge, error)
	RecurringApplicationChargeActivate(Request, RecurringApplicationCharge) (RecurringApplicationCharge, error)

	ShopGet(Request) (Shop, error)

	WebhookCreate(Ctx, Webhook) (Webhook, error)
	WebhookDelete(Ctx, int) error
	WebhookList(Ctx, WebHookRequestOptions) ([]Webhook, error)

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

type Ctx struct {
	ShopName     string
	AccessToken  string
	Ctx          context.Context
	CursorUrl    string
	AutoPaginate bool
}

type Request struct {
	Context Ctx
	Method  string
	Url     string
	Headers map[string]string
	Body    []byte
	Version ApiVersion
}

func (r *RestAdminClient) Request(request Request) (result []byte, next string, err error) {
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
		err = errors.WithMessagef(err, "request failed %v", request)
		return
	}
	next = ExtractNextCursorUrl(resp.Header.Get("Link"))

	defer resp.Body.Close()
	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.WithMessage(err, "there was a problem was reading the body")
	}

	if resp.StatusCode >= 300 {
		err = errors.New(string(result))
		err = errors.WithMessagef(err, "received %v response", resp.StatusCode)
		return
	}

	return
}

func (r *RestAdminClient) List(context Ctx, options QueryParamStringer, resource Lister) (next string, err error) {
	var request = Request{
		Context: context,
		Method:  "GET",
		Version: r.Version,
	}
	optionString, err := options.UrlOptionsString()
	if err != nil {
		return
	}
	if context.CursorUrl != "" {
		request.Url = context.CursorUrl
	} else {
		request.Url = BuildSimpleUrl(request, resource.GetResourceName()) + "?" + optionString
	}

	buf, next, err := r.Request(request)

	err = json.Unmarshal(buf, &resource)
	if err != nil {
		fmt.Println(string(buf))
		err = errors.WithMessage(err, "error while unmarshalling list response")
		return
	}

	if context.AutoPaginate {
		if next != "" {
			context.CursorUrl = next
		} else {
			return
		}

		_, err = r.List(context, options, resource)
		if err != nil {
			err = errors.WithMessage(err, "failure during pagination...aborting")
			return
		}
	}

	return
}

func (r *RestAdminClient) Get(context Ctx, resource Getter) (err error) {
	var request = Request{
		Context: context,
		Method:  "GET",
		Version: r.Version,
	}

	request.Url = resource.BuildGetUrl(request)

	buf, _, err := r.Request(request)

	err = json.Unmarshal(buf, &resource)
	if err != nil {
		err = errors.WithMessage(err, "error while unmarshalling get request")
		return
	}

	return
}

func (r *RestAdminClient) Create(context Ctx, returnResource Creator, originalResource Creator) (err error) {
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
	buf, _, err := r.Request(request)
	if err != nil {
		err = errors.WithMessage(err, "there was error while making the request")
		return
	}

	err = json.Unmarshal(buf, &returnResource)
	if err != nil {
		err = errors.WithMessage(err, "error unmarshaling request")
	}

	return
}

func (r *RestAdminClient) Delete(context Ctx, resource string, id int) (err error) {
	var request = Request{
		Context: context,
		Method:  "DELETE",
	}
	request.Url = BuildIdUrl(request, resource, id)

	_, _, err = r.Request(request)
	if err != nil {
		err = errors.WithMessagef(err, "unable to delete webhook %v", id)
	}

	return
}
