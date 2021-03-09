package caller

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Http struct {
	NoUseSSL    bool
	Method      string // default post
	ContentType string // default application/json
}

func (h *Http) Call(c *Caller) error {
	var url string
	if c.RequestConfig.ServiceName != "" {
		url = c.RequestConfig.ServiceName
	} else if c.RequestConfig.IP != "" && c.RequestConfig.Port > 0 {
		if h.NoUseSSL {
			url = fmt.Sprintf("http://%s:%d", c.RequestConfig.IP, c.RequestConfig.Port)
		} else {
			url = fmt.Sprintf("https://%s:%d", c.RequestConfig.IP, c.RequestConfig.Port)
		}
	}

	resp, err := http.Post(url, h.ContentType, bytes.NewReader(c.Req))

	if err != nil {
		return err
	}

	c.Resp, err = ioutil.ReadAll(resp.Body)
	return err
}
