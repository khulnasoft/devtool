// Copyright (c) 2022 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package io.devtool.devtoolprotocol.api.entities;

public class SetWorkspaceTimeoutResult {
    private String[] resetTimeoutOnWorkspaces;
    private String humanReadableDuration;

    public SetWorkspaceTimeoutResult(String[] resetTimeoutOnWorkspaces, String humanReadableDuration) {
        this.resetTimeoutOnWorkspaces = resetTimeoutOnWorkspaces;
        this.humanReadableDuration = humanReadableDuration;
    }

    public String[] getResetTimeoutOnWorkspaces() {
        return resetTimeoutOnWorkspaces;
    }

    public String getHumanReadableDuration() {
        return humanReadableDuration;
    }
}
