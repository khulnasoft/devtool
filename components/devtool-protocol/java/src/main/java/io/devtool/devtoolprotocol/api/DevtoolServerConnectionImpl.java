// Copyright (c) 2022 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.devtool.devtoolprotocol.api;

import javax.websocket.CloseReason;
import javax.websocket.Session;
import java.io.IOException;
import java.util.concurrent.CompletableFuture;
import java.util.logging.Level;
import java.util.logging.Logger;

public class DevtoolServerConnectionImpl extends CompletableFuture<CloseReason> implements DevtoolServerConnection {

    public static final Logger LOG = Logger.getLogger(DevtoolServerConnectionImpl.class.getName());

    private final String devtoolHost;

    private Session session;

    public DevtoolServerConnectionImpl(String devtoolHost) {
        this.devtoolHost = devtoolHost;
    }

    public void setSession(Session session) {
        this.session = session;
    }

    @Override
    public boolean cancel(boolean mayInterruptIfRunning) {
        Session session = this.session;
        this.session = null;
        if (session != null) {
            try {
                session.close();
            } catch (IOException e) {
                LOG.log(Level.WARNING, devtoolHost + ": failed to close connection:", e);
            }
        }
        return super.cancel(mayInterruptIfRunning);
    }
}
