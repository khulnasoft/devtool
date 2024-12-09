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
	"time"
)

func ExampleListTeams() {
	token := "devtool_pat_example.personal-access-token"

	devtool, err := client.New(client.WithCredentials(token))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to construct devtool client %v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	response, err := devtool.Teams.ListTeams(ctx, connect.NewRequest(&v1.ListTeamsRequest{}))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to list teams %v", err)
		return
	}

	fmt.Fprintf(os.Stdout, "Retrieved teams %v", response.Msg.GetTeams())
}

func ExampleGetTeam() {
	token := "devtool_pat_example.personal-access-token"

	devtool, err := client.New(client.WithCredentials(token))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to construct devtool client %v", err)
		return
	}

	response, err := devtool.Teams.GetTeam(context.Background(), connect.NewRequest(&v1.GetTeamRequest{
		TeamId: "<TEAM_ID>",
	}))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get team %v", err)
		return
	}

	fmt.Fprintf(os.Stdout, "Retrieved team %v", response.Msg.GetTeam())
}
