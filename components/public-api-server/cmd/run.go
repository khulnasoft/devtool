// Copyright (c) 2022 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package cmd

import (
	"github.com/khulnasoft/devtool/common-go/log"
	"github.com/khulnasoft/devtool/public-api-server/pkg/server"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCommand)
}

var runCommand = &cobra.Command{
	Use:     "run",
	Short:   "Starts the service",
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := getConfig()

		if err := server.Start(log.Log, Version, cfg); err != nil {
			log.WithError(err).Fatal("cannot start server")
		}
	},
}
