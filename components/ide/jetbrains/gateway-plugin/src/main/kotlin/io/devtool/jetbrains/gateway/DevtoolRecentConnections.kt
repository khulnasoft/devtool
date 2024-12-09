// Copyright (c) 2022 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.devtool.jetbrains.gateway

import com.intellij.ui.dsl.builder.AlignX
import com.intellij.ui.dsl.builder.AlignY
import com.intellij.ui.dsl.builder.panel
import com.jetbrains.gateway.api.GatewayRecentConnections
import com.jetbrains.rd.util.lifetime.Lifetime
import io.devtool.jetbrains.icons.DevtoolIcons
import javax.swing.JComponent

class DevtoolRecentConnections : GatewayRecentConnections {

    override val recentsIcon = DevtoolIcons.Logo

    private lateinit var view: DevtoolWorkspacesView
    override fun createRecentsView(lifetime: Lifetime): JComponent {
        this.view = DevtoolWorkspacesView(lifetime)
        return panel {
            row {
                resizableRow()
                cell(view.component)
                    .resizableColumn()
                    .align(AlignX.FILL)
                    .align(AlignY.FILL)
                cell()
            }
        }
    }

    override fun getRecentsTitle(): String {
        return "Devtool"
    }

    override fun updateRecentView() {
        if (this::view.isInitialized) {
            this.view.refresh()
        }
    }
}
