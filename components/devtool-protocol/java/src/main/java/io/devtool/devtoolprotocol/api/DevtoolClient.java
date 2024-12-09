// Copyright (c) 2021 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.devtool.devtoolprotocol.api;

import io.devtool.devtoolprotocol.api.entities.WorkspaceInstance;
import org.eclipse.lsp4j.jsonrpc.services.JsonNotification;

public class DevtoolClient {

    private DevtoolServer server;

    public void connect(DevtoolServer server) {
        this.server = server;
    }

    public DevtoolServer getServer() {
        if (this.server == null) {
            throw new IllegalStateException("not connected");
        }
        return this.server;
    }

    public void notifyConnect() {
    }

    @JsonNotification
    public void onInstanceUpdate(WorkspaceInstance instance) {

    }
}
