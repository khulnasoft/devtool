/**
 * Copyright (c) 2021 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { ContextURL, User, WorkspaceContext } from "@khulnasoft/devtool-protocol";
import { injectable } from "inversify";
import { IPrefixContextParser } from "./context-parser";

@injectable()
export class ImageBuildPrefixContextParser implements IPrefixContextParser {
    static PREFIX = ContextURL.IMAGEBUILD_PREFIX + "/";

    findPrefix(user: User, context: string): string | undefined {
        if (context.startsWith(ImageBuildPrefixContextParser.PREFIX)) {
            return ImageBuildPrefixContextParser.PREFIX;
        }
    }

    public async handle(user: User, prefix: string, context: WorkspaceContext): Promise<WorkspaceContext> {
        context.forceImageBuild = true;
        return context;
    }
}
