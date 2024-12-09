/**
 * Copyright (c) 2023 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import * as crypto from "crypto";
import { DBDevtoolToken, UserDB } from "@devtool/devtool-db/lib";
import { DevtoolToken, DevtoolTokenType } from "@khulnasoft/devtool-protocol";
import { log } from "@khulnasoft/devtool-protocol/lib/util/logging";
import { inject, injectable } from "inversify";
import { Authorizer } from "../authorization/authorizer";

@injectable()
export class DevtoolTokenService {
    constructor(
        @inject(UserDB) private readonly userDB: UserDB,
        @inject(Authorizer) private readonly auth: Authorizer,
    ) {}

    async getDevtoolTokens(requestorId: string, userId: string): Promise<DevtoolToken[]> {
        await this.auth.checkPermissionOnUser(requestorId, "read_tokens", userId);
        const devtoolTokens = await this.userDB.findAllDevtoolTokensOfUser(userId);
        return devtoolTokens;
    }

    async generateNewDevtoolToken(
        requestorId: string,
        userId: string,
        options: { name?: string; type: DevtoolTokenType; scopes?: string[] },
        oldPermissionCheck?: (dbToken: DBDevtoolToken) => Promise<void>, // @deprecated
    ): Promise<string> {
        await this.auth.checkPermissionOnUser(requestorId, "write_tokens", userId);
        const token = crypto.randomBytes(30).toString("hex");
        const tokenHash = crypto.createHash("sha256").update(token, "utf8").digest("hex");
        const dbToken: DBDevtoolToken = {
            tokenHash,
            name: options.name,
            type: options.type,
            userId,
            scopes: options.scopes || [],
            created: new Date().toISOString(),
        };
        if (oldPermissionCheck) {
            await oldPermissionCheck(dbToken);
        }
        await this.userDB.storeDevtoolToken(dbToken);
        return token;
    }

    async findDevtoolToken(requestorId: string, userId: string, tokenHash: string): Promise<DevtoolToken | undefined> {
        await this.auth.checkPermissionOnUser(requestorId, "read_tokens", userId);
        let token: DevtoolToken | undefined;
        try {
            token = await this.userDB.findDevtoolTokensOfUser(userId, tokenHash);
        } catch (error) {
            log.error({ userId }, "failed to resolve devtool token: ", error);
        }
        return token;
    }

    async deleteDevtoolToken(
        requestorId: string,
        userId: string,
        tokenHash: string,
        oldPermissionCheck?: (token: DevtoolToken) => Promise<void>, // @deprecated
    ): Promise<void> {
        await this.auth.checkPermissionOnUser(requestorId, "write_tokens", userId);
        const existingTokens = await this.getDevtoolTokens(requestorId, userId);
        const tkn = existingTokens.find((token) => token.tokenHash === tokenHash);
        if (!tkn) {
            throw new Error(`User ${requestorId} tries to delete a token ${tokenHash} that does not exist.`);
        }
        if (oldPermissionCheck) {
            await oldPermissionCheck(tkn);
        }
        await this.userDB.deleteDevtoolToken(tokenHash);
    }
}
