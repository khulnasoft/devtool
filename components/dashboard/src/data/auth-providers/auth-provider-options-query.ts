/**
 * Copyright (c) 2024 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { AuthProviderType } from "@khulnasoft/public-api/lib/devtool/v1/authprovider_pb";
import { isDevtoolIo } from "../../utils";
import { useMemo } from "react";

const optionsForPAYG = [
    { type: AuthProviderType.GITHUB, label: "GitHub" },
    { type: AuthProviderType.GITLAB, label: "GitLab" },
    { type: AuthProviderType.BITBUCKET_SERVER, label: "Bitbucket Server" },
    { type: AuthProviderType.BITBUCKET, label: "Bitbucket Cloud" },
];

const optionsForEnterprise = [...optionsForPAYG, { type: AuthProviderType.AZURE_DEVOPS, label: "Azure DevOps" }];

export const isSupportAzureDevOpsIntegration = () => {
    return isDevtoolIo();
};

export const useAuthProviderOptionsQuery = (isOrgLevel: boolean) => {
    return useMemo(() => {
        const isPAYG = isDevtoolIo();
        // Azure DevOps is not supported for PAYG users and is only available for org-level integrations
        // because auth flow is identified by auth provider's host, which will always be `dev.azure.com`
        //
        // Don't remove this until we can setup an generial application for Azure DevOps (investigate needed)
        if (isPAYG || !isOrgLevel) {
            return optionsForPAYG;
        }
        return optionsForEnterprise;
    }, [isOrgLevel]);
};
