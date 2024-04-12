package route

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRoutes),
)

type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// here should return routers
func NewRoutes(
) Routes {
	return Routes{
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
