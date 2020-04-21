package shopify

import (
	"context"
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
	Version apiVersion
}

type RequestDetails struct {
	ShopName    string
	AccessToken string
	Ctx         context.Context
}
