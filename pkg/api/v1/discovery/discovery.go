package discovery

import (
	"github.com/gofiber/fiber/v2"

	"github.com/loopholelabs/logging/types"

	"github.com/shivanshvij/flux/internal/utils"
	"github.com/shivanshvij/flux/pkg/api/v1/models"
	"github.com/shivanshvij/flux/pkg/sdcp"
)

type Discovery struct {
	logger types.Logger
	app    *fiber.App
}

func New(logger types.Logger) *Discovery {
	i := &Discovery{
		logger: logger.SubLogger("DISCOVERY"),
		app:    utils.DefaultFiberApp(),
	}

	i.init()

	return i
}

func (a *Discovery) init() {
	a.logger.Debug().Msg("initializing")
	a.app.Post("/", a.Discovery)
}

// Discovery godoc
// @Description  Discovers new Printers
// @Tags         discovery
// @Accept       application/json
// @Produce      application/json
// @Success      200  {object} models.DiscoveryResponse
// @Failure      500  {string} string
// @Router       /discovery [post]
func (a *Discovery) Discovery(ctx *fiber.Ctx) error {
	a.logger.Debug().Msgf("received Discovery request from %s", ctx.IP())

	discoveries, err := sdcp.Discover(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	res := &models.DiscoveryResponse{
		Discovered: make([]*models.DiscoveryData, len(discoveries)),
	}
	for i, d := range discoveries {
		res.Discovered[i] = &models.DiscoveryData{
			MachineName:     d.Data.MachineName,
			MachineModel:    d.Data.MachineModel,
			BrandName:       d.Data.BrandName,
			MainboardIP:     d.Data.MainboardIP,
			MainboardID:     d.Data.MainboardID,
			ProtocolVersion: d.Data.ProtocolVersion,
			FirmwareVersion: d.Data.FirmwareVersion,
		}
	}

	return ctx.JSON(res)
}

func (a *Discovery) App() *fiber.App {
	return a.app
}
