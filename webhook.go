package shopify

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
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

func (w WebhooksWrapper) NewWrapper() WebhooksWrapper {
	return WebhooksWrapper{}
}

func (w WebhooksWrapper) GetCount() int {
	return len(w.Webhooks)
}

func (w WebhooksWrapper) GetLastId() int {
	return w.Webhooks[len(w.Webhooks)-1].Id
}

func (w WebhooksWrapper) GetResourceName() string {
	return "webhooks"
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
	SinceId      int      `url:"since_id"`
	Topic        string   `url:"topic,omitempty"`
	UpdatedAtMin string   `url:"updated_at_min,omitempty"`
	UpdatedAtMax string   `url:"updated_at_max,omitempty"`
}

func (w WebHookRequestOptions) UrlOptionsString() (queryParams string, err error) {
	values, err := query.Values(w)
	if err != nil {
		err = errors.WithMessagef(err, "unable to encode options as query param", w)
		return
	}

	queryParams = values.Encode()
	return
}

func (r *RestAdminClient) WebhookCreate(context ShopifyContext, resource Webhook) (result Webhook, err error) {
	var request = Request{
		Context: context,
		Method:  "POST",
	}
	request.Url = BuildSimpleUrl(request, "webhooks")

	r.logger.Printf("create webhook for topic %s for shop %s", resource.Topic, context.ShopName)

	request.Body, err = json.Marshal(WebhookWrapper{resource})
	if err != nil {
		err = errors.WithMessage(err, "failure while marshaling the request body")
		return
	}

	buf, err := r.Request(request)
	if err != nil {
		err = errors.WithMessage(err, "failure making request")
		return
	}

	wrapper := WebhookWrapper{}

	err = json.Unmarshal(buf, &wrapper)
	if err != nil {
		err = errors.WithMessage(err, "failure while unmarshaling the response")
		return
	}

	result = wrapper.Webhook

	return
}

func (r *RestAdminClient) WebhookDelete(context ShopifyContext, id int) (err error) {
	var request = Request{
		Context: context,
		Method:  "DELETE",
	}
	request.Url = BuildIdUrl(request, "webhooks", id)
	r.logger.Printf("requesting to delete webhook with id %s", id)

	_, err = r.Request(request)
	if err != nil {
		err = errors.WithMessagef(err, "unable to delete webhook %s", id)
	}

	return
}

func (r *RestAdminClient) WebhookList(context ShopifyContext, options WebHookRequestOptions) (results []Webhook, err error) {
	var wrapper = &WebhooksWrapper{}
	err = r.List(context, options, wrapper)
	results = wrapper.Webhooks

	//TODO figure out a way to generalize the autopaginate logic
	//if context.AutoPaginate {
	//	//TODO add support for curser based pagination
	//	if len(results) < options.Limit || (options.Limit == 0 && len(results) < 50) {
	//		return
	//	}
	//
	//	options.SinceId = results[len(results)-1].Id
	//	nextResult, err := r.WebhookList(context, options)
	//	if err != nil {
	//		err = errors.WithMessage(err, "failure during pagination...aborting")
	//		results = []Webhook{}
	//		return
	//	}
	//	results = append(results, nextResult...)
	//	return
	//}

	return
}
