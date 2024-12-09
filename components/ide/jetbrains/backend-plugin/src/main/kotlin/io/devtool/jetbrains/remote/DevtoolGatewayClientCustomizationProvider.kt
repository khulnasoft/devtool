// Copyright (c) 2022 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.devtool.jetbrains.remote

import com.jetbrains.rdserver.unattendedHost.customization.GatewayClientCustomizationProvider
import com.jetbrains.rdserver.unattendedHost.customization.controlCenter.DefaultGatewayControlCenterProvider
import com.jetbrains.rdserver.unattendedHost.customization.controlCenter.GatewayControlCenterProvider
import com.jetbrains.rdserver.unattendedHost.customization.controlCenter.GatewayHostnameDisplayKind
import io.devtool.jetbrains.remote.icons.DevtoolIcons
import javax.swing.Icon

class DevtoolGatewayClientCustomizationProvider : GatewayClientCustomizationProvider {
    override val icon: Icon = DevtoolIcons.Logo
    override val title: String = System.getenv("JETBRAINS_DEVTOOL_WORKSPACE_HOST") ?: DefaultGatewayControlCenterProvider().getHostnameShort()

    override val controlCenter: GatewayControlCenterProvider = object : GatewayControlCenterProvider {
        override fun getHostnameDisplayKind() = GatewayHostnameDisplayKind.ShowHostnameOnNavbar
        override fun getHostnameShort() = title
        override fun getHostnameLong() = title
    }
}
