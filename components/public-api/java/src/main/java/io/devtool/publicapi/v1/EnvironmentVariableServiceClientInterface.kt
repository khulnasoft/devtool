// Copyright (c) 2024 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by connect-kotlin. DO NOT EDIT.
//
// Source: devtool/v1/envvar.proto
//
package io.devtool.publicapi.v1

import com.connectrpc.Headers
import com.connectrpc.ResponseMessage

public interface EnvironmentVariableServiceClientInterface {
  /**
   *  ListUserEnvironmentVariables returns all environment variables for the
   *  authenticated user.
   */
  public suspend
      fun listUserEnvironmentVariables(request: Envvar.ListUserEnvironmentVariablesRequest,
      headers: Headers = emptyMap()): ResponseMessage<Envvar.ListUserEnvironmentVariablesResponse>

  /**
   *  UpdateUserEnvironmentVariable updates an environment variable for the
   *  authenticated user.
   */
  public suspend
      fun updateUserEnvironmentVariable(request: Envvar.UpdateUserEnvironmentVariableRequest,
      headers: Headers = emptyMap()): ResponseMessage<Envvar.UpdateUserEnvironmentVariableResponse>

  /**
   *  CreateUserEnvironmentVariable creates a new environment variable for the
   *  authenticated user.
   */
  public suspend
      fun createUserEnvironmentVariable(request: Envvar.CreateUserEnvironmentVariableRequest,
      headers: Headers = emptyMap()): ResponseMessage<Envvar.CreateUserEnvironmentVariableResponse>

  /**
   *  DeleteUserEnvironmentVariable deletes an environment variable for the
   *  authenticated user.
   */
  public suspend
      fun deleteUserEnvironmentVariable(request: Envvar.DeleteUserEnvironmentVariableRequest,
      headers: Headers = emptyMap()): ResponseMessage<Envvar.DeleteUserEnvironmentVariableResponse>

  /**
   *  ListConfigurationEnvironmentVariables returns all environment variables in
   *  a configuration.
   */
  public suspend
      fun listConfigurationEnvironmentVariables(request: Envvar.ListConfigurationEnvironmentVariablesRequest,
      headers: Headers = emptyMap()):
      ResponseMessage<Envvar.ListConfigurationEnvironmentVariablesResponse>

  /**
   *  UpdateConfigurationEnvironmentVariable updates an environment variable in
   *  a configuration.
   */
  public suspend
      fun updateConfigurationEnvironmentVariable(request: Envvar.UpdateConfigurationEnvironmentVariableRequest,
      headers: Headers = emptyMap()):
      ResponseMessage<Envvar.UpdateConfigurationEnvironmentVariableResponse>

  /**
   *  CreateConfigurationEnvironmentVariable creates a new environment variable
   *  in a configuration.
   */
  public suspend
      fun createConfigurationEnvironmentVariable(request: Envvar.CreateConfigurationEnvironmentVariableRequest,
      headers: Headers = emptyMap()):
      ResponseMessage<Envvar.CreateConfigurationEnvironmentVariableResponse>

  /**
   *  DeleteConfigurationEnvironmentVariable deletes an environment variable in
   *  a configuration.
   */
  public suspend
      fun deleteConfigurationEnvironmentVariable(request: Envvar.DeleteConfigurationEnvironmentVariableRequest,
      headers: Headers = emptyMap()):
      ResponseMessage<Envvar.DeleteConfigurationEnvironmentVariableResponse>

  public suspend
      fun resolveWorkspaceEnvironmentVariables(request: Envvar.ResolveWorkspaceEnvironmentVariablesRequest,
      headers: Headers = emptyMap()):
      ResponseMessage<Envvar.ResolveWorkspaceEnvironmentVariablesResponse>
}