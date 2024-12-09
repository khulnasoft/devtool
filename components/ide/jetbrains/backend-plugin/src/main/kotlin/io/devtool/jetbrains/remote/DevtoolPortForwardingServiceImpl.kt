// Copyright (c) 2024 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.devtool.jetbrains.remote

import com.jetbrains.rd.util.lifetime.Lifetime
import com.jetbrains.rd.util.threading.coroutines.launch
import io.devtool.jetbrains.remote.AbstractDevtoolPortForwardingService
import kotlinx.coroutines.CoroutineScope

@Suppress("UnstableApiUsage")
class DevtoolPortForwardingServiceImpl : AbstractDevtoolPortForwardingService() {
    override fun runJob(lifetime: Lifetime, block: suspend CoroutineScope.() -> Unit) = lifetime.launch { block() }
}
