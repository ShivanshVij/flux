package utils

import (
	"encoding/json"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

func DefaultFiberApp(bodyLimit ...int) *fiber.App {
	config := fiber.Config{
		DisableStartupMessage: true,
		ReadTimeout:           time.Second * 10,
		WriteTimeout:          time.Second * 10,
		IdleTimeout:           time.Second * 10,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
	}
	if len(bodyLimit) > 0 {
		config.BodyLimit = bodyLimit[0]
	}
	return fiber.New(config)
}

func WaitForSignal(errChan chan error) error {
	sig := make(chan os.Signal, 2)
	defer close(sig)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sig)
	for {
		select {
		case <-sig:
			return nil
		case err := <-errChan:
			if err == nil {
				continue
			}
			return err
		}
	}
}
