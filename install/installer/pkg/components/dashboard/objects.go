// Copyright (c) 2021 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package dashboard

import (
	"github.com/khulnasoft/devtool/installer/pkg/common"
	"k8s.io/apimachinery/pkg/runtime"
)

var Objects = common.CompositeRenderFunc(
	deployment,
	networkpolicy,
	rolebinding,
	pdb,
	func(ctx *common.RenderContext) ([]runtime.Object, error) {
		return Role(ctx)
	},
	common.GenerateService(Component, []common.ServicePort{
		{
			Name:          PortName,
			ContainerPort: ContainerPort,
			ServicePort:   ServicePort,
		},
	}),
	common.DefaultServiceAccount(ComponentServiceAccount),
)
