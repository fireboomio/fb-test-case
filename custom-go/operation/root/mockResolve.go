package root
import (
	"custom-go/generated"
	"custom-go/pkg/base"
)

func MockResolve(hook *base.HookRequest, body generated.operation__rootBody) (res generated.operation__rootBody, err error) {
	hook.Logger().Info("MockResolve")
	return body, nil
}
