/**
 * Copyright (c) 2021 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

(import './dashboards/SLOs/workspace-startup-time.libsonnet') +
{
  grafanaDashboards+:: {
    // Import raw json files here.
    // Example:
    // 'my-new-dashboard.json': (import 'dashboards/components/new-component.json'),
    'devtool-cluster-autoscaler-k3s.json': (import 'dashboards/devtool-cluster-autoscaler-k3s.json'),
    'devtool-node-resource-metrics.json': (import 'dashboards/devtool-node-resource-metrics.json'),
    'devtool-grpc-server.json': (import 'dashboards/devtool-grpc-server.json'),
    'devtool-grpc-client.json': (import 'dashboards/devtool-grpc-client.json'),
    'devtool-connect-server.json': (import 'dashboards/devtool-connect-server.json'),
    'devtool-overview.json': (import 'dashboards/devtool-overview.json'),
    'devtool-nodes-overview.json': (import 'dashboards/devtool-nodes-overview.json'),
    'devtool-admin-node.json': (import 'dashboards/devtool-admin-node.json'),
    'devtool-admin-workspace.json': (import 'dashboards/devtool-admin-workspace.json'),
    'devtool-applications.json': (import 'dashboards/devtool-applications.json'),
    'redis.json': (import 'dashboards/redis.json')
  },
}
