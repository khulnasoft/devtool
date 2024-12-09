// Copyright (c) 2022 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.devtool.jetbrains.gateway

import com.intellij.openapi.components.service
import com.intellij.openapi.options.BoundConfigurable
import com.intellij.openapi.ui.DialogPanel
import com.intellij.openapi.ui.ValidationInfo
import com.intellij.ui.components.JBTextField
import com.intellij.ui.dsl.builder.*
import com.intellij.ui.layout.ValidationInfoBuilder

class DevtoolSettingsConfigurable : BoundConfigurable("Devtool") {

    override fun createPanel(): DialogPanel {
        val state = service<DevtoolSettingsState>()
        return panel {
            row {
                textField()
                    .label("Devtool Host:", LabelPosition.LEFT)
                    .align(Align.FILL)
                    .bindText(state::devtoolHost)
                    .validationOnApply(::validateDevtoolHost)
                    .validationOnInput(::validateDevtoolHost)
            }
            row {
                checkBox("Force SSH over HTTP tunnel")
                    .bindSelected(state::forceHttpTunnel)
                    .comment("Helpful if you are behind a firewall/proxy that blocks SSH or " +
                            "have complicated SSH setup (bastions, proxy jumps, etc.)")
            }
            row {
                checkBox("Persistent connection heartbeats")
                    .bindSelected(state::additionalHeartbeat)
                    .comment("Keep workspaces running as long as the IDE connection remains active")
            }

        }
    }

    private fun validateDevtoolHost(
        builder: ValidationInfoBuilder,
        devtoolHost: JBTextField
    ): ValidationInfo? {
        return builder.run {
            if (devtoolHost.text.isBlank()) {
                return@run error("may not be empty")
            }
            return@run null
        }
    }
}
