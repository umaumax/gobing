package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/umaumax/gobing"
)

var (
	apiType string
)

func init() {
	//	NOTE for debug
	log.SetFlags(log.Llongfile)

	flag.StringVar(&apiType, "api", gobing.API_TYPE_WEB, fmt.Sprintf("%s | %s", gobing.API_TYPE_WEB, gobing.API_TYPE_IMAGE))
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Printf("Usage:%s keyword", os.Args[0])
		return
	}
	keyword := flag.Arg(0)

	c := gobing.NewClient()
	c.Values.Offset = 50
	c.Values.Keyword = keyword
	//	{
	//		b, err := c.Get(c.Values)
	//		fmt.Println(string(b), err)
	//	}

	switch apiType {
	case gobing.API_TYPE_WEB:
		ret, err := c.WebSearch(nil)
		if err != nil {
			log.Fatalln("get bing search api err:", err)
		}
		for i, v := range ret.D.Results {
			fmt.Println(i, v.Title, v.DisplayUrl)
		}
	case gobing.API_TYPE_IMAGE:
		ret, err := c.ImageSearch(nil)
		if err != nil {
			log.Fatalln("get bing search api err:", err)
		}
		for i, v := range ret.D.Results {
			fmt.Println(i, v.Title, v.MediaUrl)
		}
	}
}
