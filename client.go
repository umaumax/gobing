package gobing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var DefaultClient = &defaultClient
var defaultClient = Client{
	HTTP:   http.DefaultClient,
	Values: DefaultValues,
	APIKey: "",
}

func NewClient() *Client {
	v := defaultValues
	c := defaultClient
	c.Values = &v
	return &c
}

var DefaultValues = &defaultValues
var defaultValues = Values{
	Offset:  0,
	Count:   50,
	Keyword: "",
}

func NewValues() *Values {
	v := defaultValues
	return &v
}

type Values struct {
	Offset  int
	Count   int
	Keyword string
}

func (v *Values) Build() string {
	values := url.Values{}
	values.Add("$skip", fmt.Sprint(v.Offset))
	values.Add("$top", fmt.Sprint(v.Count))
	//	values.Add("Query", buildQuery(v.Keyword))
	escaped := strings.Replace(values.Encode(), "%24", "$", -1)
	escaped += "&Query=" + buildQuery(v.Keyword)
	return escaped
}

type Client struct {
	HTTP   *http.Client
	Values *Values
	APIKey string
}

func buildQuery(keyword string) string {
	path := &url.URL{Path: keyword}
	escaped := strings.Replace(path.String(), "%27", "'", -1)
	return fmt.Sprint("%27", escaped, "%27")
}

const (
	API_TYPE_WEB   = "Web"
	API_TYPE_IMAGE = "Image"

//	API_TYPE_NEWS                = "News"
//	API_TYPE_SPELLING_SUGGESTION = "SpellingSuggestion"
//	API_TYPE_RELATED_SEARCH      = "RelatedSearch"
//	API_TYPE_COMPOSITE           = "Composite"
)

func (b *Client) Get(apiType string, values *Values) (body []byte, err error) {
	u := fmt.Sprintf("https://api.datamarket.azure.com/Bing/Search/v1/%s?$format=json&%s", apiType, b.Values.Build())

	req, err := http.NewRequest("POST", u, nil)
	if err != nil {
		return
	}
	req.SetBasicAuth(b.APIKey, b.APIKey)

	resp, err := b.HTTP.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("error %d: %s", resp.StatusCode, body)
		return
	}
	return
}

func ParseWebJson(body []byte) (ret *WebResult, err error) {
	ret = &WebResult{}
	if err = json.Unmarshal(body, &ret); err != nil {
		return
	}
	return
}

func ParseImageJson(body []byte) (ret *ImageResult, err error) {
	ret = &ImageResult{}
	if err = json.Unmarshal(body, &ret); err != nil {
		return
	}
	return
}
