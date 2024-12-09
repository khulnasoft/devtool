/**
 * Copyright (c) 2024 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { getExperimentsClientForBackend } from "@devtool/devtool-protocol/lib/experiments/configcat-server";

export async function getFeatureFlagEnableExperimentalJBTB(userId: string): Promise<boolean> {
    return getExperimentsClientForBackend().getValueAsync("enable_experimental_jbtb", false, {
        user: { id: userId },
    });
}
