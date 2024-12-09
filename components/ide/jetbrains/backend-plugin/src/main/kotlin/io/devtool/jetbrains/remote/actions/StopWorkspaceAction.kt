// Copyright (c) 2022 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.devtool.jetbrains.remote.actions

import com.intellij.openapi.actionSystem.AnAction
import com.intellij.openapi.actionSystem.AnActionEvent
import com.intellij.openapi.components.service
import io.devtool.jetbrains.remote.DevtoolManager

class StopWorkspaceAction : AnAction() {
    private val manager = service<DevtoolManager>()

    override fun actionPerformed(event: AnActionEvent) {
        manager.pendingInfo.thenAccept { workspaceInfo ->
            manager.trackEvent("jb_execute_command_devtool_workspace", mapOf(
                "action" to "stop"
            ))

            manager.client.server.stopWorkspace(workspaceInfo.workspaceId)
        }
    }
}
