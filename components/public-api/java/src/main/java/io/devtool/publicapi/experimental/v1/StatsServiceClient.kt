// Copyright (c) 2024 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by connect-kotlin. DO NOT EDIT.
//
// Source: devtool/experimental/v1/stats.proto
//
package io.devtool.publicapi.experimental.v1

import com.connectrpc.Headers
import com.connectrpc.MethodSpec
import com.connectrpc.ProtocolClientInterface
import com.connectrpc.ResponseMessage
import com.connectrpc.StreamType

public class StatsServiceClient(
  private val client: ProtocolClientInterface,
) : StatsServiceClientInterface {
  /**
   *  Retrieves the current user stats
   */
  override suspend fun getUserStats(request: Stats.GetUserStatsRequest, headers: Headers):
      ResponseMessage<Stats.GetUserStatsResponse> = client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.experimental.v1.StatsService/GetUserStats",
      io.devtool.publicapi.experimental.v1.Stats.GetUserStatsRequest::class,
      io.devtool.publicapi.experimental.v1.Stats.GetUserStatsResponse::class,
      StreamType.UNARY,
    ),
  )

}
