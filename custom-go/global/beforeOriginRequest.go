package global

import (
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
)

func BeforeOriginRequest(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientRequest, error) {
	return body.Request, nil
}
