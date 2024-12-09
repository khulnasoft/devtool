-- Copyright (c) 2020 Devtool GmbH. All rights reserved.
-- Licensed under the GNU Affero General Public License (AGPL). See License.AGPL.txt in the project root for license information.

-- must be idempotent
CREATE DATABASE IF NOT EXISTS `authorization` CHARSET utf8mb4;

-- Grant privileges
GRANT ALL ON `authorization`.* TO "__DEVTOOL_USERNAME__"@"%";
