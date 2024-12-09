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
    'devtool-component-blobserve.json': (import 'dashboards/components/blobserve.json'),
    'devtool-component-openvsx-proxy.json': (import 'dashboards/components/openvsx-proxy.json'),
    'devtool-component-openvsx-mirror.json': (import 'dashboards/components/openvsx-mirror.json'),
    'devtool-component-ssh-gateway.json': (import 'dashboards/components/ssh-gateway.json'),
    'devtool-component-supervisor.json': (import 'dashboards/components/supervisor.json'),
    'devtool-component-jb.json': (import 'dashboards/components/jb.json'),
    'devtool-component-browser-overview.json': (import 'dashboards/components/browser-overview.json'),
    'devtool-component-code-browser.json': (import 'dashboards/components/code-browser.json'),
    'devtool-component-ide-startup-time.json': (import 'dashboards/components/ide-startup-time.json'),
    'devtool-component-ide-service.json': (import 'dashboards/components/ide-service.json'),
    'devtool-component-local-ssh.json': (import 'dashboards/components/local-ssh.json'),
  },
}
