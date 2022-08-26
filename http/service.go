package http

import (
	"github.com/ability-sh/abi-lib/http"
	"github.com/ability-sh/abi-micro/micro"
)

var httpClient = http.NewClient()

type httpService struct {
	config interface{}
	name   string
}

func newHTTPService(name string, config interface{}) HTTPService {
	return &httpService{name: name, config: config}
}

/**
* 服务名称
**/
func (s *httpService) Name() string {
	return s.name
}

/**
* 服务配置
**/
func (s *httpService) Config() interface{} {
	return s.config
}

/**
* 初始化服务
**/
func (s *httpService) OnInit(ctx micro.Context) error {
	return nil
}

/**
* 校验服务是否可用
**/
func (s *httpService) OnValid(ctx micro.Context) error {
	return nil
}

func (s *httpService) Request(ctx micro.Context, method string) http.HTTPRequest {
	return http.NewHTTPRequest(method).SetHeaders(map[string]string{"Trace": ctx.Trace()})
}

func (s *httpService) Recycle() {

}
