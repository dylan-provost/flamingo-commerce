package product

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo-commerce/v3/product/interfaces/controller"
	"flamingo.me/flamingo-commerce/v3/product/interfaces/templatefunctions"
	"flamingo.me/flamingo/v3/framework/config"
	"flamingo.me/flamingo/v3/framework/flamingo"
	"flamingo.me/flamingo/v3/framework/web"
)

// Module registers our profiler
type Module struct{}

// Configure the product URL
func (m *Module) Configure(injector *dingo.Injector) {
	flamingo.BindTemplateFunc(injector, "getProduct", new(templatefunctions.GetProduct))
	flamingo.BindTemplateFunc(injector, "getProductUrl", new(templatefunctions.GetProductURL))
	flamingo.BindTemplateFunc(injector, "findProducts", new(templatefunctions.FindProducts))

	web.BindRoutes(injector, new(routes))
}

// DefaultConfig for this module
func (m *Module) DefaultConfig() config.Map {
	return config.Map{
		"commerce.product.view.template": "product/product",
		"commerce.product.priceIsGross":  true,
		"templating": config.Map{
			"product": config.Map{
				"attributeRenderer": config.Map{},
			},
		},
	}
}

type routes struct {
	controller *controller.View
}

func (r *routes) Inject(controller *controller.View) {
	r.controller = controller
}

func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.HandleGet("product.view", r.controller.Get)
	h, _ := registry.Route("/product/:marketplacecode/:name.html", `product.view(marketplacecode, name, backurl?="")`)
	h.Normalize("name")
	h, _ = registry.Route("/product/:marketplacecode/:variantcode/:name.html", `product.view(marketplacecode, variantcode, name, backurl?="")`)
	h.Normalize("name")
}
