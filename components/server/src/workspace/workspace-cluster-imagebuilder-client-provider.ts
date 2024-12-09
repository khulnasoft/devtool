/**
 * Copyright (c) 2022 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { User, Workspace, WorkspaceInstance } from "@khulnasoft/devtool-protocol";
import { defaultGRPCOptions } from "@khulnasoft/devtool-protocol/lib/util/grpc";
import { WorkspaceRegion } from "@khulnasoft/devtool-protocol/lib/workspace-cluster";
import {
    ImageBuilderClient,
    ImageBuilderClientProvider,
    PromisifiedImageBuilderClient,
} from "@devtool/image-builder/lib";
import { WorkspaceManagerClientProvider } from "@devtool/ws-manager/lib/client-provider";
import {
    WorkspaceManagerClientProviderCompositeSource,
    WorkspaceManagerClientProviderSource,
} from "@devtool/ws-manager/lib/client-provider-source";
import { inject, injectable } from "inversify";

@injectable()
export class WorkspaceClusterImagebuilderClientProvider implements ImageBuilderClientProvider {
    @inject(WorkspaceManagerClientProviderCompositeSource)
    protected readonly source: WorkspaceManagerClientProviderSource;
    @inject(WorkspaceManagerClientProvider) protected readonly clientProvider: WorkspaceManagerClientProvider;

    // gRPC connections can be used concurrently, even across services.
    // Thus it makes sense to cache them rather than create a new connection for each request.
    protected readonly connectionCache = new Map<string, ImageBuilderClient>();

    async getClient(
        user: User,
        workspace?: Workspace,
        instance?: WorkspaceInstance,
        region?: WorkspaceRegion,
    ): Promise<PromisifiedImageBuilderClient> {
        const clusters = await this.clientProvider.getStartClusterSets(user, workspace, instance, region);
        for await (const cluster of clusters) {
            const info = await this.source.getWorkspaceCluster(cluster.installation);
            if (!info) {
                continue;
            }

            let client = this.connectionCache.get(info.name);
            if (!client) {
                client = this.clientProvider.createConnection(ImageBuilderClient, info, defaultGRPCOptions);
                this.connectionCache.set(info.name, client);
            }
            return new PromisifiedImageBuilderClient(client, []);
        }

        throw new Error("no image-builder available");
    }
}
