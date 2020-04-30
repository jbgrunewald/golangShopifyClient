package shopify

import "testing"

func TestParseLinkHeader(t *testing.T) {
	header := "<https://test.myshopify.com/admin/api/2019-07/products.json?page_info=hijgklmn&limit=3>; rel=next"
	url := ExtractNextCursorUrl(header)
	if url != "https://test.myshopify.com/admin/api/2019-07/products.json?page_info=hijgklmn&limit=3" {
		t.Errorf("the extracted header value is not what was expected %s", url)
	}
}
