package shopify

import (
	"errors"
)

type ShopifyTestImpl struct {
	fakeOAuthResponses              map[string]OAuthResponse
	fakeShopResponses               map[string]Shop
	fakeBillingSetupResponses       map[string]RecurringApplicationCharge
	fakeBillingActivateResponses    map[string]RecurringApplicationCharge
	fakeCreateWebhookResponses      map[string]Webhook
	fakeGetWebhookResponse          map[string][]Webhook
	fakeProductsResponses           map[string][]Product
	fakeCollectsResponses           map[string][]Collect
	fakeRecurringApplicationCharges map[string][]RecurringApplicationCharge
	fakeScriptTag                   map[string]ScriptTag
}

func NewShopifyTestImp() *ShopifyTestImpl {
	imp := ShopifyTestImpl{}
	imp.fakeOAuthResponses = make(map[string]OAuthResponse)
	imp.fakeShopResponses = make(map[string]Shop)
	imp.fakeBillingSetupResponses = make(map[string]RecurringApplicationCharge)
	imp.fakeBillingActivateResponses = make(map[string]RecurringApplicationCharge)
	imp.fakeCreateWebhookResponses = make(map[string]Webhook)
	imp.fakeProductsResponses = make(map[string][]Product)
	imp.fakeRecurringApplicationCharges = make(map[string][]RecurringApplicationCharge)
	return &imp
}

func (client *ShopifyTestImpl) OAuthRequest(details ShopifyContext, request OAuthRequest) (result OAuthResponse, err error) {
	result, ok := client.fakeOAuthResponses[details.ShopName]
	if !ok {
		err = errors.New("something has gone wrong with oAuth")
		return
	}

	return
}

func (client *ShopifyTestImpl) RecurringApplicationChargeCreate(details ShopifyContext, request RecurringApplicationCharge) (result RecurringApplicationCharge, err error) {
	result, ok := client.fakeBillingSetupResponses[details.ShopName]
	if !ok {
		err = errors.New("something has gone wrong with setting up the billing")
	}

	return
}

func (client *ShopifyTestImpl) RecurringApplicationChargeActivate(details ShopifyContext, request RecurringApplicationCharge) (result RecurringApplicationCharge, err error) {
	result, ok := client.fakeBillingActivateResponses[details.ShopName]
	if !ok {
		err = errors.New("something has gone wrong with activate billing")
		return
	}

	return
}

func (client *ShopifyTestImpl) ShopGet(details ShopifyContext) (result Shop, err error) {
	result, ok := client.fakeShopResponses[details.ShopName]
	if !ok {
		err = errors.New("something has gone wrong with shop details")
	}

	return
}

func (client *ShopifyTestImpl) WebhookCreate(details ShopifyContext, request Webhook) (result Webhook, err error) {
	result, ok := client.fakeCreateWebhookResponses[details.ShopName]
	if !ok {
		err = errors.New("something has gone wrong with requesting the webhook")
	}

	return
}

func (client *ShopifyTestImpl) WebhookDelete(details ShopifyContext, request Webhook) (err error) {
	delete(client.fakeCreateWebhookResponses, details.ShopName)
	return
}

func (client *ShopifyTestImpl) WebhookList(details ShopifyContext, options WebHookRequestOptions) (result []Webhook, err error) {
	result, ok := client.fakeGetWebhookResponse[details.ShopName]
	if !ok {
		err = errors.New("something has gone wrong requesting the webhooks")
	}

	return
}

func (client *ShopifyTestImpl) ProductList(details ShopifyContext, options ProductRequestOptions) (result []Product, err error) {
	result, ok := client.fakeProductsResponses[details.ShopName]
	if !ok {
		err = errors.New("something has gone wrong with requesting the products")
	}

	return
}

func (client *ShopifyTestImpl) CollectList(details ShopifyContext, options CollectRequestOptions) (result []Collect, err error) {
	result, ok := client.fakeCollectsResponses[details.ShopName]
	if !ok {
		err = errors.New("there was an issue with requesting the collects")
	}

	return
}

func (client *ShopifyTestImpl) RecurringApplicationChargeList(details ShopifyContext, options RecurringApplicationChargeOptons) (result []RecurringApplicationCharge, err error) {
	result, ok := client.fakeRecurringApplicationCharges[details.ShopName]
	if !ok {
		err = errors.New("there was an issue with getting the reccurring application charges")
	}

	return
}

func (c *ShopifyTestImpl) ScriptTagCreate(details ShopifyContext, scriptTag ScriptTag) (result ScriptTag, err error) {
	result, ok := c.fakeScriptTag[details.ShopName]
	if !ok {
		err = errors.New("there was an issue with getting the script tag")
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

func (client *ShopifyTestImpl) RegisterFakeRecurringApplicationCharges(shopName string, charges []RecurringApplicationCharge) {
	client.fakeRecurringApplicationCharges[shopName] = charges
}

func (client *ShopifyTestImpl) RegisterFakeScriptTag(shopName string, response ScriptTag) {
	client.fakeScriptTag[shopName] = response
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
