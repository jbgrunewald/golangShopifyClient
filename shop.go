package shopify

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Shop struct {
	AddressOne                     string   `json:"address1,omitempty"`
	AddressTwo                     string   `json:"address2,omitempty"`
	CheckoutApiSupported           bool     `json:"checkout_api_supported,omitempty"`
	City                           string   `json:"city,omitempty"`
	Country                        string   `json:"country,omitempty"`
	CountryCode                    string   `json:"country_code,omitempty"`
	CountryTaxes                   bool     `json:"country_taxes,omitempty"`
	CreateAt                       string   `json:"create_at,omitempty"`
	CustomerEmail                  string   `json:"customer_email,omitempty"`
	Currency                       string   `json:"currency,omitempty"`
	Domain                         string   `json:"domain,omitempty"`
	EnabledPresentmenCurrencies    []string `json:"enabled_presentmen_currencies,omitempty"`
	EligibleForCardReaderGiveaway  bool     `json:"eligible_for_card_reader_giveaway,omitempty"`
	EligibleForPayments            bool     `json:"eligible_for_payments,omitempty"`
	Email                          string   `json:"email,omitempty"`
	ForceSsl                       bool     `json:"force_ssl,omitempty"`
	GoogleAppsDomain               string   `json:"google_apps_domain,omitempty"`
	GoogleAppsLoginEnabled         bool     `json:"google_apps_login_enabled,omitempty"`
	HasDiscounts                   bool     `json:"has_discounts,omitempty"`
	HasGiftCards                   bool     `json:"has_gift_cards,omitempty"`
	HasStorefront                  bool     `json:"has_storefront,omitempty"`
	IanaTimezone                   string   `json:"iana_timezone,omitempty"`
	Id                             int      `json:"id,omitempty"`
	Latitude                       float64  `json:"latitude,omitempty"`
	Longitude                      float64  `json:"longitude,omitempty"`
	MoneyFormat                    string   `json:"money_format,omitempty"`
	MoneyInEmailsFormat            string   `json:"money_in_emails_format,omitempty"`
	MoneyWithCurrencyFormat        string   `json:"money_with_currency_format,omitempty"`
	MultiLocationEnabled           bool     `json:"multi_location_enabled,omitempty"`
	MyshopifyDomain                string   `json:"myshopify_domain,omitempty"`
	Name                           string   `json:"name,omitempty"`
	PasswordEnabled                bool     `json:"password_enabled,omitempty"`
	Phone                          string   `json:"phone,omitempty"`
	PlanDisplayName                string   `json:"plan_display_name,omitempty"`
	PreLaunchEnabled               bool     `json:"pre_launch_enabled,omitempty"`
	PlaneName                      string   `json:"plane_name,omitempty"`
	PrimaryLocale                  string   `json:"primary_locale,omitempty"`
	Province                       string   `json:"province,omitempty"`
	ProvinceCode                   string   `json:"province_code,omitempty"`
	RequiresExtraPaymentsAgreement bool     `json:"requires_extra_payments_agreement,omitempty"`
	SetupRequired                  bool     `json:"setup_required,omitempty"`
	ShopOwner                      string   `json:"shop_owner,omitempty"`
	Source                         string   `json:"source,omitempty"`
	TaxesIncluded                  bool     `json:"taxes_included,omitempty"`
	TaxShipping                    bool     `json:"tax_shipping,omitempty"`
	Timezone                       string   `json:"timezone,omitempty"`
	UpdatedAt                      string   `json:"updated_at,omitempty"`
	WeightUnit                     string   `json:"weight_unit,omitempty"`
	Zip                            string   `json:"zip,omitempty"`
}

type ShopWrapper struct {
	Shop Shop `json:"shop,omitempty"`
}

func (client *ShopifyApiImpl) GetShopDetails(details ShopifyRequestDetails) (result Shop, err error) {
	requestUrl := "https://" + details.ShopName + "/admin/shop.json"

	log.Printf("Requesting the shop details for shop %s using URL %s\n", details.ShopName, requestUrl)

	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return
	}

	req.Header.Add("X-Shopify-Access-Token", details.AccessToken)

	resp, err := client.Http.Do(req)
	if err != nil {
		return
	}

	buf, _ := ioutil.ReadAll(resp.Body)
	log.Println("This is the response for the shopify details: ", string(buf))
	wrapper := ShopWrapper{}
	err = json.Unmarshal(buf, &wrapper)
	if err != nil {
		return
	}

	result = wrapper.Shop

	return
}
