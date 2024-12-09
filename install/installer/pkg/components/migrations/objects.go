// Copyright (c) 2021 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package migrations

import (
	"github.com/khulnasoft/devtool/installer/pkg/common"
)

var Objects = common.CompositeRenderFunc(
	job,
	rolebinding,
	common.DefaultServiceAccount(Component),
)
