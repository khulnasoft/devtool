#!/bin/bash
# Copyright (c) 2022 Devtool GmbH. All rights reserved.
# Licensed under the GNU Affero General Public License (AGPL).
# See License.AGPL.txt in the project root for license information.

set -Eeuo pipefail
source /workspace/devtool/scripts/ws-deploy.sh deployment ws-proxy false
