/**
 * Copyright (c) 2024 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

// @generated by protoc-gen-connect-es v1.1.2 with parameter "target=ts"
// @generated from file devtool/experimental/v1/stats.proto (package devtool.experimental.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetUserStatsRequest, GetUserStatsResponse } from "./stats_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service devtool.experimental.v1.StatsService
 */
export const StatsService = {
  typeName: "devtool.experimental.v1.StatsService",
  methods: {
    /**
     * Retrieves the current user stats
     *
     * @generated from rpc devtool.experimental.v1.StatsService.GetUserStats
     */
    getUserStats: {
      name: "GetUserStats",
      I: GetUserStatsRequest,
      O: GetUserStatsResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
