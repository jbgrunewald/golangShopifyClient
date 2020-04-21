package shopify

import (
	"context"
	"log"
	"net/http"
)

type Client interface {
	OAuthRequest(details ShopifyRequestDetails, request OAuthRequest) (result OAuthResponse, err error)
	RecurringApplicationChargeCreate(details ShopifyRequestDetails, request RecurringApplicationCharge) (result RecurringApplicationCharge, err error)
	RecurringApplicationChargeActivate(details ShopifyRequestDetails, request RecurringApplicationCharge) (result RecurringApplicationCharge, err error)
	ShopGet(details ShopifyRequestDetails) (result Shop, err error)
	WebhookCreate(details ShopifyRequestDetails, request Webhook) (result Webhook, err error)
	WebhookDelete(details ShopifyRequestDetails, request Webhook) (err error)
	WebhookList(details ShopifyRequestDetails, options WebHookRequestOptions) (webhooks []Webhook, err error)
	ProductList(details ShopifyRequestDetails, options ProductRequestOptions) (products []Product, err error)
	CollectList(details ShopifyRequestDetails, options CollectRequestOptions) (result []Collect, err error)
	RecurringApplicationChargeList(details ShopifyRequestDetails, options RecurringApplicationChargeOptons) (charges []RecurringApplicationCharge, err error)
	ScriptTagCreate(details ShopifyRequestDetails, scriptTag ScriptTag) (result ScriptTag, err error)
}


type ShopifyApiImpl struct {
	Http    *http.Client
	Logger  *log.Logger
	Version string
}

type ShopifyRequestDetails struct {
	ShopName    string
	AccessToken string
	Ctx         context.Context
}