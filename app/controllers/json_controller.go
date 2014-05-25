package controllers

import (
	"fmt"
	"github.com/revel/revel"
)

const JsonContentType string = "application/json; charset=utf-8"

type JsonController struct{}

type JsonErrorResult struct {
	code int
	err  error
}

type JsonOkResult struct{}

func (c JsonController) RenderJsonError(code int, err error) revel.Result {
	return JsonErrorResult{code: code, err: err}
}

func (r JsonErrorResult) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(r.code, JsonContentType)
	json := fmt.Sprintf("{\"message\": \"%s\"}", r.err)
	resp.Out.Write([]byte(json))
}

func (c JsonController) RenderJsonOk() revel.Result {
	return JsonOkResult{}
}

func (r JsonOkResult) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(200, JsonContentType)
	json := fmt.Sprintf("{\"message\": \"ok\"}")
	resp.Out.Write([]byte(json))
}
