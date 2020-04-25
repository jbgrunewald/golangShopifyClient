package shopify

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

func (s ShopWrapper) GetResourceName() string {
	return "shop"
}

func (s ShopWrapper) GetId() int {
	return s.Shop.Id
}
func (s ShopWrapper) BuildGetUrl(request Request) string {
	return BuildSimpleUrl(request, s.GetResourceName())
}

func (c *RestAdminClient) ShopGet(context ShopifyContext) (result Shop, err error) {
	wrapper := &ShopWrapper{}
	err = c.Get(context, wrapper)
	result = wrapper.Shop

	return
}
