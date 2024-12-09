// Copyright (c) 2022 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.devtool.jetbrains.gateway

import com.jetbrains.gateway.api.GatewayConnector
import com.jetbrains.gateway.api.GatewayConnectorDocumentationPage
import com.jetbrains.rd.util.lifetime.Lifetime
import io.devtool.jetbrains.icons.DevtoolIcons
import java.awt.Component

class DevtoolConnector : GatewayConnector {
    override val icon = DevtoolIcons.Logo

    override fun createView(lifetime: Lifetime) = DevtoolConnectorView(lifetime)

    override fun getActionText() = "Connect to Devtool"

    override fun getDescription() = "Connect to Devtool workspaces"

    override fun getDocumentationAction() = GatewayConnectorDocumentationPage("https://www.devtool.io/docs/ides-and-editors/jetbrains-gateway")

    override fun getConnectorId() = "devtool.connector"

    override fun getRecentConnections(setContentCallback: (Component) -> Unit) = DevtoolRecentConnections()

    override fun getTitle() = "Devtool"

    @Deprecated("Not used", ReplaceWith("null"))
    override fun getTitleAdornment() = null

    override fun initProcedure() {}
}
