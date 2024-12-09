// Copyright (c) 2022 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

//go:build !sentinel
// +build !sentinel

package main

import (
	"github.com/khulnasoft/devtool/agent-smith/cmd"
)

func main() {
	cmd.Execute()
}
