/**
 * Copyright (c) 2023 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { FC } from "react";
import { Configuration } from "@khulnasoft/public-api/lib/devtool/v1/configuration_pb";
import { ConfigurationWorkspaceClassesOptions } from "./workspaces/ConfigurationWorkspaceClassesOptions";

type Props = {
    configuration: Configuration;
};
export const ConfigurationDetailWorkspaces: FC<Props> = ({ configuration }) => {
    return <ConfigurationWorkspaceClassesOptions configuration={configuration} />;
};
