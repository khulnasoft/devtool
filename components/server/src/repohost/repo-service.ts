/**
 * Copyright (c) 2020 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { User } from "@khulnasoft/devtool-protocol";
import { injectable } from "inversify";

@injectable()
export class RepositoryService {
    async isDevtoolWebhookEnabled(user: User, cloneUrl: string): Promise<boolean> {
        throw new Error("unsupported");
    }
}
