/**
 * Copyright (c) 2022 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

export interface BlockedRepository {
    id: number;
    urlRegexp: string;
    blockUser: boolean;
    blockFreeUsage: boolean;
    createdAt: string;
    updatedAt: string;
}
