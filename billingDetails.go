package shopify

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type RecurringApplicationCharge struct {
	ConfirmationUrl string `json:"confirmation_url,omitempty"`
	Id              int    `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Price           string `json:"price,omitempty"`
	ReturnUrl       string `json:"return_url,omitempty"`
	Status          string `json:"status,omitempty"`
	Test            bool   `json:"test,omitempty"`
	TrialDays       int    `json:"trial_days,omitempty"`
}

type RecurringApplicationChargeWrapper struct {
	RecurringCharge RecurringApplicationCharge `json:"recurring_application_charge,omitempty"`
}

func (client *ShopifyApiImpl) BillingRequest(details ShopifyRequestDetails, request RecurringApplicationCharge) (result RecurringApplicationCharge, err error) {
	requestUrl := "https://" + details.ShopName + "/admin/recurring_application_charges.json"

	log.Printf("Making the recurring application charge request for shop %s using URL %s\n", details.ShopName, requestUrl, request)

	requestStr, err := json.Marshal(RecurringApplicationChargeWrapper{request})
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(requestStr))
	if err != nil {
		return
	}

	req.WithContext(details.Ctx)

	req.Header.Add("X-Shopify-Access-Token", details.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Http.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return
	}

	if resp.StatusCode != 201 {
		log.Println("The billing request response status code is: ", resp.StatusCode)
		return result, errors.New("The response from the server was not the expected status code")
	}

	buf, _ := ioutil.ReadAll(resp.Body)
	log.Println("The response for the recurring billing request is: ", string(buf))
	wrapper := RecurringApplicationChargeWrapper{}
	err = json.Unmarshal(buf, &wrapper)
	if err != nil {
		log.Println("Error unmarshaling the billing response", err.Error())
	}

	log.Println("The result of the unmarshaling: ", wrapper)
	result = wrapper.RecurringCharge

	return
}

func (client *ShopifyApiImpl) ActivateBilling(details ShopifyRequestDetails, request RecurringApplicationCharge) (result RecurringApplicationCharge, err error) {
	if details.AccessToken == "" || details.ShopName == "" {
		err = errors.New("Missing the shop name or the access token from the details object inside the activate billing call.")
		return
	}

	requestUrl := "https://" + details.ShopName + "/admin/recurring_application_charges/" + strconv.Itoa(request.Id) + "/activate.json"

	log.Printf("Requesting to activate recurring application charge with id %s for shop %s using URL %s\n", strconv.Itoa(request.Id), details.ShopName, requestUrl)

	requestStr, err := json.Marshal(RecurringApplicationChargeWrapper{request})
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(requestStr))
	if err != nil {
		return
	}

	req.WithContext(details.Ctx)
	req.Header.Add("X-Shopify-Access-Token", details.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Http.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return
	}

	buf, _ := ioutil.ReadAll(resp.Body)
	log.Println("This is the response recieved from activating the billing: ", string(buf))
	wrapper := RecurringApplicationChargeWrapper{}
	err = json.Unmarshal(buf, &wrapper)
	if err != nil {
		return
	}

	result = wrapper.RecurringCharge

	if resp.StatusCode != 200 {
		err = errors.New("Expected status code to be 201")
		return
	}

	return
}
