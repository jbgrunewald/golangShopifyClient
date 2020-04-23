package shopify

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
)

type Product struct {
	BodyHtml                       string           `json:"body_html,omitempty"`
	CreatedAt                      string           `json:"created_at,omitempty"`
	Handle                         string           `json:"handle,omitempty"`
	Id                             int              `json:"id,omitempty"`
	Images                         []ProductImage   `json:"images,omitempty"`
	Title                          string           `json:"title,omitempty"`
	Variants                       []ProductVariant `json:"variants,omitempty"`
	Options                        []ProductOption  `json:"options,omitempty"`
	ProductType                    string           `json:"product_type,omitempty"`
	PublishedAt                    string           `json:"published_at,omitempty"`
	PublishedScope                 string           `json:"published_scope,omitempty"`
	Tags                           string           `json:"tags,omitempty"`
	TemplateSuffix                 string           `json:"template_suffix,omitempty"`
	MetafieldsGlobalTitleTag       string           `json:"metafields_global_title_tag,omitempty"`
	MetafieldsGlobalDescriptionTag string           `json:"metafields_global_description_tag,omitempty"`
	UpdateAt                       string           `json:"update_at,omitempty"`
	Vendor                         string           `json:"vendor,omitempty"`
	Image                          ProductImage     `json:"image,omitempty"`
}

type ProductOption struct {
	Id        int      `json:"id"`
	ProductId int      `json:"product_id"`
	Name      string   `json:"name"`
	Position  int      `json:"position"`
	Values    []string `json:"values"`
}

type ProductImage struct {
	Id         int    `json:"id"`
	ProductId  int    `json:"product_id"`
	Position   int    `json:"position"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	Src        string `json:"src"`
	VariantIds []int  `json:"variant_ids"`
}

type ProductWrapper struct {
	Products []Product `json:"products"`
}

type ProductRequestOptions struct {
	Ids                   []int    `url:"ids,omitempty"`
	Limit                 int      `url:"limit,omitempty"`
	SinceId               int      `url:"since_id"`
	Title                 string   `url:"title,omitempty"`
	Vendor                string   `url:"vendor,omitempty"`
	Handle                string   `url:"handle,omitempty"`
	ProductType           string   `url:"product_type,omitempty"`
	CollectionId          int      `url:"collection_id,omitempty"`
	CreatedAtMin          string   `url:"created_at_min,omitempty"`
	CreatedAtMax          string   `url:"created_at_max,omitempty"`
	UpdatedAtMin          string   `url:"updated_at_min,omitempty"`
	UpdatedAtMax          string   `url:"updated_at_max,omitempty"`
	PublishedAtMin        string   `url:"published_at_min,omitempty"`
	PublishedAtMax        string   `url:"published_at_max,omitempty"`
	PublishedStatus       string   `url:"published_status,omitempty"`
	Fields                []string `url:"fields,omitempty"`
	PresentmentCurrencies string   `url:"presentment_currencies,omitempty"`
	All                   bool
}

func (c *RestAdminClient) ProductList(details Request, options ProductRequestOptions) (products []Product, err error) {
	v, err := query.Values(options)
	if err != nil {
		c.Logger.Println("there's an issue setting up the query params")
		return
	}
	requestUrl := "https://" + details.ShopName + "/admin/products.json?" + v.Encode()
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
	c.Logger.Println("This is the response for the products: ", string(buf))
	wrapper := ProductWrapper{}
	err = json.Unmarshal(buf, &wrapper)
	if err != nil {
		return
	}

	products = wrapper.Products

	if len(products) < options.Limit || !options.All || (options.Limit == 0 && len(products) < 50) {
		return
	}

	options.SinceId = products[len(products)-1].Id
	nextResult, err := c.ProductList(details, options)

	products = append(products, nextResult...)

	return
}
