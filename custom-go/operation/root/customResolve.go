package root
import (
	"custom-go/generated"
	"custom-go/pkg/base"
)

func CustomResolve(hook *base.HookRequest, body generated.operation__rootBody) (res generated.operation__rootBody, err error) {
	hook.Logger().Info("CustomResolve")
	return body, nil
}
