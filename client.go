package shopify

import (
	"context"
	"net/http"
)

type ShopifyClient interface {
	OAuthRequest(details ShopifyRequestDetails, request OAuthRequest) (result OAuthResponse, err error)
	BillingRequest(details ShopifyRequestDetails, request RecurringApplicationCharge) (result RecurringApplicationCharge, err error)
	ActivateBilling(details ShopifyRequestDetails, request RecurringApplicationCharge) (result RecurringApplicationCharge, err error)
	GetShopDetails(details ShopifyRequestDetails) (result Shop, err error)
	CreateWebhook(details ShopifyRequestDetails, request Webhook) (result Webhook, err error)
	DeleteWebhook(details ShopifyRequestDetails, request Webhook) (err error)
	GetWebhooks(details ShopifyRequestDetails, options WebHookRequestOptions) (webhooks []Webhook, err error)
	GetProducts(details ShopifyRequestDetails, options ProductRequestOptions) (products []Product, err error)
	GetCollects(details ShopifyRequestDetails, options CollectRequestOptions) (result []Collect, err error)
}

type ShopifyApiImpl struct {
	Http *http.Client
}

type ShopifyRequestDetails struct {
	ShopName    string
	AccessToken string
	Ctx         context.Context
}
