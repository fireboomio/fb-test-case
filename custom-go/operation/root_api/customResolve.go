package root_api
import (
	"custom-go/generated"
	"custom-go/pkg/base"
)

func CustomResolve(hook *base.HookRequest, body generated.operation__root_apiBody) (res generated.operation__root_apiBody, err error) {
	hook.Logger().Info("CustomResolve")
	return body, nil
}
