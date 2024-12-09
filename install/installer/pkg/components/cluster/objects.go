// Copyright (c) 2021 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// The cluster package is designed for cluster-level objects that are always installed
// These will mostly be security related

package cluster

import "github.com/khulnasoft/devtool/installer/pkg/common"

var Objects = common.CompositeRenderFunc(
	certmanager,
	clusterrole,
	resourcequota,
	rolebinding,
	common.DefaultServiceAccount(NobodyComponent),
)
