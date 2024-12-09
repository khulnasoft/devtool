// Copyright (c) 2023 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package cmd

import (
	"encoding/json"
	"net/http"
	"runtime"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/khulnasoft/local-app/pkg/config"
	"github.com/khulnasoft/local-app/pkg/constants"
	"github.com/khulnasoft/local-app/pkg/selfupdate"
	"github.com/opencontainers/go-digest"
)

func TestVersionUpdateCmd(t *testing.T) {
	RunCommandTests(t, []CommandTest{
		{
			Name:        "happy path",
			Commandline: []string{"version", "update"},
			PrepServer: func(mux *http.ServeMux) {
				newBinary := []byte("#!/bin/bash\necho hello world")
				mux.HandleFunc(selfupdate.DevtoolCLIBasePath+"/manifest.json", func(w http.ResponseWriter, r *http.Request) {
					mf, err := json.Marshal(selfupdate.Manifest{
						Version: semver.MustParse("v9999.0"),
						Binaries: []selfupdate.Binary{
							{
								Filename: "devtool",
								OS:       runtime.GOOS,
								Arch:     runtime.GOARCH,
								Digest:   digest.FromBytes(newBinary),
							},
						},
					})
					if err != nil {
						t.Fatal(err)
					}
					_, _ = w.Write(mf)
				})
				mux.HandleFunc(selfupdate.DevtoolCLIBasePath+"/devtool", func(w http.ResponseWriter, r *http.Request) {
					_, _ = w.Write(newBinary)
				})
			},
			Config: AddActiveTestContext(&config.Config{}),
		},
		{
			Name:        "no update needed",
			Commandline: []string{"version", "update"},
			PrepServer: func(mux *http.ServeMux) {
				mux.HandleFunc(selfupdate.DevtoolCLIBasePath+"/manifest.json", func(w http.ResponseWriter, r *http.Request) {
					mf, err := json.Marshal(selfupdate.Manifest{
						Version: constants.Version,
					})
					if err != nil {
						t.Fatal(err)
					}
					_, _ = w.Write(mf)
				})
			},
			Config: AddActiveTestContext(&config.Config{}),
		},
	})
}
