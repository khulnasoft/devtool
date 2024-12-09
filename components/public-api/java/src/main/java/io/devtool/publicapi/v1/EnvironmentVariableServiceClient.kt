// Copyright (c) 2024 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by connect-kotlin. DO NOT EDIT.
//
// Source: devtool/v1/envvar.proto
//
package io.devtool.publicapi.v1

import com.connectrpc.Headers
import com.connectrpc.MethodSpec
import com.connectrpc.ProtocolClientInterface
import com.connectrpc.ResponseMessage
import com.connectrpc.StreamType

public class EnvironmentVariableServiceClient(
  private val client: ProtocolClientInterface,
) : EnvironmentVariableServiceClientInterface {
  /**
   *  ListUserEnvironmentVariables returns all environment variables for the
   *  authenticated user.
   */
  override suspend
      fun listUserEnvironmentVariables(request: Envvar.ListUserEnvironmentVariablesRequest,
      headers: Headers): ResponseMessage<Envvar.ListUserEnvironmentVariablesResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.v1.EnvironmentVariableService/ListUserEnvironmentVariables",
      io.devtool.publicapi.v1.Envvar.ListUserEnvironmentVariablesRequest::class,
      io.devtool.publicapi.v1.Envvar.ListUserEnvironmentVariablesResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  UpdateUserEnvironmentVariable updates an environment variable for the
   *  authenticated user.
   */
  override suspend
      fun updateUserEnvironmentVariable(request: Envvar.UpdateUserEnvironmentVariableRequest,
      headers: Headers): ResponseMessage<Envvar.UpdateUserEnvironmentVariableResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.v1.EnvironmentVariableService/UpdateUserEnvironmentVariable",
      io.devtool.publicapi.v1.Envvar.UpdateUserEnvironmentVariableRequest::class,
      io.devtool.publicapi.v1.Envvar.UpdateUserEnvironmentVariableResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  CreateUserEnvironmentVariable creates a new environment variable for the
   *  authenticated user.
   */
  override suspend
      fun createUserEnvironmentVariable(request: Envvar.CreateUserEnvironmentVariableRequest,
      headers: Headers): ResponseMessage<Envvar.CreateUserEnvironmentVariableResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.v1.EnvironmentVariableService/CreateUserEnvironmentVariable",
      io.devtool.publicapi.v1.Envvar.CreateUserEnvironmentVariableRequest::class,
      io.devtool.publicapi.v1.Envvar.CreateUserEnvironmentVariableResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  DeleteUserEnvironmentVariable deletes an environment variable for the
   *  authenticated user.
   */
  override suspend
      fun deleteUserEnvironmentVariable(request: Envvar.DeleteUserEnvironmentVariableRequest,
      headers: Headers): ResponseMessage<Envvar.DeleteUserEnvironmentVariableResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.v1.EnvironmentVariableService/DeleteUserEnvironmentVariable",
      io.devtool.publicapi.v1.Envvar.DeleteUserEnvironmentVariableRequest::class,
      io.devtool.publicapi.v1.Envvar.DeleteUserEnvironmentVariableResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  ListConfigurationEnvironmentVariables returns all environment variables in
   *  a configuration.
   */
  override suspend
      fun listConfigurationEnvironmentVariables(request: Envvar.ListConfigurationEnvironmentVariablesRequest,
      headers: Headers): ResponseMessage<Envvar.ListConfigurationEnvironmentVariablesResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.v1.EnvironmentVariableService/ListConfigurationEnvironmentVariables",
      io.devtool.publicapi.v1.Envvar.ListConfigurationEnvironmentVariablesRequest::class,
      io.devtool.publicapi.v1.Envvar.ListConfigurationEnvironmentVariablesResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  UpdateConfigurationEnvironmentVariable updates an environment variable in
   *  a configuration.
   */
  override suspend
      fun updateConfigurationEnvironmentVariable(request: Envvar.UpdateConfigurationEnvironmentVariableRequest,
      headers: Headers): ResponseMessage<Envvar.UpdateConfigurationEnvironmentVariableResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.v1.EnvironmentVariableService/UpdateConfigurationEnvironmentVariable",
      io.devtool.publicapi.v1.Envvar.UpdateConfigurationEnvironmentVariableRequest::class,
      io.devtool.publicapi.v1.Envvar.UpdateConfigurationEnvironmentVariableResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  CreateConfigurationEnvironmentVariable creates a new environment variable
   *  in a configuration.
   */
  override suspend
      fun createConfigurationEnvironmentVariable(request: Envvar.CreateConfigurationEnvironmentVariableRequest,
      headers: Headers): ResponseMessage<Envvar.CreateConfigurationEnvironmentVariableResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.v1.EnvironmentVariableService/CreateConfigurationEnvironmentVariable",
      io.devtool.publicapi.v1.Envvar.CreateConfigurationEnvironmentVariableRequest::class,
      io.devtool.publicapi.v1.Envvar.CreateConfigurationEnvironmentVariableResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  DeleteConfigurationEnvironmentVariable deletes an environment variable in
   *  a configuration.
   */
  override suspend
      fun deleteConfigurationEnvironmentVariable(request: Envvar.DeleteConfigurationEnvironmentVariableRequest,
      headers: Headers): ResponseMessage<Envvar.DeleteConfigurationEnvironmentVariableResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.v1.EnvironmentVariableService/DeleteConfigurationEnvironmentVariable",
      io.devtool.publicapi.v1.Envvar.DeleteConfigurationEnvironmentVariableRequest::class,
      io.devtool.publicapi.v1.Envvar.DeleteConfigurationEnvironmentVariableResponse::class,
      StreamType.UNARY,
    ),
  )


  override suspend
      fun resolveWorkspaceEnvironmentVariables(request: Envvar.ResolveWorkspaceEnvironmentVariablesRequest,
      headers: Headers): ResponseMessage<Envvar.ResolveWorkspaceEnvironmentVariablesResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.v1.EnvironmentVariableService/ResolveWorkspaceEnvironmentVariables",
      io.devtool.publicapi.v1.Envvar.ResolveWorkspaceEnvironmentVariablesRequest::class,
      io.devtool.publicapi.v1.Envvar.ResolveWorkspaceEnvironmentVariablesResponse::class,
      StreamType.UNARY,
    ),
  )

}
