package gobing

func (b *Client) DefaultWebSearch(keyword string) (ret *WebResult, err error) {
	v := defaultValues
	v.Keyword = keyword
	return b.WebSearch(&v)
}

func (b *Client) WebSearch(v *Values) (ret *WebResult, err error) {
	if v == nil {
		v = b.Values
	}

	var body []byte
	body, err = b.Get(API_TYPE_WEB, v)
	if err != nil {
		return
	}
	return ParseWebJson(body)
}

func (b *Client) DefaultImageSearch(keyword string) (ret *ImageResult, err error) {
	v := defaultValues
	v.Keyword = keyword
	return b.ImageSearch(&v)
}

func (b *Client) ImageSearch(v *Values) (ret *ImageResult, err error) {
	if v == nil {
		v = b.Values
	}

	var body []byte
	body, err = b.Get(API_TYPE_IMAGE, v)
	if err != nil {
		return
	}
	return ParseImageJson(body)
}
