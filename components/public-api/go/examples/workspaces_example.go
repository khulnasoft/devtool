// Copyright (c) 2023 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package examples

import (
	"context"
	"fmt"
	"os"

	"github.com/bufbuild/connect-go"
	"github.com/khulnasoft/devtool/components/public-api/go/client"
	v1 "github.com/khulnasoft/devtool/components/public-api/go/experimental/v1"
)

func ExampleListAllWorkspaces() {
	token := "devtool_pat_example.personal-access-token"

	devtool, err := client.New(client.WithCredentials(token))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to construct devtool client %v", err)
		return
	}

	response, err := devtool.Workspaces.ListWorkspaces(context.Background(), connect.NewRequest(&v1.ListWorkspacesRequest{}))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to list workspaces %v", err)
		return
	}

	fmt.Fprintf(os.Stdout, "Retrieved workspaces %v", response.Msg.GetResult())
}

func ExampleGetWorkspace() {
	token := "devtool_pat_example.personal-access-token"

	devtool, err := client.New(client.WithCredentials(token))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to construct devtool client %v", err)
		return
	}

	response, err := devtool.Workspaces.GetWorkspace(context.Background(), connect.NewRequest(&v1.GetWorkspaceRequest{
		WorkspaceId: "<WORKSPACE_ID>",
	}))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get workspace %v", err)
		return
	}

	fmt.Fprintf(os.Stdout, "Retrieved workspace %v", response.Msg.GetResult())
}
