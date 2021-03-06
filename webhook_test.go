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

func TestWebhookList(t *testing.T) {
	var listResponse = WebhooksWrapper{
		Webhooks: make([]Webhook, 0),
	}
	listResponse.Webhooks = append(listResponse.Webhooks, *createResponse.Webhook)
	server := httptest.NewTLSServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("X-Shopify-Access-Token")
		if req.URL.Path == "/admin/api/2019-07/webhooks.json" && req.Method == "GET" && req.URL.RawQuery == "limit=65&since_id=0" && authHeader == "thisisatoken" {
			response, _ := json.Marshal(listResponse)
			_, _ = rw.Write(response)
			return
		}

		rw.Write([]byte("something went wrong"))
	}))
	defer server.Close()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	client := RestAdminClient{
		Http:    server.Client(),
		Logger:  logger,
		Version: VERSION_2019_07,
	}

	serverUrl, _ := url.Parse(server.URL)
	requestContext := Ctx{
		AccessToken: "thisisatoken",
		ShopName:    serverUrl.Host,
		Ctx:         context.Background(),
	}

	options := WebHookRequestOptions{
		Limit: 65,
	}
	result, _, err := client.WebhookList(requestContext, options)
	if err != nil {
		t.Error(err, result)
	}
	resultStringed, err := json.Marshal(result)
	if err != nil {
		t.Error(err)
	}
	expected, _ := json.Marshal(listResponse.Webhooks)
	if bytes.Compare(resultStringed, expected) != 0 {
		t.Error()
	}
}

func TestWebhookListAutoPaginate(t *testing.T) {
	var listResponse = WebhooksWrapper{
		Webhooks: make([]Webhook, 0),
	}
	listResponse.Webhooks = append(listResponse.Webhooks, *createResponse.Webhook)
	server := httptest.NewTLSServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("X-Shopify-Access-Token")
		response, _ := json.Marshal(listResponse)
		if req.URL.Path == "/admin/api/2019-07/webhooks.json" && req.Method == "GET" && req.URL.RawQuery == "limit=65&since_id=0" && authHeader == "thisisatoken" {
			rw.Header().Set("Link", "<https://"+req.Host+"/admin/api/2019-07/webhooks.json?page_info=hijgklmn&limit=3>; rel=next")
			_, _ = rw.Write(response)
			return
		}

		if req.URL.Path == "/admin/api/2019-07/webhooks.json" && req.Method == "GET" && req.URL.RawQuery == "page_info=hijgklmn&limit=3" && authHeader == "thisisatoken" {
			_, _ = rw.Write(response)
			return
		}

		fmt.Println("something's not working")
		rw.Write([]byte("something went wrong"))
	}))
	defer server.Close()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	client := RestAdminClient{
		Http:    server.Client(),
		Logger:  logger,
		Version: VERSION_2019_07,
	}

	serverUrl, _ := url.Parse(server.URL)
	requestContext := Ctx{
		AccessToken:  "thisisatoken",
		ShopName:     serverUrl.Host,
		Ctx:          context.Background(),
		AutoPaginate: true,
	}

	options := WebHookRequestOptions{
		Limit: 65,
	}
	result, _, err := client.WebhookList(requestContext, options)
	if err != nil {
		t.Error(err, result)
	}
	resultStringed, err := json.Marshal(result)
	if err != nil {
		t.Error(err)
	}
	listResponse.Webhooks = append(listResponse.Webhooks, *createResponse.Webhook)
	expected, _ := json.Marshal(listResponse.Webhooks)
	if bytes.Compare(resultStringed, expected) != 0 {
		t.Error()
	}
}

func TestWebhookListWithCursorUrl(t *testing.T) {
	var listResponse = WebhooksWrapper{
		Webhooks: make([]Webhook, 0),
	}
	listResponse.Webhooks = append(listResponse.Webhooks, *createResponse.Webhook)
	server := httptest.NewTLSServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("X-Shopify-Access-Token")
		if req.URL.Path == "/admin/api/2019-07/webhooks.json" && req.Method == "GET" && req.URL.RawQuery == "limit=65&since_id=0" && authHeader == "thisisatoken" {
			response, _ := json.Marshal(listResponse)
			_, _ = rw.Write(response)
			return
		}

		rw.Write([]byte("something went wrong"))
	}))
	defer server.Close()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	client := RestAdminClient{
		Http:    server.Client(),
		Logger:  logger,
		Version: VERSION_2019_07,
	}

	serverUrl, _ := url.Parse(server.URL)
	requestContext := Ctx{
		AccessToken: "thisisatoken",
		ShopName:    serverUrl.Host,
		Ctx:         context.Background(),
	}

	options := WebHookRequestOptions{
		Limit: 65,
	}
	result, _, err := client.WebhookList(requestContext, options)
	if err != nil {
		t.Error(err, result)
	}
	resultStringed, err := json.Marshal(result)
	if err != nil {
		t.Error(err)
	}
	expected, _ := json.Marshal(listResponse.Webhooks)
	if bytes.Compare(resultStringed, expected) != 0 {
		fmt.Println(string(resultStringed))
		fmt.Println(string(expected))
		t.Error()
	}
}

func TestWebhookCreate(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("X-Shopify-Access-Token")
		if req.URL.Path == "/admin/api/2019-07/webhooks.json" && req.Method == "POST" && authHeader == "thisisatoken" {
			response, _ := json.Marshal(createResponse)
			_, _ = rw.Write(response)
			return
		}

		rw.Write([]byte("something went wrong"))
	}))
	defer server.Close()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	client := RestAdminClient{
		Http:    server.Client(),
		Logger:  logger,
		Version: VERSION_2019_07,
	}

	serverUrl, _ := url.Parse(server.URL)
	requestContext := Ctx{
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

func TestWebhookDelete(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("X-Shopify-Access-Token")
		if req.URL.Path == "/admin/api/2019-07/webhooks/12345.json" && req.Method == "DELETE" && authHeader == "thisisatoken" {
			response, _ := json.Marshal(createResponse)
			_, _ = rw.Write(response)
			return
		}

		rw.Write([]byte("something went wrong"))
	}))
	defer server.Close()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	client := RestAdminClient{
		Http:    server.Client(),
		Logger:  logger,
		Version: VERSION_2019_07,
	}

	serverUrl, _ := url.Parse(server.URL)
	requestContext := Ctx{
		AccessToken: "thisisatoken",
		ShopName:    serverUrl.Host,
		Ctx:         context.Background(),
	}

	err := client.WebhookDelete(requestContext, 12345)
	if err != nil {
		t.Error(err)
	}
}
