#!/bin/bash
# Copyright (c) 2021 Devtool GmbH. All rights reserved.
# Licensed under the GNU Affero General Public License (AGPL).
# See License.AGPL.txt in the project root for license information.

# This scipt connects via Google's cloud_sql_proxy to a database and runs the db-migrations

# ENV variables for configuration:
# * GOOGLE_APPLICATION_CREDENTIALS_DATA: contents of the crendetials files that cloud_sql_proxy uses for authentication
# * GCP_DATABASE: database name
# * DB_PASSWORD: database password

# Example usage:
# docker run --rm \
#            --env GOOGLE_APPLICATION_CREDENTIALS_DATA='...' \
#            --env GCP_DATABASE="devtool-foobar:europe-west1:devtool-foobar-baz" \
#            --env DB_PASSWORD="..." \
#            gcr.io/devtool-core-dev/build/db-migrations:x1 /app/migrate_gcp.sh

set -euo pipefail

/app/typeorm_gcp.sh migrations:run
