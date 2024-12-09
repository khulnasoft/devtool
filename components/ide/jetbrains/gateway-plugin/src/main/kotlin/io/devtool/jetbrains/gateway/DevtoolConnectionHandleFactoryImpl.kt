// Copyright (c) 2024 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.devtool.jetbrains.gateway

import com.intellij.openapi.components.Service
import com.jetbrains.gateway.api.GatewayConnectionHandle
import com.jetbrains.gateway.ssh.ClientOverSshTunnelConnector
import com.jetbrains.gateway.ssh.HostTunnelConnector
import com.jetbrains.gateway.thinClientLink.ThinClientHandle
import com.jetbrains.rd.util.lifetime.Lifetime
import io.devtool.jetbrains.gateway.DevtoolConnectionProvider.ConnectParams
import io.devtool.jetbrains.gateway.common.DevtoolConnectionHandle
import io.devtool.jetbrains.gateway.common.DevtoolConnectionHandleFactory
import java.net.URI
import javax.swing.JComponent

@Suppress("UnstableApiUsage")
class DevtoolConnectionHandleFactoryImpl: DevtoolConnectionHandleFactory {
    override fun createDevtoolConnectionHandle(
        lifetime: Lifetime,
        component: JComponent,
        params: ConnectParams
    ): GatewayConnectionHandle {
        return DevtoolConnectionHandle(lifetime, component, params)
    }

    override suspend fun connect(lifetime: Lifetime, connector: HostTunnelConnector, tcpJoinLink: URI): ThinClientHandle {
        return ClientOverSshTunnelConnector(
            lifetime,
            connector
        ).connect(tcpJoinLink, null)
    }
}
