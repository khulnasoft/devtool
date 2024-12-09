// Copyright (c) 2024 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by connect-kotlin. DO NOT EDIT.
//
// Source: devtool/experimental/v1/ide_client.proto
//
package io.devtool.publicapi.experimental.v1

import com.connectrpc.Headers
import com.connectrpc.MethodSpec
import com.connectrpc.ProtocolClientInterface
import com.connectrpc.ResponseMessage
import com.connectrpc.StreamType

public class IDEClientServiceClient(
  private val client: ProtocolClientInterface,
) : IDEClientServiceClientInterface {
  /**
   *  SendHeartbeat sends a clientheartbeat signal for a running workspace.
   */
  override suspend fun sendHeartbeat(request: IdeClient.SendHeartbeatRequest, headers: Headers):
      ResponseMessage<IdeClient.SendHeartbeatResponse> = client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.experimental.v1.IDEClientService/SendHeartbeat",
      io.devtool.publicapi.experimental.v1.IdeClient.SendHeartbeatRequest::class,
      io.devtool.publicapi.experimental.v1.IdeClient.SendHeartbeatResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  SendDidClose sends a client close signal for a running workspace.
   */
  override suspend fun sendDidClose(request: IdeClient.SendDidCloseRequest, headers: Headers):
      ResponseMessage<IdeClient.SendDidCloseResponse> = client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.experimental.v1.IDEClientService/SendDidClose",
      io.devtool.publicapi.experimental.v1.IdeClient.SendDidCloseRequest::class,
      io.devtool.publicapi.experimental.v1.IdeClient.SendDidCloseResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  UpdateGitStatus updates the status of a repository in a workspace.
   */
  override suspend fun updateGitStatus(request: IdeClient.UpdateGitStatusRequest, headers: Headers):
      ResponseMessage<IdeClient.UpdateGitStatusResponse> = client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.experimental.v1.IDEClientService/UpdateGitStatus",
      io.devtool.publicapi.experimental.v1.IdeClient.UpdateGitStatusRequest::class,
      io.devtool.publicapi.experimental.v1.IdeClient.UpdateGitStatusResponse::class,
      StreamType.UNARY,
    ),
  )

}
