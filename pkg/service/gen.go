package service

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/hdget/hdproject/g"
	"github.com/hdget/hdsdk/lib/svc"
	"github.com/hdget/hdsdk/utils"
	"github.com/pkg/errors"
	"strings"
)

// Generators 定义所有生成器
var (
	Generators = []svc.Generator{
		NewRouteGenerator(),
	}
)

// route generator
type routeGenerator struct {
	svc.Generator
}

func NewRouteGenerator() svc.Generator {
	return &routeGenerator{
		Generator: svc.NewGenerator(),
	}
}

func (m *routeGenerator) Gen(srcPath string) error {
	// 获取routeItems
	routeItems := make([]any, 0)
	fmt.Printf("Generating routes in: %s...\n", srcPath)
	for moduleName, module := range svc.GetInvocationModules() {
		routes, err := module.GetRoutes(srcPath)
		if err != nil {
			return err
		}

		handlerNames := make([]string, 0)
		for _, r := range routes {
			// add route items here
			//routeItems = append(routeItems, &pb.UpdateRouteRequestItem{
			//	App:           g.APP,
			//	Version:       int32(r.Version),
			//	Namespace:     r.Namespace,
			//	Handler:       r.Handler,
			//	Endpoint:      r.Endpoint,
			//	IsPublic:      r.IsPublic,
			//	Methods:       r.Methods,
			//	CallerId:      r.CallerId,
			//	IsRawResponse: r.IsRawResponse,
			//	Comments:      r.Comments,
			//})
			handlerNames = append(handlerNames, r.Handler)
		}

		fmt.Printf(" - module: %-25s total: %-5d handlers: [%s]\n", moduleName, len(routes), strings.Join(handlerNames, ", "))
	}

	varName := "routes"
	pbPath := fmt.Sprintf("%s/autogen/pb", g.APP)
	return utils.
		NewGoFile("service", map[string]string{pbPath: "pb"}).
		DeclareSliceVar(varName, pbPath, routeItems).
		AddMethod(utils.Reflect().GetStructName(m), utils.Reflect().GetFuncName(m.Get), nil, []string{"any"}, []jen.Code{jen.Return(jen.Id(varName))}).
		Save("pkg/service/autogen_routes.go")
}

func (m *routeGenerator) Register() error {
	generated := m.Get()
	if generated == nil {
		return errors.New("invalid generated stuff")
	}

	//routeItems := generated.([]*pb.UpdateRouteRequestItem)
	//if len(routeItems) == 0 {
	//	return nil
	//}
	//
	//_, err := dapr.Invoke("base", 2, "route", "update", &pb.UpdateRouteRequest{
	//	Items: routeItems,
	//})
	return nil
}
