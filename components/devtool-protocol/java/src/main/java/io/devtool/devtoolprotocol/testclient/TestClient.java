// Copyright (c) 2021 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.devtool.devtoolprotocol.testclient;

import io.devtool.devtoolprotocol.api.DevtoolClient;
import io.devtool.devtoolprotocol.api.DevtoolServer;
import io.devtool.devtoolprotocol.api.DevtoolServerLauncher;
import io.devtool.devtoolprotocol.api.entities.SendHeartBeatOptions;
import io.devtool.devtoolprotocol.api.entities.User;

import java.util.Collections;

public class TestClient {
    public static void main(String[] args) throws Exception {
        String uri = "wss://devtool.io/api/v1";
        String token = "CHANGE-ME";
        String origin = "https://CHANGE-ME.devtool.io/";

        DevtoolClient client = new DevtoolClient();
        DevtoolServerLauncher.create(client).listen(uri, origin, token, "Test", "Test", Collections.emptyList(), null);
        DevtoolServer devtoolServer = client.getServer();
        User user = devtoolServer.getLoggedInUser().join();
        System.out.println("logged in user:" + user);

        Void result = devtoolServer
                .sendHeartBeat(new SendHeartBeatOptions("CHANGE-ME", false)).join();
        System.out.println("send heart beat:" + result);
    }
}
