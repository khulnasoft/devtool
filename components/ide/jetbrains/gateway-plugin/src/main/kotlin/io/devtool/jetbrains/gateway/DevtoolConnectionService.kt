// Copyright (c) 2022 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.devtool.jetbrains.gateway

import com.intellij.ide.plugins.PluginManagerCore
import com.intellij.openapi.components.Service
import com.intellij.openapi.diagnostic.thisLogger
import com.intellij.openapi.extensions.PluginId
import com.intellij.remoteDev.util.onTerminationOrNow
import com.intellij.util.ExceptionUtil
import com.intellij.util.io.DigestUtil
import com.intellij.util.net.JdkProxyProvider
import com.intellij.util.net.ssl.CertificateManager
import com.jetbrains.rd.util.concurrentMapOf
import com.jetbrains.rd.util.lifetime.Lifetime
import io.devtool.devtoolprotocol.api.DevtoolServerLauncher
import io.devtool.jetbrains.auth.DevtoolAuthService
import kotlinx.coroutines.*
import kotlinx.coroutines.future.await
import org.eclipse.jetty.websocket.api.UpgradeException
import java.net.URL
import java.nio.charset.StandardCharsets

@Service
class DevtoolConnectionService {

    private val clients = concurrentMapOf<String, GatewayDevtoolClient>()

    fun obtainClient(devtoolHost: String): GatewayDevtoolClient {
        return clients.getOrPut(devtoolHost) {
            val lifetime = Lifetime.Eternal.createNested()
            val client = GatewayDevtoolClient(lifetime, devtoolHost)
            val launcher = DevtoolServerLauncher.create(client)
            val job = GlobalScope.launch {
                var accessToken = DevtoolAuthService.getAccessToken(devtoolHost)
                val authorize = suspend {
                    ensureActive()
                    accessToken = DevtoolAuthService.authorize(devtoolHost)
                }
                if (accessToken == null) {
                    authorize()
                }

                val plugin = PluginManagerCore.getPlugin(PluginId.getId("io.devtool.jetbrains.gateway"))!!
                val connect = suspend {
                    ensureActive()

                    val originalClassLoader = Thread.currentThread().contextClassLoader
                    val connection = try {
                        val origin = "https://$devtoolHost/"
                        val proxies = JdkProxyProvider.getInstance().proxySelector.select(URL(origin).toURI())
                        val sslContext = CertificateManager.getInstance().sslContext

                        // see https://intellij-support.jetbrains.com/hc/en-us/community/posts/360003146180/comments/360000376240
                        Thread.currentThread().contextClassLoader = DevtoolConnectionProvider::class.java.classLoader

                        launcher.listen(
                                "wss://$devtoolHost/api/v1",
                                origin,
                                plugin.pluginId.idString,
                                plugin.version,
                                accessToken,
                                proxies,
                                sslContext
                        )
                    } catch (t: Throwable) {
                        val badUpgrade = ExceptionUtil.findCause(t, UpgradeException::class.java)
                        if (badUpgrade?.responseStatusCode == 401 || badUpgrade?.responseStatusCode == 403) {
                            throw InvalidTokenException("failed web socket handshake (${badUpgrade.responseStatusCode})")
                        }
                        throw t
                    } finally {
                        Thread.currentThread().contextClassLoader = originalClassLoader
                    }
                    val tokenHash = DigestUtil.sha256Hex(accessToken!!.toByteArray(StandardCharsets.UTF_8))
                    val tokenScopes = client.server.getDevtoolTokenScopes(tokenHash).await()
                    for (scope in DevtoolAuthService.scopes) {
                        if (!tokenScopes.contains(scope)) {
                            connection.cancel(false)
                            throw InvalidTokenException("$scope scope is not granted")
                        }
                    }
                    connection
                }

                val minReconnectionDelay = 2 * 1000L
                val maxReconnectionDelay = 30 * 1000L
                val reconnectionDelayGrowFactor = 1.5
                var reconnectionDelay = minReconnectionDelay
                while (isActive) {
                    try {
                        val connection = try {
                            connect()
                        } catch (t: Throwable) {
                            val e = ExceptionUtil.findCause(t, InvalidTokenException::class.java) ?: throw t
                            thisLogger().warn("$devtoolHost: invalid token, authorizing again and reconnecting:", e)
                            authorize()
                            connect()
                        }
                        reconnectionDelay = minReconnectionDelay
                        thisLogger().info("$devtoolHost: connected")
                        val reason = connection.await()
                        if (isActive) {
                            thisLogger().warn("$devtoolHost: connection closed, reconnecting after $reconnectionDelay milliseconds: $reason")
                        } else {
                            thisLogger().info("$devtoolHost: connection permanently closed: $reason")
                        }
                    } catch (t: Throwable) {
                        if (isActive) {
                            thisLogger().warn(
                                    "$devtoolHost: failed to connect, trying again after $reconnectionDelay milliseconds:",
                                    t
                            )
                        } else {
                            thisLogger().error("$devtoolHost: connection permanently closed:", t)
                        }
                    }
                    delay(reconnectionDelay)
                    reconnectionDelay = (reconnectionDelay * reconnectionDelayGrowFactor).toLong()
                    if (reconnectionDelay > maxReconnectionDelay) {
                        reconnectionDelay = maxReconnectionDelay
                    }
                }
            }
            lifetime.onTerminationOrNow {
                clients.remove(devtoolHost)
                job.cancel()
            }
            return@getOrPut client
        }
    }

    private class InvalidTokenException(message: String) : Exception(message)
}
