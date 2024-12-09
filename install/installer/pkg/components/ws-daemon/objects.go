// Copyright (c) 2021 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package wsdaemon

import (
	"github.com/khulnasoft/devtool/installer/pkg/common"
)

var Objects = common.CompositeRenderFunc(
	role,
	clusterrole,
	configmap,
	common.DefaultServiceAccount(Component),
	daemonset,
	rolebinding,
	common.GenerateService(Component, []common.ServicePort{
		{
			Name:          "rpc",
			ContainerPort: ServicePort,
			ServicePort:   ServicePort,
		},
	}),
	tlssecret,
)
