// Copyright (c) 2021 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package cmd

import (
	"fmt"

	"github.com/khulnasoft/devtool/installer/pkg/config"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:        "init",
	Deprecated: `use "config init" command instead`,
	Short:      "Create a base config file",
	Long: `Create a base config file

This file contains all the credentials to install a Devtool instance and
be saved to a repository.`,
	Example: `  # Save config to config.yaml.
  devtool-installer init > config.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.NewDefaultConfig()
		if err != nil {
			panic(err)
		}
		fc, err := config.Marshal(config.CurrentVersion, cfg)
		if err != nil {
			panic(err)
		}

		fmt.Print(string(fc))
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
