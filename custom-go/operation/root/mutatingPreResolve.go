package root
import (
	"custom-go/generated"
	"custom-go/pkg/base"
)

func MutatingPreResolve(hook *base.HookRequest, body generated.operation__rootBody) (res generated.operation__rootBody, err error) {
	hook.Logger().Info("MutatingPreResolve")
	return body, nil
}
