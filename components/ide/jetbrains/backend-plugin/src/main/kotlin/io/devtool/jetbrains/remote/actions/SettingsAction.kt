// Copyright (c) 2022 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.devtool.jetbrains.remote.actions

import com.intellij.openapi.actionSystem.AnAction
import com.intellij.openapi.actionSystem.AnActionEvent
import com.intellij.openapi.components.service
import io.devtool.jetbrains.remote.DevtoolManager
import org.apache.http.client.utils.URIBuilder

class SettingsAction : AnAction() {
    private val manager = service<DevtoolManager>()

    override fun actionPerformed(event: AnActionEvent) {
        manager.pendingInfo.thenAccept { workspaceInfo ->
            URIBuilder(workspaceInfo.devtoolHost).setPath("settings").build().toString().let { url ->
                manager.openUrlFromAction(url)
            }
        }
    }
}
