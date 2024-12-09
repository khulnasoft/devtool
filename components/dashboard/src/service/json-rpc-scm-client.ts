/**
 * Copyright (c) 2023 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { PartialMessage } from "@bufbuild/protobuf";
import { PromiseClient } from "@connectrpc/connect";
import { SCMService } from "@devtool/public-api/lib/devtool/v1/scm_connect";
import { converter } from "./public-api";
import { getDevtoolService } from "./service";
import { ApplicationError, ErrorCodes } from "@devtool/devtool-protocol/lib/messaging/error";
import {
    SearchSCMTokensRequest,
    SearchSCMTokensResponse,
    GuessTokenScopesRequest,
    SearchRepositoriesRequest,
    ListSuggestedRepositoriesRequest,
    ListSuggestedRepositoriesResponse,
    SearchRepositoriesResponse,
    GuessTokenScopesResponse,
} from "@devtool/public-api/lib/devtool/v1/scm_pb";

export class JsonRpcScmClient implements PromiseClient<typeof SCMService> {
    async searchSCMTokens({ host }: PartialMessage<SearchSCMTokensRequest>): Promise<SearchSCMTokensResponse> {
        if (!host) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "host is required");
        }
        const response = new SearchSCMTokensResponse();
        const token = await getDevtoolService().server.getToken({ host });
        if (token) {
            response.tokens.push(converter.toSCMToken(token));
        }
        return response;
    }

    async guessTokenScopes({
        gitCommand,
        host,
        repoUrl,
    }: PartialMessage<GuessTokenScopesRequest>): Promise<GuessTokenScopesResponse> {
        if (!host) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "host is required");
        }
        const response = await getDevtoolService().server.guessGitTokenScopes({
            gitCommand: gitCommand || "",
            host,
            repoUrl: repoUrl || "",
        });
        return new GuessTokenScopesResponse({
            message: response.message,
            scopes: response.scopes || [],
        });
    }

    async searchRepositories(request: PartialMessage<SearchRepositoriesRequest>): Promise<SearchRepositoriesResponse> {
        const { limit, searchString } = request;
        if (!searchString) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "searchString is required");
        }
        const repos = await getDevtoolService().server.searchRepositories({ searchString, limit });
        return new SearchRepositoriesResponse({
            repositories: repos.map((r) => converter.toSuggestedRepository(r)),
        });
    }

    async listSuggestedRepositories(
        request: PartialMessage<ListSuggestedRepositoriesRequest>,
    ): Promise<ListSuggestedRepositoriesResponse> {
        const { organizationId } = request;
        if (!organizationId) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "organizationId is required");
        }
        const repos = await getDevtoolService().server.getSuggestedRepositories(organizationId);
        return new SearchRepositoriesResponse({
            repositories: repos.map((r) => converter.toSuggestedRepository(r)),
        });
    }
}
