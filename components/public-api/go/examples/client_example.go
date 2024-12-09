// Copyright (c) 2022 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package examples

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	"github.com/khulnasoft/devtool/components/public-api/go/client"
	v1 "github.com/khulnasoft/devtool/components/public-api/go/experimental/v1"
	"os"
)

func ExampleClient() {
	token := "devtool_pat_example.personal-access-token"
	devtool, err := client.New(client.WithCredentials(token))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to construct devtool client %v", err)
		return
	}

	// use the devtool client to access resources
	devtool.Teams.ListTeams(context.Background(), connect.NewRequest(&v1.ListTeamsRequest{}))
}
