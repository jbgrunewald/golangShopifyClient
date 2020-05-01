package shopify

import (
	"encoding/json"
	"fmt"
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

type WebhookWrapper struct {
	Webhook *Webhook `json:"webhook"`
}

func (w WebhookWrapper) GetResourceName() string {
	return "webhooks"
}

func (w WebhookWrapper) BuildCreateUrl(request Request) string {
	return BuildSimpleUrl(request, w.GetResourceName())
}

type Webhooks struct {
	Webhooks []Webhook `json:"webhooks"`
}

type WebhooksWrapper struct {
	Webhooks []Webhook
}

func (w *WebhooksWrapper) UnmarshalJSON(data []byte) (err error) {
	var wrapper Webhooks
	if err = json.Unmarshal(data, &wrapper); err != nil {
		fmt.Println(err)
		return
	}

	w.Webhooks = append(w.Webhooks, wrapper.Webhooks...)
	return
}

func (w WebhooksWrapper) GetResourceName() string {
	return "webhooks"
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
		err = errors.WithMessagef(err, "unable to encode options as query param %v", w)
		return
	}

	queryParams = values.Encode()
	return
}

func (r *RestAdminClient) WebhookCreate(context ShopifyContext, request Webhook) (result *Webhook, err error) {
	var returnWrapper = new(WebhookWrapper)
	requestWrapper := WebhookWrapper{Webhook: &request}
	err = r.Create(context, returnWrapper, requestWrapper)
	result = returnWrapper.Webhook

	return
}

func (r *RestAdminClient) WebhookDelete(context ShopifyContext, id int) (err error) {
	err = r.Delete(context, "webhooks", id)
	if err != nil {
		err = errors.WithMessagef(err, "unable to delete webhook %v", id)
	}

	return
}

func (r *RestAdminClient) WebhookList(context ShopifyContext, options WebHookRequestOptions) (results []Webhook, next string, err error) {
	var wrapper = &WebhooksWrapper{}
	next, err = r.List(context, options, wrapper)
	results = wrapper.Webhooks
	return
}
