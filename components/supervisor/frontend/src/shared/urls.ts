/**
 * Copyright (c) 2020 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { DevtoolHostUrl } from "@khulnasoft/devtool-protocol/lib/util/devtool-host-url";

export const workspaceUrl = new DevtoolHostUrl(window.location.href);

export const serverUrl = workspaceUrl.withoutWorkspacePrefix();

export const startUrl = serverUrl.asStart(workspaceUrl.workspaceId);
