// Copyright (c) 2022 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.devtool.jetbrains.gateway

import com.intellij.openapi.Disposable
import com.intellij.openapi.components.PersistentStateComponent
import com.intellij.openapi.components.State
import com.intellij.openapi.components.Storage
import com.intellij.util.EventDispatcher
import com.intellij.util.xmlb.XmlSerializerUtil
import java.net.URL
import java.util.*

@State(
    name = "io.devtool.jetbrains.gateway.DevtoolSettingsState",
    storages = [Storage("devtool.xml")]
)
class DevtoolSettingsState : PersistentStateComponent<DevtoolSettingsState> {

    var devtoolHost: String = "devtool.io"
        set(value) {
            if (value.isNullOrBlank()) {
                return
            }
            val devtoolHost = try {
                URL(value.trim()).host
            } catch (t: Throwable) {
                value.trim()
            }
            if (devtoolHost == field) {
                return
            }
            field = devtoolHost
            dispatcher.multicaster.didChange()
        }

    var forceHttpTunnel: Boolean = false
        set(value) {
            if (value == field) {
                return
            }
            field = value
            dispatcher.multicaster.didChange()
        }

    var additionalHeartbeat: Boolean = false
        set(value) {
            if (value == field) {
                return
            }
            field = value
            dispatcher.multicaster.didChange()
        }

    private interface Listener : EventListener {
        fun didChange()
    }

    private val dispatcher = EventDispatcher.create(Listener::class.java)
    fun addListener(listener: () -> Unit): Disposable {
        val internalListener = object : Listener {
            override fun didChange() {
                listener()
            }
        }
        dispatcher.addListener(internalListener)
        return Disposable { dispatcher.removeListener(internalListener) }
    }

    override fun getState(): DevtoolSettingsState {
        return this
    }

    override fun loadState(state: DevtoolSettingsState) {
        XmlSerializerUtil.copyBean(state, this)
    }
}
