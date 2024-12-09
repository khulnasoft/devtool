/**
 * Copyright (c) 2023 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { ListUsageRequest, ListUsageResponse } from "@khulnasoft/devtool-protocol/lib/usage";
import { useQuery } from "@tanstack/react-query";
import { getDevtoolService } from "../../service/service";

export function useListUsage(request: ListUsageRequest) {
    const query = useQuery<ListUsageResponse, Error>(
        ["usage", request],
        () => {
            return getDevtoolService().server.listUsage(request);
        },
        {
            cacheTime: 1000 * 60 * 1, // 1 minutes
            staleTime: 1000 * 60 * 1, // 1 minutes
            retry: false,
        },
    );
    return query;
}
