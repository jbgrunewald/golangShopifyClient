package shopify

import "strings"

/*
Used to extract the cursor based url from the response header.
*/
func ExtractNextCursorUrl(header string) string {
	headerSplit := strings.Split(header, " ")
	next := headerSplit[0]
	if next != "" {
		return next[1 : len(headerSplit[0])-2]
	}
	return next
}

/*
A helper for building the base url for a shopify request
*/
func BuildBaseUrl(request Request) (url string) {
	pathBuilder := strings.Builder{}
	pathBuilder.WriteString("https://")
	pathBuilder.WriteString(request.Context.ShopName)
	pathBuilder.WriteString("/admin/")
	pathBuilder.WriteString(request.Version.String())
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
