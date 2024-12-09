-- Copyright (c) 2020 Devtool GmbH. All rights reserved.
-- Licensed under the GNU Affero General Public License (AGPL). See License.AGPL.txt in the project root for license information.

-- must be idempotent

-- @devtoolDB contains name of the DB the script manipulates, and is replaced by the file reader
SET
@devtoolDB = IFNULL(@devtoolDB, '`__DEVTOOL_DB_NAME__`');

SET
@statementStr = CONCAT('DROP DATABASE IF EXISTS ', @devtoolDB);
PREPARE statement FROM @statementStr;
EXECUTE statement;

SET
@statementStr = CONCAT('CREATE DATABASE ', @devtoolDB, ' CHARSET utf8mb4');
PREPARE statement FROM @statementStr;
EXECUTE statement;
