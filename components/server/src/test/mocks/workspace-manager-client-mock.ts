/**
 * Copyright (c) 2023 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { User, Workspace, WorkspaceInstance } from "@khulnasoft/devtool-protocol";
import {
    WorkspaceClusterWoTLS,
    WorkspaceManagerConnectionInfo,
} from "@khulnasoft/devtool-protocol/lib/workspace-cluster";
import { PromisifiedWorkspaceManagerClient } from "@khulnasoft/ws-manager/lib";
import { IWorkspaceClusterStartSet, WorkspaceManagerClientProvider } from "@khulnasoft/ws-manager/lib/client-provider";
import { ChannelCredentials, Client, ClientOptions } from "@grpc/grpc-js";

export class WorkspaceManagerClientProviderMock extends WorkspaceManagerClientProvider {
    getStartClusterSets(
        user: User,
        workspace: Workspace,
        instance: WorkspaceInstance,
        region?: "" | "europe" | "north-america" | "south-america" | "africa" | "asia" | undefined,
    ): Promise<IWorkspaceClusterStartSet> {
        throw new Error("Method not implemented.");
    }
    get(name: string, grpcOptions?: object | undefined): Promise<PromisifiedWorkspaceManagerClient> {
        throw new Error("Method not implemented.");
    }
    getAllWorkspaceClusters(): Promise<WorkspaceClusterWoTLS[]> {
        throw new Error("Method not implemented.");
    }
    createConnection<T extends Client>(
        creator: new (address: string, credentials: ChannelCredentials, options?: ClientOptions | undefined) => T,
        info: WorkspaceManagerConnectionInfo,
        grpcOptions?: object | undefined,
    ): T {
        throw new Error("Method not implemented.");
    }
    dispose(): void {
        throw new Error("Method not implemented.");
    }
}
