/**
 * Copyright (c) 2023 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { LinkedInProfile } from "@khulnasoft/devtool-protocol";

export const LinkedInProfileDB = Symbol("LinkedInProfileDB");
export interface LinkedInProfileDB {
    storeProfile(userId: string, profile: LinkedInProfile): Promise<void>;
}
