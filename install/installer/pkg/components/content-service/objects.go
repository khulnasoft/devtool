// Copyright (c) 2021 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package content_service

import (
	"github.com/khulnasoft/devtool/common-go/baseserver"
	"github.com/khulnasoft/devtool/installer/pkg/common"
)

var Objects = common.CompositeRenderFunc(
	configmap,
	deployment,
	pdb,
	networkpolicy,
	rolebinding,
	common.GenerateService(Component, []common.ServicePort{
		{
			Name:          RPCServiceName,
			ContainerPort: RPCPort,
			ServicePort:   RPCPort,
		},
		{
			Name:          baseserver.BuiltinMetricsPortName,
			ContainerPort: baseserver.BuiltinMetricsPort,
			ServicePort:   baseserver.BuiltinMetricsPort,
		},
	}),
	common.DefaultServiceAccount(Component),
)
