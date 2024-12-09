/**
 * Copyright (c) 2024 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { useContext } from "react";
import { useIsDataOps } from "../data/featureflag-query";
import { ThemeContext } from "../theme-context";
import devtoolIcon from "../icons/devtool.svg";
import dataOpsIconLight from "../icons/dataops-light.svg";
import dataOpsIconDark from "../icons/dataops-dark.svg";

export function ProductLogo(props: { className: string }) {
    const { isDark } = useContext(ThemeContext);
    const isDataOps = useIsDataOps();
    return (
        <>
            {isDataOps ? (
                <img src={isDark ? dataOpsIconDark : dataOpsIconLight} alt="Dataops logo" className={props.className} />
            ) : (
                <img src={devtoolIcon} alt="Devtool's logo" className={props.className} />
            )}
        </>
    );
}
