package shopify

type ProductVariant struct {
	Barcode             string             `json:"barcode,omitempty"`
	CompareAtPrice      string             `json:"compare_at_price,omitempty"`
	CreateAt            string             `json:"create_at,omitempty"`
	FulfillmentService  string             `json:"fulfillment_service,omitempty"`
	Grams               int                `json:"grams,omitempty"`
	Id                  int                `json:"id,omitempty"`
	ImageId             int                `json:"image_id,omitempty"`
	InventoryItemId     int                `json:"inventory_item_id,omitempty"`
	InventoryManagement string             `json:"inventory_management,omitempty"`
	InventoryPolicy     string             `json:"inventory_policy,omitempty"`
	InventoryQuantity   int                `json:"inventory_quantity,omitempty"`
	Metafields          []MetaFields       `json:"metafields,omitempty"`
	OptionOne           string             `json:"option1,omitempty"`
	OptionTwo           string             `json:"option2,omitempty"`
	OptionThree         string             `json:"option3,omitempty"`
	PresentmentPrices   []PresentmentPrice `json:"presentment_prices,omitempty"`
	Position            int                `json:"position,omitempty"`
	Price               string             `json:"price,omitempty"`
	ProductId           int                `json:"product_id,omitempty"`
	Sku                 string             `json:"sku,omitempty"`
	Taxable             bool               `json:"taxable,omitempty"`
	TaxCode             string             `json:"tax_code,omitempty"`
	Title               string             `json:"title,omitempty"`
	UpdatedAt           string             `json:"updated_at,omitempty"`
	Weight              float32            `json:"weight,omitempty"`
	WeightUnit          string             `json:"weight_unit,omitempty"`
}

type PresentmentPrice struct {
	Price          Price `json:"price"`
	CompareAtPrice Price `json:"compare_at_price"`
}

type Price struct {
	CurrencyCode string `json:"currency_code"`
	Amount       string `json:"amount"`
}

type ProductMetafield struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	ValueType   string `json:"value_type"`
	Namespace   string `json:"namespace"`
	Description string `json:"description,omitempty"`
}
