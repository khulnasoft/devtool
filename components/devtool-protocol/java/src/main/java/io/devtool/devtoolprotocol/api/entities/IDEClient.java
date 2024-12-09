// Copyright (c) 2022 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.devtool.devtoolprotocol.api.entities;

import java.util.List;

public class IDEClient {

    private String defaultDesktopIDE;

    private List<String> desktopIDEs;

    public String getDefaultDesktopIDE() {
        return defaultDesktopIDE;
    }

    public void setDefaultDesktopIDE(String defaultDesktopIDE) {
        this.defaultDesktopIDE = defaultDesktopIDE;
    }

    public List<String> getDesktopIDEs() {
        return desktopIDEs;
    }

    public void setDesktopIDEs(List<String> desktopIDEs) {
        this.desktopIDEs = desktopIDEs;
    }
}
