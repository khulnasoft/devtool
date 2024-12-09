/**
 * Copyright (c) 2023 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { DBWithTracing, OneTimeSecretDB, TracedOneTimeSecretDB } from "@devtool/devtool-db/lib";
import { log } from "@khulnasoft/devtool-protocol/lib/util/logging";
import { inject, injectable } from "inversify";
import { Job } from "./runner";

@injectable()
export class OTSGarbageCollector implements Job {
    @inject(TracedOneTimeSecretDB) protected readonly oneTimeSecretDB: DBWithTracing<OneTimeSecretDB>;

    public name = "ots-gc";
    public frequencyMs = 5 * 60 * 1000; // every 5 minutes

    public async run(): Promise<number | undefined> {
        try {
            await this.oneTimeSecretDB.trace({}).pruneExpired();

            return undefined;
        } catch (err) {
            log.error("Failed to garbage collect OTS", err);
            throw err;
        }
    }
}
