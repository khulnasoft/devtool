/**
 * Copyright (c) 2023 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { PartialMessage } from "@bufbuild/protobuf";
import { CallOptions, PromiseClient } from "@connectrpc/connect";
import { PrebuildService } from "@khulnasoft/public-api/lib/devtool/v1/prebuild_connect";
import {
    StartPrebuildRequest,
    GetPrebuildRequest,
    ListPrebuildsRequest,
    WatchPrebuildRequest,
    StartPrebuildResponse,
    GetPrebuildResponse,
    ListPrebuildsResponse,
    WatchPrebuildResponse,
    CancelPrebuildRequest,
    CancelPrebuildResponse,
    ListOrganizationPrebuildsRequest,
    ListOrganizationPrebuildsResponse,
} from "@khulnasoft/public-api/lib/devtool/v1/prebuild_pb";
import { getDevtoolService } from "./service";
import { converter } from "./public-api";
import { PrebuildWithStatus } from "@khulnasoft/devtool-protocol";
import { generateAsyncGenerator } from "@khulnasoft/devtool-protocol/lib/generate-async-generator";
import { ApplicationError, ErrorCodes } from "@khulnasoft/devtool-protocol/lib/messaging/error";

export class JsonRpcPrebuildClient implements PromiseClient<typeof PrebuildService> {
    async startPrebuild(
        request: PartialMessage<StartPrebuildRequest>,
        options?: CallOptions,
    ): Promise<StartPrebuildResponse> {
        if (!request.configurationId) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "configurationId is required");
        }
        const result = await getDevtoolService().server.triggerPrebuild(request.configurationId, request.gitRef || null);
        return new StartPrebuildResponse({
            prebuildId: result.prebuildId,
        });
    }

    async cancelPrebuild(
        request: PartialMessage<CancelPrebuildRequest>,
        options?: CallOptions,
    ): Promise<CancelPrebuildResponse> {
        const response = await this.getPrebuild(request, options);
        await getDevtoolService().server.cancelPrebuild(response.prebuild!.configurationId, response.prebuild!.id);
        return new CancelPrebuildResponse();
    }

    get devtoolHost(): string {
        return window.location.protocol + "//" + window.location.host;
    }

    async getPrebuild(
        request: PartialMessage<GetPrebuildRequest>,
        options?: CallOptions,
    ): Promise<GetPrebuildResponse> {
        if (!request.prebuildId) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "prebuildId is required");
        }
        const result = await getDevtoolService().server.getPrebuild(request.prebuildId);
        if (!result) {
            throw new ApplicationError(ErrorCodes.NOT_FOUND, `prebuild ${request.prebuildId} not found`);
        }
        return new GetPrebuildResponse({
            prebuild: converter.toPrebuild(this.devtoolHost, result),
        });
    }

    async listPrebuilds(
        request: PartialMessage<ListPrebuildsRequest>,
        options?: CallOptions,
    ): Promise<ListPrebuildsResponse> {
        if (request.workspaceId) {
            const pbws = await getDevtoolService().server.findPrebuildByWorkspaceID(request.workspaceId);
            if (pbws) {
                const prebuild = await getDevtoolService().server.getPrebuild(pbws.id);
                if (prebuild) {
                    return new ListPrebuildsResponse({
                        prebuilds: [converter.toPrebuild(this.devtoolHost, prebuild)],
                    });
                }
            }
            return new ListPrebuildsResponse({
                prebuilds: [],
            });
        }
        if (!request.configurationId) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "configurationId is required");
        }
        const result = await getDevtoolService().server.findPrebuilds({
            projectId: request.configurationId,
            branch: request.gitRef || undefined,
            limit: request.pagination?.pageSize || undefined,
        });
        return new ListPrebuildsResponse({
            prebuilds: converter.toPrebuilds(this.devtoolHost, result),
        });
    }

    async *watchPrebuild(
        request: PartialMessage<WatchPrebuildRequest>,
        options?: CallOptions,
    ): AsyncIterable<WatchPrebuildResponse> {
        if (!options?.signal) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "signal is required");
        }
        if (!request.scope?.value) {
            throw new ApplicationError(ErrorCodes.BAD_REQUEST, "scope is required");
        }
        if (request.scope.case === "prebuildId") {
            const prebuild = await this.getPrebuild({ prebuildId: request.scope.value }, options);
            yield new WatchPrebuildResponse({
                prebuild: prebuild.prebuild,
            });
        }
        const it = generateAsyncGenerator<PrebuildWithStatus>(
            (queue) => {
                try {
                    const dispose = getDevtoolService().registerClient({
                        onPrebuildUpdate: (prebuild) => {
                            queue.push(prebuild);
                        },
                    });
                    return () => {
                        dispose.dispose();
                    };
                } catch (e) {
                    queue.fail(e);
                }
            },
            { signal: options.signal },
        );
        for await (const pb of it) {
            if (request.scope.case === "prebuildId") {
                if (pb.info.id !== request.scope.value) {
                    continue;
                }
            } else if (pb.info.projectId !== request.scope.value) {
                continue;
            }
            const prebuild = converter.toPrebuild(this.devtoolHost, pb);
            if (prebuild) {
                yield new WatchPrebuildResponse({ prebuild });
            }
        }
    }

    async listOrganizationPrebuilds(
        request: PartialMessage<ListOrganizationPrebuildsRequest>,
    ): Promise<ListOrganizationPrebuildsResponse> {
        throw new ApplicationError(ErrorCodes.UNIMPLEMENTED, "Not implemented (for jrpc)");
    }
}
