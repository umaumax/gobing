package gobing

type ImageResult struct {
	D struct {
		Next    string `json:"__next"`
		Results []struct {
			ContentType string `json:"ContentType"`
			DisplayUrl  string `json:"DisplayUrl"`
			FileSize    string `json:"FileSize"`
			Height      string `json:"Height"`
			ID          string `json:"ID"`
			MediaUrl    string `json:"MediaUrl"`
			SourceUrl   string `json:"SourceUrl"`
			Thumbnail   struct {
				ContentType string `json:"ContentType"`
				FileSize    string `json:"FileSize"`
				Height      string `json:"Height"`
				MediaUrl    string `json:"MediaUrl"`
				Width       string `json:"Width"`
				Metadata    struct {
					Type string `json:"type"`
				} `json:"__metadata"`
			} `json:"Thumbnail"`
			Title    string `json:"Title"`
			Width    string `json:"Width"`
			Metadata struct {
				Type string `json:"type"`
				Uri  string `json:"uri"`
			} `json:"__metadata"`
		} `json:"results"`
	} `json:"d"`
}
