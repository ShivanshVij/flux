package v1

import (
	"github.com/gofiber/fiber/v2"

	"github.com/loopholelabs/logging/types"

	"github.com/shivanshvij/flux/internal/utils"
	"github.com/shivanshvij/flux/pkg/api/v1/discovery"
	"github.com/shivanshvij/flux/pkg/api/v1/docs"
	"github.com/shivanshvij/flux/pkg/api/v1/models"
)

//go:generate go run -mod=mod github.com/swaggo/swag/cmd/swag@v1.16.3 init -g v1.go -o docs --instanceName api -d ./
type V1 struct {
	logger types.Logger
	app    *fiber.App
}

func New(logger types.Logger) *V1 {
	v := &V1{
		logger: logger.SubLogger("V1"),
		app:    utils.DefaultFiberApp(1024 * 1024 * 500),
	}

	v.init()

	return v
}

// @title Flux API V1
// @version 1.0
// @description API for Flux, V1
// @license.name Apache 2.0
// @license.url https://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @schemes https
// @BasePath /v1
func (v *V1) init() {
	v.logger.Debug().Msg("initializing")
	v.app.Get("/swagger.json", func(ctx *fiber.Ctx) error {
		ctx.Response().Header.SetContentType("application/json")
		return ctx.SendString(docs.SwaggerInfoapi.ReadDoc())
	})

	v.app.Mount("/discovery", discovery.New(v.logger).App())

	v.app.Get("/health", v.Health)
}

// Health godoc
// @Description  Returns the health and status of the various services that make up the API.
// @Tags         health
// @Accept       application/json
// @Produce      application/json
// @Success      200 {object} models.HealthResponse
// @Failure      500 {string} string
// @Router       /health [get]
func (v *V1) Health(ctx *fiber.Ctx) error {
	return ctx.JSON(&models.HealthResponse{})
}

func (v *V1) App() *fiber.App {
	return v.app
}
