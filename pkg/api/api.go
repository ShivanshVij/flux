package api

import (
	"github.com/shivanshvij/flux/pkg/sdcp"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/loopholelabs/logging/types"

	"github.com/shivanshvij/flux/internal/config"
	"github.com/shivanshvij/flux/internal/utils"

	v1 "github.com/shivanshvij/flux/pkg/api/v1"
	v1Docs "github.com/shivanshvij/flux/pkg/api/v1/docs"
)

const (
	V1Path = "/v1"
)

type API struct {
	logger types.Logger
	config *config.Config
	app    *fiber.App

	sdcp *sdcp.SDCP
}

func New(config *config.Config, logger types.Logger) *API {
	return &API{
		logger: logger.SubLogger("api"),
		config: config,
		app:    utils.DefaultFiberApp(1024 * 1024 * 500),
	}
}

func (s *API) Start() error {
	listener, err := net.Listen("tcp", s.config.ListenAddress)
	if err != nil {
		return err
	}

	s.sdcp = sdcp.New(s.logger)
	v1Docs.SwaggerInfoapi.Host = s.config.Endpoint
	v1Docs.SwaggerInfoapi.Schemes = []string{"http"}

	s.app.Use(cors.New())
	s.app.Mount(V1Path, v1.New(s.sdcp, s.logger).App())

	return s.app.Listener(listener)
}

func (s *API) Stop() error {
	s.sdcp.Close()
	return s.app.Shutdown()
}
