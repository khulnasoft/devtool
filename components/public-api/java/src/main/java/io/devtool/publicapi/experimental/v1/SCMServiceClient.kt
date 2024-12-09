// Copyright (c) 2024 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by connect-kotlin. DO NOT EDIT.
//
// Source: devtool/experimental/v1/scm.proto
//
package io.devtool.publicapi.experimental.v1

import com.connectrpc.Headers
import com.connectrpc.MethodSpec
import com.connectrpc.ProtocolClientInterface
import com.connectrpc.ResponseMessage
import com.connectrpc.StreamType

public class SCMServiceClient(
  private val client: ProtocolClientInterface,
) : SCMServiceClientInterface {
  /**
   *  GetSuggestedRepoURLs returns a list of suggested repositories to open for
   *  the user.
   */
  override suspend fun getSuggestedRepoURLs(request: Scm.GetSuggestedRepoURLsRequest,
      headers: Headers): ResponseMessage<Scm.GetSuggestedRepoURLsResponse> = client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.experimental.v1.SCMService/GetSuggestedRepoURLs",
      io.devtool.publicapi.experimental.v1.Scm.GetSuggestedRepoURLsRequest::class,
      io.devtool.publicapi.experimental.v1.Scm.GetSuggestedRepoURLsResponse::class,
      StreamType.UNARY,
    ),
  )

}
