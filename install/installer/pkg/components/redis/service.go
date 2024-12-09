// Copyright (c) 2022 Devtool GmbH. All rights reserved.
/// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package redis

import (
	"github.com/khulnasoft/devtool/installer/pkg/common"
	"k8s.io/apimachinery/pkg/runtime"
)

func service(ctx *common.RenderContext) ([]runtime.Object, error) {
	return common.GenerateService(Component, []common.ServicePort{
		{
			Name:          PortName,
			ContainerPort: Port,
			ServicePort:   Port,
		},
	})(ctx)
}
