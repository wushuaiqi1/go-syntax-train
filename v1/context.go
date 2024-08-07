package v1

import (
	"encoding/json"
	"io"
	"net/http"
)

type Context struct {
	Response http.ResponseWriter `json:"response"`
	Request  *http.Request       `json:"request"`
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Response: w,
		Request:  r,
	}
}

// ReadJson 序列化
func (c *Context) ReadJson() (data any, r error) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return
	}
	return data, nil
}

func (c *Context) WriteJson(code int, resp interface{}) error {
	c.Response.WriteHeader(code)
	json, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = c.Response.Write(json)
	if err != nil {
		return err
	}
	return nil
}

func (c *Context) OfSuccess(resp interface{}) {
	err := c.WriteJson(http.StatusOK, resp)
	if err != nil {
		return
	}
}

func (c *Context) OfFail(resp interface{}) {
	err := c.WriteJson(http.StatusBadRequest, resp)
	if err != nil {
		return
	}
}
