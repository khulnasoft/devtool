/**
 * Copyright (c) 2020 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { User, Token } from "@khulnasoft/devtool-protocol";
import {
    RepositoryNotFoundError,
    UnauthorizedRepositoryAccessError,
} from "@khulnasoft/public-api-common/lib/public-api-errors";
import { RepositoryUnauthorizedError } from "@khulnasoft/public-api/lib/devtool/v1/error_pb";

export namespace NotFoundError {
    export async function create(
        token: Token | undefined,
        user: User,
        host: string,
        owner: string,
        repoName: string,
        errorMessage: string = "Repository not found.",
    ) {
        const lastUpdate = (token && token.updateDate) ?? "";
        const userScopes = token ? [...token.scopes] : [];

        const userIsOwner = owner == user.name; // TODO: shouldn't this be a comparison with `identity.authName`?
        return new RepositoryNotFoundError({
            host,
            owner,
            repoName,
            userIsOwner,
            userScopes,
            lastUpdate,
            errorMessage,
        });
    }
    export function is(error: any): error is RepositoryNotFoundError {
        return error instanceof RepositoryNotFoundError;
    }
}

export namespace UnauthorizedError {
    export function create(props: Partial<RepositoryUnauthorizedError>) {
        return new UnauthorizedRepositoryAccessError(props);
    }
    export function is(error: any): error is UnauthorizedRepositoryAccessError {
        return error instanceof UnauthorizedRepositoryAccessError;
    }
}
