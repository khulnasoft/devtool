// Copyright (c) 2022 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.devtool.jetbrains.remote.actions

import com.intellij.openapi.actionSystem.AnAction
import com.intellij.openapi.actionSystem.AnActionEvent
import com.intellij.openapi.components.service
import io.devtool.jetbrains.remote.DevtoolManager

class ReportIssueAction : AnAction() {
    private val manager = service<DevtoolManager>()

    override fun actionPerformed(event: AnActionEvent) {
       manager.openUrlFromAction("https://github.com/khulnasoft/devtool/issues/new/choose")
    }
}
