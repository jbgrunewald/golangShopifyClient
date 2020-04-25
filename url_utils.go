package shopify

import "strings"

/*
A helper for building the base url for a shopify request
*/
func BuildBaseUrl(request Request) (url string) {
	pathBuilder := strings.Builder{}
	pathBuilder.WriteString("https://")

	//support authentication with api key and password
	if request.Context.ApiKey != "" && request.Context.Password != "" {
		pathBuilder.WriteString(request.Context.ApiKey)
		pathBuilder.WriteString(":")
		pathBuilder.WriteString(request.Context.Password)
		pathBuilder.WriteString("@")
	}

	pathBuilder.WriteString(request.Context.ShopName)
	pathBuilder.WriteString("/admin/")
	pathBuilder.WriteString(request.Version.String())
	pathBuilder.WriteString("/")
	url = pathBuilder.String()
	return
}

/*
A Helper function to build a request url with only a resource component.
*/
func BuildSimpleUrl(request Request, resource string) (url string) {
	pathBuilder := strings.Builder{}
	pathBuilder.WriteString(BuildBaseUrl(request))
	pathBuilder.WriteString("/")
	pathBuilder.WriteString(resource)
	pathBuilder.WriteString(".json")
	url = pathBuilder.String()
	return
}

/*
A helper for building a url that has an id
*/
func BuildIdUrl(request Request, resource string, id int) (url string) {
	pathBuilder := strings.Builder{}
	pathBuilder.WriteString(BuildBaseUrl(request))
	pathBuilder.WriteString("/")
	pathBuilder.WriteString(resource)
	pathBuilder.WriteString(string(id))
	pathBuilder.WriteString(".json")
	url = pathBuilder.String()
	return
}
