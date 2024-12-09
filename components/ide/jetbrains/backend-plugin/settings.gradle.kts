// Copyright (c) 2021 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

rootProject.name = "devtool-remote"

include(":supervisor-api")
val supervisorApiProjectPath: String by settings
project(":supervisor-api").projectDir = File(supervisorApiProjectPath)

include(":devtool-protocol")
val devtoolProtocolProjectPath: String by settings
project(":devtool-protocol").projectDir = File(devtoolProtocolProjectPath)
