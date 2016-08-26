package gobing

type WebResult struct {
	D struct {
		Next    string `json:"__next"`
		Results []struct {
			Description string `json:"Description"`
			DisplayUrl  string `json:"DisplayUrl"`
			ID          string `json:"ID"`
			Title       string `json:"Title"`
			URL         string `json:"Url"`
			Metadata    struct {
				Type string `json:"type"`
				Uri  string `json:"uri"`
			} `json:"__metadata"`
		} `json:"results"`
	} `json:"d"`
}
