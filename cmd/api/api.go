package api

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/loopholelabs/cmdutils"
	"github.com/loopholelabs/cmdutils/pkg/command"

	"github.com/shivanshvij/flux/internal/config"
	"github.com/shivanshvij/flux/internal/utils"
	"github.com/shivanshvij/flux/pkg/api"
)

// Cmd encapsulates the commands for starting the API
func Cmd() command.SetupCommand[*config.Config] {
	ListenAddress := config.DefaultListenAddress
	Endpoint := config.DefaultEndpoint

	return func(cmd *cobra.Command, ch *cmdutils.Helper[*config.Config]) {
		apiCmd := &cobra.Command{
			Use:   "api",
			Short: "Start the Flux API",
			PreRunE: func(cmd *cobra.Command, args []string) error {
				err := ch.Config.GlobalRequiredFlags(cmd)
				if err != nil {
					return err
				}

				ch.Config.ListenAddress = ListenAddress
				ch.Config.Endpoint = Endpoint

				return ch.Config.Validate()
			},
			RunE: func(cmd *cobra.Command, args []string) error {
				ch.Printer.Println("starting Flux API listening on ", ch.Config.ListenAddress)
				errCh := make(chan error, 1)

				a := api.New(ch.Config, ch.Logger)
				go func() {
					errCh <- a.Start()
				}()

				err := utils.WaitForSignal(errCh)
				if err != nil {
					return fmt.Errorf("error while starting Flux API: %w", err)
				}

				err = a.Stop()
				if err != nil {
					return fmt.Errorf("failed to stop Flux API: %w", err)
				}
				return nil
			},
		}
		cmd.AddCommand(apiCmd)

		apiCmd.Flags().StringVar(&ListenAddress, "listen-address", config.DefaultListenAddress, "The address to listen on")
		apiCmd.Flags().StringVar(&Endpoint, "endpoint", config.DefaultEndpoint, "The endpoint to listen on")
	}
}
