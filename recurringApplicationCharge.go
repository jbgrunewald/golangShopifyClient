package shopify

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/google/go-querystring/query"
	"io/ioutil"
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

type RecurringApplicationChargeOptons struct {
	SinceId int    `json:"since_id, omitempty"`
	Fields  string `json:"fields, omitempty"`
	All     bool
}

type RecurringApplicationChargeWrapper struct {
	RecurringApplicationCharges RecurringApplicationCharge `json:"recurring_application_charge,omitempty"`
}

type RecurringApplicationChargesWrapper struct {
	RecurringApplicationCharges []RecurringApplicationCharge `json:"recurring_application_charges,omitempty"`
}

func (c *ShopifyApiImpl) RecurringApplicationChargeCreate(details ShopifyRequestDetails, request RecurringApplicationCharge) (result RecurringApplicationCharge, err error) {
	requestUrl := "https://" + details.ShopName + "/admin/recurring_application_charges.json"

	c.Logger.Printf("Making the recurring application charge request for shop %s using URL %s\n", details.ShopName, requestUrl)

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

	resp, err := c.Http.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return
	}

	if resp.StatusCode != 201 {
		c.Logger.Println("The billing request response status code is: ", resp.StatusCode)
		return result, errors.New("The response from the server was not the expected status code")
	}

	buf, _ := ioutil.ReadAll(resp.Body)
	c.Logger.Println("The response for the recurring billing request is: ", string(buf))
	wrapper := RecurringApplicationChargeWrapper{}
	err = json.Unmarshal(buf, &wrapper)
	if err != nil {
		c.Logger.Println("Error unmarshaling the billing response", err.Error())
	}

	c.Logger.Println("The result of the unmarshaling: ", wrapper)
	result = wrapper.RecurringApplicationCharges

	return
}

func (c *ShopifyApiImpl) RecurringApplicationChargeActivate(details ShopifyRequestDetails, request RecurringApplicationCharge) (result RecurringApplicationCharge, err error) {
	if details.AccessToken == "" || details.ShopName == "" {
		err = errors.New("Missing the shop name or the access token from the details object inside the activate billing call.")
		return
	}

	requestUrl := "https://" + details.ShopName + "/admin/recurring_application_charges/" + strconv.Itoa(request.Id) + "/activate.json"

	c.Logger.Printf("Requesting to activate recurring application charge with id %s for shop %s using URL %s\n", strconv.Itoa(request.Id), details.ShopName, requestUrl)

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

	resp, err := c.Http.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return
	}

	buf, _ := ioutil.ReadAll(resp.Body)
	c.Logger.Println("This is the response recieved from activating the billing: ", string(buf))
	wrapper := RecurringApplicationChargeWrapper{}
	err = json.Unmarshal(buf, &wrapper)
	if err != nil {
		return
	}

	result = wrapper.RecurringApplicationCharges

	if resp.StatusCode != 200 {
		err = errors.New("Expected status code to be 201")
		return
	}

	return
}

func (c *ShopifyApiImpl) RecurringApplicationChargeList(details ShopifyRequestDetails, options RecurringApplicationChargeOptons) (charges []RecurringApplicationCharge, err error) {
	v, err := query.Values(options)
	if err != nil {
		c.Logger.Println("there's an issue setting up the query params while request the recurring application charges")
		return
	}
	requestUrl := "https://" + details.ShopName + "/admin/" + c.Version + "recurring_application_charges.json?" + v.Encode()
	c.Logger.Println(requestUrl)

	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return
	}

	req.Header.Add("X-Shopify-Access-Token", details.AccessToken)

	resp, err := c.Http.Do(req)
	if err != nil {
		return
	}

	buf, _ := ioutil.ReadAll(resp.Body)
	c.Logger.Println("This is the response for the recurring application charge: ", string(buf))
	wrapper := RecurringApplicationChargesWrapper{}
	err = json.Unmarshal(buf, &wrapper)
	if err != nil {
		return
	}

	charges = wrapper.RecurringApplicationCharges

	if len(charges) == 0 || !options.All {
		return
	}

	options.SinceId = charges[len(charges)-1].Id
	nextResult, err := c.RecurringApplicationChargeList(details, options)

	charges = append(charges, nextResult...)

	return
}
