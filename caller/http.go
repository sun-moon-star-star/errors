package caller

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Http struct {
	NoUseSSL    bool
	Method      uint8 // 0 - default post
	ContentType uint8 // 0 - default application/json
}

var HTTPUnknownMethodError = fmt.Errorf("%s", "Unknown Method")
var HTTPUnknownContentTypeError = fmt.Errorf("%s", "Unknown ContentType")

func (h *Http) url(c *Caller) string {
	var url string
	if c.RequestConfig.IP != "" && c.RequestConfig.Port > 0 {
		if h.NoUseSSL {
			url = fmt.Sprintf("http://%s:%d/%s", c.RequestConfig.IP,
				c.RequestConfig.Port, c.RequestConfig.ServiceName)
		} else {
			url = fmt.Sprintf("https://%s:%d/%s", c.RequestConfig.IP,
				c.RequestConfig.Port, c.RequestConfig.ServiceName)
		}
	} else if c.RequestConfig.ServiceName != "" {
		url = c.RequestConfig.ServiceName
	}
	return url
}

func (h *Http) request(url string, c *Caller) error {
	var contentType string

	if h.ContentType == 0 {
		contentType = "application/json"
	} else {
		return HTTPUnknownContentTypeError
	}

	var err error

	switch h.Method {
	case 0:
		resp, err := http.Post(url, contentType, bytes.NewReader(c.Req))

		if err != nil {
			break
		}

		c.Resp, err = ioutil.ReadAll(resp.Body)
	default:
		return HTTPUnknownMethodError
	}

	return err
}

func (h *Http) Call(c *Caller) error {
	return h.request(h.url(c), c)
}
