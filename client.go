package yaxml

import(
	"fmt"
	"net/http"
	"net/url"
	"encoding/xml"
	"bytes"
	"strconv"
)

const (
	defaultBaseURL = "https://yandex.ru/search/xml"
	mediaType      = "application/xml"
)

type Client struct {
	baseURL *url.URL

	auth 	auth
	filter 	string
	lr 		int
	l10n 	string
}

type auth struct {
	login 	string
	key 	string
}

func NewClient() *Client {
	c := new(Client)
	c.baseURL, _ 	= url.Parse(defaultBaseURL)
	c.filter 		= "none"
	c.lr 			= 1
	c.l10n 			= "ru"
	return c
}

func (c *Client) SetAuth(login, key string) {
	c.auth.login = login
	c.auth.key = key
	q := c.baseURL.Query()
	q.Add("user", login)
	q.Add("key", key)
	c.baseURL.RawQuery = q.Encode()	
}

func (c *Client) SetFilter(filter string) {
	c.filter = filter
}

func (c *Client) SetLR(lr int) {
	c.lr = lr
}

func (c *Client) SetL10N(l10n string) {
	c.l10n = l10n
}

func (c *Client) GetYandexSearch(req *YandexSearchRequest) (YandexSearch, error) {
	var ys YandexSearch

	b, err := xml.Marshal(req)
	if err != nil {
		return ys, fmt.Errorf("cannot encode Request to XML: %v", err)
	}

	u := c.baseURL
	q := u.Query()
	q.Add("filter", c.filter)
	q.Add("lr", strconv.Itoa(c.lr))
	q.Add("l10n", c.l10n)
	u.RawQuery = q.Encode()

	resp, err := c.request("POST", u.String(), b)
	if err != nil {
		return ys, err
	}

	defer resp.Body.Close()

    if err = xml.NewDecoder(resp.Body).Decode(&ys); err != nil {
    	return ys, fmt.Errorf("cannot decode XML to YandexSearch: %v", err)
    }

    return ys, nil
}

func (c *Client) GetLimits() (*Limits, error) {
	ls := NewLimits()

	u := c.baseURL
	q := u.Query()
	q.Add("action", "limits-info")
	u.RawQuery = q.Encode()

	resp, err := c.request("GET", u.String(), nil)
	if err != nil {
		return ls, err
	}

	defer resp.Body.Close()

    if err = xml.NewDecoder(resp.Body).Decode(&ls); err != nil {
    	return ls, fmt.Errorf("cannot decode XML to XMLYandexSearch: %v", err)
    }

    if err = ls.createIndex(); err != nil {
    	return ls, err
    }

    return ls, nil
}

func (c *Client) request(method, url string, body []byte) (*http.Response, error) {
	var req *http.Request
	var resp *http.Response
	var err error

    if req, err = http.NewRequest(method, url, bytes.NewBuffer(body)); err != nil {
        return resp, fmt.Errorf("cannot create new (http) request: %v", err)
    }

    req.Header.Add("Content-Type", mediaType)

    var client http.Client
    if resp, err = client.Do(req); err != nil {
    	return resp, fmt.Errorf("cannot send request: %v", err)
    }

    return resp, nil
}

