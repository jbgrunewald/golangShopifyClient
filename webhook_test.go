package shopify

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

var createResponse = WebhookWrapper{
	Webhook: &Webhook{
		Address:    "https://8aea8006.ngrok.io/",
		Created_at: "2020-04-26T22:41:14-07:00",
		Format:     "json",
		Id:         750230437990,
		Topic:      "products/create",
		UpdatedAt:  "2020-04-26T22:41:14-07:00",
	},
}

func TestWebhookCreate(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/admin/api/2019-07/webhooks.json" {
			rw.Write([]byte("something went wrong"))
		}

		response, _ := json.Marshal(createResponse)
		_, _ = rw.Write(response)
	}))
	defer server.Close()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	client := RestAdminClient{
		Http:    server.Client(),
		Logger:  logger,
		Version: VERSION_2019_07,
	}

	serverUrl, _ := url.Parse(server.URL)
	requestContext := ShopifyContext{
		AccessToken: "thisisatoken",
		ShopName:    serverUrl.Host,
		Ctx:         context.Background(),
	}

	webhook := Webhook{
		Address: "https://8aea8006.ngrok.io",
		Topic:   ProductCreate,
		Format:  "json",
	}
	result, err := client.WebhookCreate(requestContext, webhook)
	fmt.Printf("%+v\n", result)
	if err != nil {
		t.Error(err)
	}
	resultStringed, err := json.Marshal(result)
	if err != nil {
		t.Error(err)
	}
	expected, _ := json.Marshal(createResponse.Webhook)
	if bytes.Compare(resultStringed, expected) != 0 {
		t.Error()
	}
}
