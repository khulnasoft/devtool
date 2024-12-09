// Copyright (c) 2023 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package activity

import (
	"time"

	workspacev1 "github.com/khulnasoft/devtool/ws-manager/api/crd/v1"
)

func Last(ws *workspacev1.Workspace) *time.Time {
	lastActivity := ws.Status.LastActivity
	if lastActivity != nil {
		return &lastActivity.Time
	}

	return nil
}
