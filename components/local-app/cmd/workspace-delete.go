// Copyright (c) 2023 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package cmd

import (
	"context"
	"log/slog"
	"time"

	"github.com/bufbuild/connect-go"
	v1 "github.com/khulnasoft/devtool/components/public-api/go/experimental/v1"
	"github.com/spf13/cobra"
)

// stopWorkspaceCommand stops to a given workspace
var workspaceDeleteCmd = &cobra.Command{
	Use:   "delete <workspace-id>",
	Short: "Deletes a given workspace",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		workspaceID := args[0]

		ctx, cancel := context.WithTimeout(cmd.Context(), 30*time.Second)
		defer cancel()

		devtool, err := getDevtoolClient(ctx)
		if err != nil {
			return err
		}

		slog.Debug("Attempting to delete workspace...")
		_, err = devtool.Workspaces.DeleteWorkspace(ctx, connect.NewRequest(&v1.DeleteWorkspaceRequest{WorkspaceId: workspaceID}))

		return err
	},
}

func init() {
	workspaceCmd.AddCommand(workspaceDeleteCmd)
}
