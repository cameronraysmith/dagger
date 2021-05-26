package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.dagger.io/dagger/cmd/dagger/cmd/common"
	"go.dagger.io/dagger/cmd/dagger/logger"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new empty environment",
	Args:  cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		// Fix Viper bug for duplicate flags:
		// https://github.com/spf13/viper/issues/233
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			panic(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		lg := logger.New()
		ctx := lg.WithContext(cmd.Context())

		workspace := common.CurrentWorkspace(ctx)

		if viper.GetString("environment") != "" {
			lg.
				Fatal().
				Msg("cannot use option -e,--environment for this command")
		}
		name := args[0]
		if _, err := workspace.Create(ctx, name); err != nil {
			lg.Fatal().Err(err).Msg("failed to create environment")
		}
	},
}

func init() {
	if err := viper.BindPFlags(newCmd.Flags()); err != nil {
		panic(err)
	}
}
