package route

import (
	"main/internal/application/question_controller"

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
	questionRoutes question_controller.QuestionRoutes,
) Routes {
	return Routes{
	questionRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
