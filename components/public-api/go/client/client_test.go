// Copyright (c) 2022 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {

	t.Run("with all options", func(t *testing.T) {
		expectedOptions := &options{
			url:         "https://foo.bar.com",
			client:      &http.Client{},
			credentials: "my_awesome_credentials",
		}
		devtool, err := New(
			WithURL(expectedOptions.url),
			WithCredentials(expectedOptions.credentials),
			WithHTTPClient(expectedOptions.client),
		)
		require.NoError(t, err)
		require.Equal(t, expectedOptions, devtool.cfg)

		require.NotNil(t, devtool.PersonalAccessTokens)
		require.NotNil(t, devtool.Workspaces)
		require.NotNil(t, devtool.Projects)
		require.NotNil(t, devtool.PersonalAccessTokens)
		require.NotNil(t, devtool.User)
	})

	t.Run("fails when no credentials specified", func(t *testing.T) {
		_, err := New()
		require.Error(t, err)
	})

	t.Run("defaults to https://api.devtool.io", func(t *testing.T) {
		devtool, err := New(WithCredentials("foo"))
		require.NoError(t, err)

		require.Equal(t, "https://api.devtool.io", devtool.cfg.url)
	})

}
