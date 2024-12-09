// Copyright (c) 2022 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.devtool.devtoolprotocol.api;

import javax.websocket.CloseReason;
import java.util.concurrent.CompletionStage;
import java.util.concurrent.Future;

public interface DevtoolServerConnection extends Future<CloseReason>, CompletionStage<CloseReason> {
}
