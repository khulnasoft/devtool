// Copyright (c) 2022 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package ide_metrics

import "github.com/khulnasoft/devtool/installer/pkg/common"

var Objects = common.CompositeRenderFunc(
	configmap,
	deployment,
	rolebinding,
	service,
	networkpolicy,
	common.DefaultServiceAccount(Component),
)
