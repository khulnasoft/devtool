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
    'devtool-component-dashboard.json': (import 'dashboards/components/dashboard.json'),
    'devtool-component-db.json': (import 'dashboards/components/db.json'),
    'devtool-component-ws-manager-bridge.json': (import 'dashboards/components/ws-manager-bridge.json'),
    'devtool-component-proxy.json': (import 'dashboards/components/proxy.json'),
    'devtool-component-server.json': (import 'dashboards/components/server.json'),
    'devtool-component-server-garbage-collector.json': (import 'dashboards/components/server-garbage-collector.json'),
    'devtool-component-usage.json': (import 'dashboards/components/usage.json'),
    'devtool-slo-login.json': (import 'dashboards/SLOs/login.json'),
    'devtool-meta-overview.json': (import 'dashboards/components/meta-overview.json'),
    'devtool-meta-services.json': (import 'dashboards/components/meta-services.json'),
    'devtool-components-spicedb.json': (import 'dashboards/components/spicedb.json'),
  },
}
