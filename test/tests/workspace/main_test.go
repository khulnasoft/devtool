// Copyright (c) 2021 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package workspace

import (
	"context"
	"os"
	"testing"

	"github.com/khulnasoft/devtool/test/pkg/integration"
	"sigs.k8s.io/e2e-framework/pkg/env"
)

var (
	testEnv    env.Environment
	username   string
	namespace  string
	kubeconfig string
	gitlab     bool
)

func TestMain(m *testing.M) {
	username, namespace, testEnv, _, kubeconfig, gitlab = integration.Setup(context.Background())
	os.Exit(testEnv.Run(m))
}