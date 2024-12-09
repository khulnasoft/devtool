/**
 * Copyright (c) 2024 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { AuthProviderEntry } from "@devtool/devtool-protocol";

export function oauthUrls(customConfig: Required<AuthProviderEntry.OAuth2CustomConfig>) {
    return {
        authorizationUrl: customConfig.authorizationUrl,
        tokenUrl: customConfig.tokenUrl,
        settingsUrl: "",
    };
}
