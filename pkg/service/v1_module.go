package service

import (
	"context"
	"github.com/dapr/go-sdk/service/common"
	"github.com/hdget/hdproject/g"
	"github.com/hdget/hdsdk/lib/svc"
)

type v1_module struct{}

func init() {
	v := &v1_module{}
	err := svc.RegisterAsDaprModule(g.APP, v, map[string]svc.DaprInvocationHandler{
		"example": v.exampleHandler,
	})
	if err != nil {
		panic(err)
	}
}

// exampleHandler 示例handler
// @hd.route
func (*v1_module) exampleHandler(ctx context.Context, event *common.InvocationEvent) (any, error) {
	// write invocation handler logic here
	return nil, nil
}
