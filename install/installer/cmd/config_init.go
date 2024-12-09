// Copyright (c) 2022 Devtool GmbH. All rights reserved.
/// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package cmd

import (
	"fmt"

	"github.com/khulnasoft/devtool/installer/pkg/config"
	"github.com/spf13/cobra"
)

var configInitOpts struct {
	OverwriteConfig bool
}

// configInitCmd represents the validate command
var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a base config file",
	Long: `Create a base config file

This file contains all the credentials to install a Devtool instance and
be saved to a repository.`,
	Example: `  # Save config to config.yaml.
devtool-installer config init -c ./devtool.config.yaml`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check file isn't present
		exists, err := configFileExists()
		if err != nil {
			return err
		}
		if *exists && !configInitOpts.OverwriteConfig {
			return fmt.Errorf("file %s exists - to overwrite add --overwrite flag", configOpts.ConfigFile)
		}

		cfg, err := config.NewDefaultConfig()
		if err != nil {
			return err
		}

		return saveConfigFile(cfg)
	},
}

func init() {
	configCmd.AddCommand(configInitCmd)

	configInitCmd.Flags().BoolVar(&configInitOpts.OverwriteConfig, "overwrite", false, "overwrite config file if it exists")
}
