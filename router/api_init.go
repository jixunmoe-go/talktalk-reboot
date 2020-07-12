package router

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/http/cookiejar"
)

type csrfParams struct {
	Param string `json:"csrf_param"`
	Token string `json:"csrf_token"`
}

type Session struct {
	csrfParams
	UnknownN string
}

type Client struct {
	BaseURL string
	client  http.Client
	session Session
}

func (r *Client) Init(baseURL string) {
	if baseURL[len(baseURL)-1] == '/' {
		r.BaseURL = baseURL[:len(baseURL)-1]
	} else {
		r.BaseURL = baseURL
	}

	cookieJar, err := cookiejar.New(nil)
	check(err)

	r.session = Session{}
	r.client = http.Client{
		Jar: cookieJar,
	}
	resp, err := r.client.Get(r.url("/"))
	check(err)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	check(err)

	mapping := map[string]*string{
		"csrf_param": &r.session.Param,
		"csrf_token": &r.session.Token,
		"n":          &r.session.UnknownN,
	}

	doc.Find("meta[content]").Each(func(i int, el *goquery.Selection) {
		name, exists := el.Attr("name")
		s, keyExists := mapping[name]
		if exists && keyExists {
			*s = el.AttrOr("content", "")
		}
	})

	_ = resp.Body.Close()
}
