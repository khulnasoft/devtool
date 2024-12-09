/**
 * Copyright (c) 2023 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { FC } from "react";
import { ConfigurationNameForm } from "./general/ConfigurationName";
import { RemoveConfiguration } from "./general/RemoveConfiguration";
import { Configuration } from "@devtool/public-api/lib/devtool/v1/configuration_pb";

type Props = {
    configuration: Configuration;
};
export const ConfigurationDetailGeneral: FC<Props> = ({ configuration }) => {
    return (
        <>
            <ConfigurationNameForm configuration={configuration} />
            <RemoveConfiguration configuration={configuration} />
        </>
    );
};
