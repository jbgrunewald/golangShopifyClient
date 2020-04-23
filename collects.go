package shopify

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
)

type Collect struct {
	CollectionId int    `json:"collection_id"`
	CreateAt     string `json:"create_at"`
	Featured     bool   `json:"featured"`
	Id           int    `json:"id"`
	Position     int    `json:"position"`
	ProdutId     int    `json:"produt_id"`
	SortValue    string `json:"sort_value"`
	UpdatedAt    string `json:"updated_at"`
}

type CollectWrapper struct {
	Collects []Collect `json:"collect"`
}

type CollectRequestOptions struct {
	Limit        int      `url:"limit,omitempty"`
	SinceId      int      `url:"since_id,omitempty"`
	Fields       []string `url:"fields,omitempty"`
	CollectionId int      `url:"collection_id,omitempty"`
	ProductId    int      `url:"product_id,omitempty"`
	All          bool
}

func (c *RestAdminClient) CollectList(details RequestDetails, options CollectRequestOptions) (result []Collect, err error) {
	v, err := query.Values(options)
	requestUrl := "https://" + details.ShopName + "/admin/api/2019-04/collects.json?" + v.Encode()
	c.Logger.Println("This is the request url for the collects", requestUrl)

	c.Logger.Printf("Requesting collects for shop %s using options %v\n", details.ShopName, options)

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
	c.Logger.Println("This is the response from the request for the collects: ", string(buf))
	wrapper := CollectWrapper{}
	err = json.Unmarshal(buf, &wrapper)
	if err != nil {
		return
	}

	result = wrapper.Collects

	if len(result) < options.Limit || !options.All {
		return
	}

	options.SinceId = result[len(result)-1].Id
	nextResult, err := c.CollectList(details, options)

	result = append(result, nextResult...)

	return
}
