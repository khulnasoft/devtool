// Copyright (c) 2023 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package cmd

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/bufbuild/connect-go"
	v1 "github.com/khulnasoft/devtool/components/public-api/go/experimental/v1"
	devtool_experimental_v1connect "github.com/khulnasoft/devtool/components/public-api/go/experimental/v1/v1connect"
	"github.com/khulnasoft/local-app/pkg/config"
)

func TestLoginCmd(t *testing.T) {
	RunCommandTests(t, []CommandTest{
		{
			Name:        "test unauthenticated",
			Commandline: []string{"login", "--token", "foo"},
			Config: &config.Config{
				ActiveContext: "test",
			},
			PrepServer: func(mux *http.ServeMux) {
				mux.Handle(devtool_experimental_v1connect.NewTeamsServiceHandler(&testLoginCmdSrv{
					Err: connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("cannot establish caller identity")),
				}))
			},
			Expectation: CommandTestExpectation{
				Error:          "unauthenticated",
				HasResolutions: true,
			},
		},
	})
}

type testLoginCmdSrv struct {
	Err error
	devtool_experimental_v1connect.UnimplementedTeamsServiceHandler
}

func (srv testLoginCmdSrv) ListTeams(context.Context, *connect.Request[v1.ListTeamsRequest]) (*connect.Response[v1.ListTeamsResponse], error) {
	return nil, srv.Err
}
