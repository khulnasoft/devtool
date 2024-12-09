// Copyright (c) 2022 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.devtool.jetbrains.gateway

import com.intellij.icons.AllIcons
import com.intellij.ide.BrowserUtil
import com.intellij.ide.plugins.PluginManagerCore
import com.intellij.openapi.CompositeDisposable
import com.intellij.openapi.actionSystem.AnActionEvent
import com.intellij.openapi.application.ApplicationManager
import com.intellij.openapi.components.service
import com.intellij.openapi.diagnostic.thisLogger
import com.intellij.openapi.extensions.PluginId
import com.intellij.openapi.project.DumbAwareAction
import com.intellij.openapi.wm.impl.welcomeScreen.WelcomeScreenUIManager
import com.intellij.remoteDev.util.onTerminationOrNow
import com.intellij.ui.components.JBScrollPane
import com.intellij.ui.dsl.builder.*
import com.intellij.ui.layout.ComponentPredicate
import com.intellij.ui.layout.not
import com.intellij.util.ui.JBFont
import com.jetbrains.gateway.api.GatewayUI
import com.jetbrains.rd.util.lifetime.Lifetime
import com.jetbrains.rd.util.lifetime.LifetimeDefinition
import com.jetbrains.rd.util.lifetime.isAlive
import com.jetbrains.rd.util.lifetime.isNotAlive
import io.devtool.devtoolprotocol.api.entities.GetWorkspacesOptions
import io.devtool.devtoolprotocol.api.entities.WorkspaceInstance
import io.devtool.devtoolprotocol.api.entities.WorkspaceType
import io.devtool.jetbrains.auth.DevtoolAuthService
import io.devtool.jetbrains.icons.DevtoolIcons
import kotlinx.coroutines.*
import kotlinx.coroutines.channels.Channel
import kotlinx.coroutines.channels.actor
import kotlinx.coroutines.future.await
import org.apache.http.client.utils.URIBuilder
import java.time.OffsetDateTime
import java.time.ZonedDateTime
import java.time.temporal.ChronoUnit
import javax.swing.text.SimpleAttributeSet
import javax.swing.text.StyleConstants
import javax.swing.text.StyledDocument

@Suppress("UnstableApiUsage")
class DevtoolWorkspacesView(
    val lifetime: Lifetime
) {

    private val settings = service<DevtoolSettingsState>()

    private val loggedIn = object : ComponentPredicate() {
        override fun addListener(listener: (Boolean) -> Unit) {
            val toDispose = CompositeDisposable()
            toDispose.add(settings.addListener { listener(invoke()) })
            toDispose.add(DevtoolAuthService.addListener { listener(invoke()) })
            lifetime.onTerminationOrNow { toDispose.dispose() }
        }

        override fun invoke(): Boolean {
            return DevtoolAuthService.hasAccessToken(settings.devtoolHost)
        }
    }

    private val startWorkspaceView = DevtoolStartWorkspaceView(lifetime)

    private lateinit var workspacesPane: JBScrollPane
    val component = panel {
        indent {
            row {
                panel {
                    align(AlignY.CENTER)
                    for (i in 1..10) {
                        row {
                            label("")
                        }
                    }
                    row {
                        resizableRow()
                        icon(DevtoolIcons.Logo4x).align(AlignX.CENTER)
                    }
                    row {
                        text(
                            "Spin up fresh, automated dev environments for each task, in the cloud, in seconds.",
                            35
                        ).applyToComponent {
                            val attrs = SimpleAttributeSet()
                            StyleConstants.setAlignment(attrs, StyleConstants.ALIGN_CENTER)
                            (document as StyledDocument).setParagraphAttributes(
                                0,
                                document.length - 1,
                                attrs,
                                false
                            )
                        }.align(AlignX.CENTER)
                    }
                    row {
                        browserLink("Explore Devtool", "https://www.devtool.io").align(AlignX.CENTER)
                    }.bottomGap(BottomGap.MEDIUM)
                    row {
                        button("Connect in Browser") {
                            GlobalScope.launch {
                                DevtoolAuthService.authorize(settings.devtoolHost)
                            }
                        }.align(AlignX.CENTER)
                    }
                }
            }.visibleIf(loggedIn.not())

            val pluginVersion = PluginManagerCore.getPlugin(PluginId.getId("io.devtool.jetbrains.gateway"))?.version
            val pluginVersionLabel = if (pluginVersion?.contains("-local") == true) " (${pluginVersion})" else ""
            rowsRange {
                row {
                    icon(DevtoolIcons.Logo).gap(RightGap.SMALL)
                    label("Devtool${pluginVersionLabel}").applyToComponent {
                        this.font = JBFont.h3().asBold()
                    }
                    label("").resizableColumn().align(AlignX.FILL)
                    actionsButton(object :
                        DumbAwareAction("Dashboard", "Dashboard", AllIcons.Nodes.Servlet) {
                        override fun actionPerformed(e: AnActionEvent) {
                            BrowserUtil.browse("https://${settings.devtoolHost}")
                        }
                    }, object : DumbAwareAction("Usage", "Usage", AllIcons.Actions.DynamicUsages) {
                        override fun actionPerformed(e: AnActionEvent) {
                            BrowserUtil.browse("https://${settings.devtoolHost}/usage")
                        }
                    }, object : DumbAwareAction("Documentation", "Documentation", AllIcons.Toolwindows.Documentation) {
                        override fun actionPerformed(e: AnActionEvent) {
                            BrowserUtil.browse("https://www.devtool.io/docs/integrations/jetbrains-gateway")
                        }
                    }, object : DumbAwareAction("Feedback", "Feedback", AllIcons.Actions.IntentionBulb) {
                        override fun actionPerformed(e: AnActionEvent) {
                            BrowserUtil.browse("https://github.com/khulnasoft/devtool/issues/6576")
                        }
                    }, object : DumbAwareAction("Help", "Help", AllIcons.Actions.Help) {
                        override fun actionPerformed(e: AnActionEvent) {
                            BrowserUtil.browse("https://www.devtool.io/contact/support?subject=technical%20support")
                        }
                    }, object : DumbAwareAction("Log Out", "Log out", AllIcons.Actions.Exit) {
                        override fun actionPerformed(e: AnActionEvent) {
                            DevtoolAuthService.setAccessToken(settings.devtoolHost, null)
                        }
                    })
                    cell()
                }.topGap(TopGap.MEDIUM).bottomGap(BottomGap.SMALL)
                row {
                    cell(startWorkspaceView.component).align(AlignX.FILL)
                }.bottomGap(BottomGap.SMALL)
                row {
                    label("Recent Workspaces").bold()
                    label("").resizableColumn().align(AlignX.FILL)
                    actionButton(object :
                        DumbAwareAction("Refresh", "Refresh recent workspaces", AllIcons.Actions.Refresh) {
                        override fun actionPerformed(e: AnActionEvent) {
                            refresh()
                        }
                    })
                    cell()
                }
                row {
                    resizableRow()
                    workspacesPane = cell(JBScrollPane())
                        .resizableColumn()
                        .align(AlignX.FILL)
                        .align(AlignY.FILL)
                        .component
                    cell()
                }.bottomGap(BottomGap.SMALL)
            }.visibleIf(loggedIn)
        }
    }.apply {
        this.background = WelcomeScreenUIManager.getMainAssociatedComponentBackground()
    }

    val refresh = startUpdateLoop(lifetime, workspacesPane)

    init {
        refresh()
        loggedIn.addListener { refresh() }
    }

    @OptIn(DelicateCoroutinesApi::class, ObsoleteCoroutinesApi::class)
    private fun startUpdateLoop(lifetime: Lifetime, workspacesPane: JBScrollPane): () -> Unit {
        val updateJob = Job()
        lifetime.onTerminationOrNow { updateJob.cancel() }

        val updateActor = GlobalScope.actor<Void?>(updateJob, capacity = Channel.CONFLATED) {
            var updateLifetime: LifetimeDefinition? = null
            for (event in channel) {
                ensureActive()
                updateLifetime?.terminate()
                updateLifetime = lifetime.createNested()
                doUpdate(updateLifetime, workspacesPane)
            }
        }
        lifetime.onTerminationOrNow { updateActor.close() }

        return { updateActor.trySend(null) }
    }

    private fun getRelativeTimeSpan(creationTime: String): String {
        val fromDate = ZonedDateTime.parse(creationTime)
        val toDate = ZonedDateTime.now()
        val days = ChronoUnit.DAYS.between(fromDate, toDate)
        if (days > 0) {
            return "$days days ago"
        }
        val hours = ChronoUnit.HOURS.between(fromDate, toDate)
        if (hours > 0) {
            return "$hours hours ago"
        }
        val minutes = ChronoUnit.MINUTES.between(fromDate, toDate)
        if (minutes > 0) {
            return "$minutes minutes ago"
        }
        return "a few seconds ago"
    }

    private fun doUpdate(updateLifetime: Lifetime, workspacesPane: JBScrollPane) {
        val devtoolHost = settings.devtoolHost
        if (!DevtoolAuthService.hasAccessToken(devtoolHost)) {
            ApplicationManager.getApplication().invokeLater {
                if (updateLifetime.isAlive) {
                    workspacesPane.viewport.view = panel {
                        row {
                            comment("Loading...")
                        }
                    }
                }
            }
            return
        }
        val job = GlobalScope.launch {
            val client = service<DevtoolConnectionService>().obtainClient(devtoolHost)
            val workspaces = client.server.getWorkspaces(GetWorkspacesOptions().apply {
                this.limit = 20
            }).await()
            val workspacesMap = workspaces.associateBy { it.workspace.id }.toMutableMap()
            fun updateView() {
                val view = panel {
                    val sortedInfos = workspacesMap.values.toMutableList()
                        .sortedByDescending {
                            val creationTime = it.latestInstance?.creationTime ?: it.workspace.creationTime
                            try {
                                if (creationTime != null) {
                                    OffsetDateTime.parse(creationTime)
                                } else {
                                    null
                                }
                            } catch (e: Throwable) {
                                thisLogger().error(
                                    "$devtoolHost: ${it.workspace.id}: failed to parse creation time",
                                    e
                                )
                                null
                            }
                        }
                    for (info in sortedInfos) {
                        if (info.latestInstance == null || info.workspace.type != WorkspaceType.regular) {
                            continue
                        }
                        indent {
                            row {
                                var canConnect = false
                                icon(
                                    if (info.latestInstance.status.phase == "running") {
                                        canConnect = true
                                        DevtoolIcons.Running
                                    } else if (info.latestInstance.status.phase == "stopped") {
                                        if (info.latestInstance.status.conditions.failed.isNullOrBlank()) {
                                            DevtoolIcons.Stopped
                                        } else {
                                            DevtoolIcons.Failed
                                        }
                                    } else if (info.latestInstance.status.phase == "interrupted") {
                                        DevtoolIcons.Failed
                                    } else if (info.latestInstance.status.phase == "unknown") {
                                        DevtoolIcons.Failed
                                    } else {
                                        canConnect = true
                                        DevtoolIcons.Starting
                                    }
                                ).gap(RightGap.SMALL)
                                panel {
                                    val contextUrlRow = row {
                                        browserLink(info.workspace.id, info.latestInstance.ideUrl)
                                    }
                                    info.workspace.context.normalizedContextURL.ifPresent {
                                        contextUrlRow.rowComment("<a href='$it'>$it</a>")
                                    }
                                }
                                label("").resizableColumn().align(AlignX.FILL)
                                panel {
                                    val repo = info.latestInstance.status.repo
                                    val changes = repo?.let {
                                        it.totalUncommitedFiles + it.totalUntrackedFiles + it.totalUnpushedCommits
                                    } ?: 0
                                    row {
                                        info.workspace.context.ref.ifPresentOrElse({ label(it) }, { label("(detached)") })
                                    }.rowComment(
                                        when {
                                            changes == 1 -> "<b>$changes Change</b>"
                                            changes > 0 -> "<b>$changes Changes</b>"
                                            else -> "No Changes"
                                        }
                                    )
                                }
                                label(getRelativeTimeSpan(info.latestInstance.creationTime))
                                button("Connect") {
                                    if (!canConnect) {
                                        val startUrl = URIBuilder()
                                            .setScheme("https")
                                            .setHost(devtoolHost)
                                            .setPath("start")
                                            .setFragment(info.workspace.id)
                                            .build()
                                            .toString()
                                        BrowserUtil.browse(startUrl)
                                    } else {
                                        GatewayUI.getInstance().connect(
                                            mapOf(
                                                "devtoolHost" to devtoolHost,
                                                "workspaceId" to info.workspace.id
                                            )
                                        )
                                    }
                                }
                                cell()
                            }.layout(RowLayout.PARENT_GRID)
                        }
                    }
                }
                ApplicationManager.getApplication().invokeLater {
                    if (updateLifetime.isAlive) {
                        workspacesPane.viewport.view = view
                    }
                }
            }
            updateView()
            val updates = client.listenToWorkspace(updateLifetime, "*")
            for (update in updates) {
                if (updateLifetime.isNotAlive) {
                    return@launch
                }
                var info = workspacesMap[update.workspaceId]
                if (info == null) {
                    try {
                        info = client.syncWorkspace(update.workspaceId)
                    } catch (t: Throwable) {
                        thisLogger().error("$devtoolHost: ${update.workspaceId}: failed to sync", t)
                        continue
                    }
                    if (info.workspace.type == WorkspaceType.regular) {
                        workspacesMap[update.workspaceId] = info
                    }
                } else if (WorkspaceInstance.isUpToDate(info.latestInstance, update)) {
                    continue
                } else {
                    info.latestInstance = update
                }
                updateView()
            }
        }
        updateLifetime.onTerminationOrNow { job.cancel() }
    }
}