package shopify

func main() {
	var client = &RestAdminClient{}

	details := ShopifyContext{
		Password: "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",
		ApiKey:   "79752807306004d647f91d254eaa8a60",
	}

	webhook := Webhook{}
	result, err := client.WebhookCreate(details, webhook)
	println(result, err)
}
