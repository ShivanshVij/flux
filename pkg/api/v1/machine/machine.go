package machine

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shivanshvij/flux/pkg/api/v1/models"

	"github.com/loopholelabs/logging/types"

	"github.com/shivanshvij/flux/internal/utils"
	"github.com/shivanshvij/flux/pkg/sdcp"
)

type Machine struct {
	logger types.Logger
	app    *fiber.App

	sdcp *sdcp.SDCP
}

func New(sdcp *sdcp.SDCP, logger types.Logger) *Machine {
	i := &Machine{
		logger: logger.SubLogger("machine"),
		app:    utils.DefaultFiberApp(),
		sdcp:   sdcp,
	}

	i.init()

	return i
}

func (a *Machine) init() {
	a.logger.Debug().Msg("initializing")
	a.app.Post("/register", a.Register)
	//a.app.Post("/unregister", a.Unregister)
	a.app.Get("/status/:id", a.Status)
}

// Register godoc
// @Description  Registers a new machine
// @Tags         machine
// @Accept       application/json
// @Produce      application/json
// @Param        request  body models.MachineRegisterRequest true  "Machine Register Request"
// @Success      200  {string} string
// @Failure      400  {string} string
// @Failure      404  {string} string
// @Failure      500  {string} string
// @Router       /machine/register [post]
func (a *Machine) Register(ctx *fiber.Ctx) error {
	a.logger.Debug().Msgf("received Register request from %s", ctx.IP())

	body := new(models.MachineRegisterRequest)
	err := ctx.BodyParser(body)
	if err != nil {
		a.logger.Error().Err(err).Msg("failed to parse body")
		return fiber.NewError(fiber.StatusBadRequest, "failed to parse body")
	}

	err = a.sdcp.Register(body.MachineID, body.MachineIP)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusOK).SendString("OK")
}

// Status godoc
// @Description  Retrieves the status of a machine
// @Tags         machine
// @Accept       application/json
// @Produce      application/json
// @Param        id path string true "id"
// @Success      200  {string} string
// @Failure      400  {string} string
// @Failure      404  {string} string
// @Failure      500  {string} string
// @Router       /machine/status/{id} [get]
func (a *Machine) Status(ctx *fiber.Ctx) error {
	a.logger.Debug().Msgf("received Status request from %s", ctx.IP())

	id := ctx.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	m, ok := a.sdcp.GetMachine(id)
	if !ok {
		return fiber.NewError(fiber.StatusNotFound, "machine not found")
	}

	err := m.StatusRefresh()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusOK).SendString("OK")
}

func (a *Machine) App() *fiber.App {
	return a.app
}
