// Copyright (c) 2021 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package external

import (
	"github.com/khulnasoft/devtool/installer/pkg/common"
	dbinit "github.com/khulnasoft/devtool/installer/pkg/components/database/init"
)

var Objects = common.CompositeRenderFunc(
	dbinit.Objects,
)
