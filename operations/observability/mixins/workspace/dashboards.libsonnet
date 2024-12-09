/**
 * Copyright (c) 2021 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

{
  grafanaDashboards+:: {
    // Import raw json files here.
    // Example:
    // 'my-new-dashboard.json': (import 'dashboards/components/new-component.json'),
    'devtool-component-agent-smith.json': (import 'dashboards/components/agent-smith.json'),
    'devtool-component-content-service.json': (import 'dashboards/components/content-service.json'),
    'devtool-component-registry-facade.json': (import 'dashboards/components/registry-facade.json'),
    'devtool-component-ws-daemon.json': (import 'dashboards/components/ws-daemon.json'),
    'devtool-component-ws-manager-mk2.json': (import 'dashboards/components/ws-manager-mk2.json'),
    'devtool-component-ws-proxy.json': (import 'dashboards/components/ws-proxy.json'),
    'devtool-workspace-success-criteria.json': (import 'dashboards/success-criteria.json'),
    'devtool-workspace-coredns.json': (import 'dashboards/coredns.json'),
    'devtool-node-swap.json': (import 'dashboards/node-swap.json'),
    'devtool-node-ephemeral-storage.json': (import 'dashboards/ephemeral-storage.json'),
    'devtool-node-problem-detector.json': (import 'dashboards/node-problem-detector.json'),
    'devtool-network-limiting.json': (import 'dashboards/network-limiting.json'),
    'devtool-component-image-builder.json': (import 'dashboards/components/image-builder.json'),
    'devtool-psi.json': (import 'dashboards/node-psi.json'),
    'devtool-workspace-psi.json': (import 'dashboards/workspace-psi.json'),
    'devtool-workspace-registry-facade-blobsource.json': (import 'dashboards/registry-facade-blobsource.json'),
  },
}
