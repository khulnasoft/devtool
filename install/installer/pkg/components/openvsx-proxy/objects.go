// Copyright (c) 2021 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package openvsx_proxy

import "github.com/khulnasoft/devtool/installer/pkg/common"

// todo(sje): conditionally deploy this component

var Objects = common.CompositeRenderFunc(
	configmap,
	networkpolicy,
	rolebinding,
	statefulset,
	pdb,
	service,
	common.DefaultServiceAccount(Component),
)
