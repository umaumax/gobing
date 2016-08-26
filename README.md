# golang Bing Search API library

https://onedrive.live.com/view.aspx?resid=9C9479871FBFA822!109&app=Word&authkey=!ACvyZ_MNtngQyCU

## アカウント キー の取得方法
[PHPからBing Search APIを使って画像を収集する - Qiita](http://qiita.com/sadapon2008/items/38628ae8266495b3d3a3)

```
type Result struct {
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
```

DisplayUrlはhttpの場合省略されうる...?

## FYI
### Bing Search API
[bing search apiの使い方 - Qiita](http://qiita.com/ysekky/items/c418919ca436f3104dbe)
[Bing Search API プログラミング解説](http://so-zou.jp/web-app/tech/web-api/bing/search/)
[Bing Search API をつかってURLを取得する - karie's blog](http://karie042.hatenablog.com/entry/2014/12/21/201921)

### golang lib
[maxhawkins/bing-search: Bing Search API Client in Go](https://github.com/maxhawkins/bing-search)

[raitucarp/bing-search: Bing web search for Go language](https://github.com/raitucarp/bing-search)
