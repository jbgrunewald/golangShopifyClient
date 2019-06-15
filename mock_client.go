package shopify

import (
	"errors"
)

type ShopifyTestImpl struct {
	fakeOAuthResponses           map[string]OAuthResponse
	fakeShopResponses            map[string]Shop
	fakeBillingSetupResponses    map[string]RecurringApplicationCharge
	fakeBillingActivateResponses map[string]RecurringApplicationCharge
	fakeCreateWebhookResponses   map[string]Webhook
	fakeGetWebhookResponse       map[string][]Webhook
	fakeProductsResponses        map[string][]Product
	fakeCollectsResponses        map[string][]Collect
}

func NewShopifyTestImp() *ShopifyTestImpl {
	imp := ShopifyTestImpl{}
	imp.fakeOAuthResponses = make(map[string]OAuthResponse)
	imp.fakeShopResponses = make(map[string]Shop)
	imp.fakeBillingSetupResponses = make(map[string]RecurringApplicationCharge)
	imp.fakeBillingActivateResponses = make(map[string]RecurringApplicationCharge)
	imp.fakeCreateWebhookResponses = make(map[string]Webhook)
	imp.fakeProductsResponses = make(map[string][]Product)
	return &imp
}

func (client *ShopifyTestImpl) OAuthRequest(details ShopifyRequestDetails, request OAuthRequest) (result OAuthResponse, err error) {

	result, ok := client.fakeOAuthResponses[details.ShopName]
	if !ok {
		err = errors.New("something has gone wrong with oAuth")
		return
	}

	return
}

func (client *ShopifyTestImpl) BillingRequest(details ShopifyRequestDetails, request RecurringApplicationCharge) (result RecurringApplicationCharge, err error) {

	result, ok := client.fakeBillingSetupResponses[details.ShopName]
	if !ok {
		err = errors.New("something has gone wrong with setting up the billing")
	}

	return
}

func (client *ShopifyTestImpl) ActivateBilling(details ShopifyRequestDetails, request RecurringApplicationCharge) (result RecurringApplicationCharge, err error) {

	result, ok := client.fakeBillingActivateResponses[details.ShopName]
	if !ok {
		err = errors.New("something has gone wrong with activate billing")
		return
	}

	return
}

func (client *ShopifyTestImpl) GetShopDetails(details ShopifyRequestDetails) (result Shop, err error) {
	result, ok := client.fakeShopResponses[details.ShopName]
	if !ok {
		err = errors.New("something has gone wrong with shop details")
	}

	return
}

func (client *ShopifyTestImpl) CreateWebhook(details ShopifyRequestDetails, request Webhook) (result Webhook, err error) {
	result, ok := client.fakeCreateWebhookResponses[details.ShopName]
	if !ok {
		err = errors.New("something has gone wrong with requesting the webhook")
	}

	return
}

func (client *ShopifyTestImpl) DeleteWebhook(details ShopifyRequestDetails, request Webhook) (err error) {
	delete(client.fakeCreateWebhookResponses, details.ShopName)
	return
}

func (client *ShopifyTestImpl) GetWebhooks(details ShopifyRequestDetails, options WebHookRequestOptions) (result []Webhook, err error) {
	result, ok := client.fakeGetWebhookResponse[details.ShopName]
	if !ok {
		err = errors.New("something has gone wrong requesting the webhooks")
	}

	return
}

func (client *ShopifyTestImpl) GetProducts(details ShopifyRequestDetails, options ProductRequestOptions) (result []Product, err error) {
	result, ok := client.fakeProductsResponses[details.ShopName]
	if !ok {
		err = errors.New("something has gone wrong with requesting the products")
	}

	return
}

func (client *ShopifyTestImpl) GetCollects(details ShopifyRequestDetails, options CollectRequestOptions) (result []Collect, err error) {
	result, ok := client.fakeCollectsResponses[details.ShopName]
	if !ok {
		err = errors.New("there was an issue with requesting the collects")
	}

	return
}

func (client *ShopifyTestImpl) RegisterOAuthResponse(shopName string, response OAuthResponse) {
	client.fakeOAuthResponses[shopName] = response
}

func (client *ShopifyTestImpl) RegisterShopResponse(shopName string, response Shop) {
	client.fakeShopResponses[shopName] = response
}

func (client *ShopifyTestImpl) RegisterBillingSetupResponse(shopName string, response RecurringApplicationCharge) {
	client.fakeBillingSetupResponses[shopName] = response
}

func (client *ShopifyTestImpl) RegisterFakeBillingActivateResponse(shopName string, response RecurringApplicationCharge) {
	client.fakeBillingActivateResponses[shopName] = response
}

func (client *ShopifyTestImpl) RegisterWebhookResponse(shopName string, response Webhook) {
	client.fakeCreateWebhookResponses[shopName] = response
}

func (client *ShopifyTestImpl) RegisterFakeProductsResponse(shopName string, products []Product) {
	client.fakeProductsResponses[shopName] = products
}

func (client *ShopifyTestImpl) ClearShopResponses() {
	client.fakeShopResponses = make(map[string]Shop)
}

func (client *ShopifyTestImpl) ClearBillingSetupResponses() {
	client.fakeBillingSetupResponses = make(map[string]RecurringApplicationCharge)
}

func (client *ShopifyTestImpl) ClearBillingActivateResponses() {
	client.fakeBillingActivateResponses = make(map[string]RecurringApplicationCharge)
}

func (client *ShopifyTestImpl) ClearOAuthResponses() {
	client.fakeOAuthResponses = make(map[string]OAuthResponse)
}

func (client *ShopifyTestImpl) ClearWebhookResponse() {
	client.fakeCreateWebhookResponses = make(map[string]Webhook)
}