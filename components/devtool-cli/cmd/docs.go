// Copyright (c) 2023 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package cmd

import (
	"github.com/spf13/cobra"
)

// URL of the Devtool documentation
const DocsUrl = "https://www.devtool.io/docs/introduction"

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Open Devtool Documentation in default browser",
	RunE: func(cmd *cobra.Command, args []string) error {
		return openPreview("GP_EXTERNAL_BROWSER", DocsUrl)
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)
}
