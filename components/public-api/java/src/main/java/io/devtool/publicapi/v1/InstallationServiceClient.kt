// Copyright (c) 2024 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by connect-kotlin. DO NOT EDIT.
//
// Source: devtool/v1/installation.proto
//
package io.devtool.publicapi.v1

import com.connectrpc.Headers
import com.connectrpc.MethodSpec
import com.connectrpc.ProtocolClientInterface
import com.connectrpc.ResponseMessage
import com.connectrpc.StreamType

public class InstallationServiceClient(
  private val client: ProtocolClientInterface,
) : InstallationServiceClientInterface {
  /**
   *  GetInstallationWorkspaceDefaultImage returns the default image for current
   *  Devtool Installation.
   */
  override suspend
      fun getInstallationWorkspaceDefaultImage(request: Installation.GetInstallationWorkspaceDefaultImageRequest,
      headers: Headers): ResponseMessage<Installation.GetInstallationWorkspaceDefaultImageResponse>
      = client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.v1.InstallationService/GetInstallationWorkspaceDefaultImage",
      io.devtool.publicapi.v1.Installation.GetInstallationWorkspaceDefaultImageRequest::class,
      io.devtool.publicapi.v1.Installation.GetInstallationWorkspaceDefaultImageResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  ListBlockedRepositories lists blocked repositories.
   */
  override suspend fun listBlockedRepositories(request: Installation.ListBlockedRepositoriesRequest,
      headers: Headers): ResponseMessage<Installation.ListBlockedRepositoriesResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.v1.InstallationService/ListBlockedRepositories",
      io.devtool.publicapi.v1.Installation.ListBlockedRepositoriesRequest::class,
      io.devtool.publicapi.v1.Installation.ListBlockedRepositoriesResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  CreateBlockedRepository creates a new blocked repository.
   */
  override suspend fun createBlockedRepository(request: Installation.CreateBlockedRepositoryRequest,
      headers: Headers): ResponseMessage<Installation.CreateBlockedRepositoryResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.v1.InstallationService/CreateBlockedRepository",
      io.devtool.publicapi.v1.Installation.CreateBlockedRepositoryRequest::class,
      io.devtool.publicapi.v1.Installation.CreateBlockedRepositoryResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  DeleteBlockedRepository deletes a blocked repository.
   */
  override suspend fun deleteBlockedRepository(request: Installation.DeleteBlockedRepositoryRequest,
      headers: Headers): ResponseMessage<Installation.DeleteBlockedRepositoryResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.v1.InstallationService/DeleteBlockedRepository",
      io.devtool.publicapi.v1.Installation.DeleteBlockedRepositoryRequest::class,
      io.devtool.publicapi.v1.Installation.DeleteBlockedRepositoryResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  ListBlockedEmailDomains lists blocked email domains.
   */
  override suspend fun listBlockedEmailDomains(request: Installation.ListBlockedEmailDomainsRequest,
      headers: Headers): ResponseMessage<Installation.ListBlockedEmailDomainsResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.v1.InstallationService/ListBlockedEmailDomains",
      io.devtool.publicapi.v1.Installation.ListBlockedEmailDomainsRequest::class,
      io.devtool.publicapi.v1.Installation.ListBlockedEmailDomainsResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  CreateBlockedEmailDomain creates a new blocked email domain.
   */
  override suspend
      fun createBlockedEmailDomain(request: Installation.CreateBlockedEmailDomainRequest,
      headers: Headers): ResponseMessage<Installation.CreateBlockedEmailDomainResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.v1.InstallationService/CreateBlockedEmailDomain",
      io.devtool.publicapi.v1.Installation.CreateBlockedEmailDomainRequest::class,
      io.devtool.publicapi.v1.Installation.CreateBlockedEmailDomainResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  GetOnboardingState returns the onboarding state of the installation.
   */
  override suspend fun getOnboardingState(request: Installation.GetOnboardingStateRequest,
      headers: Headers): ResponseMessage<Installation.GetOnboardingStateResponse> = client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.v1.InstallationService/GetOnboardingState",
      io.devtool.publicapi.v1.Installation.GetOnboardingStateRequest::class,
      io.devtool.publicapi.v1.Installation.GetOnboardingStateResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  GetInstallationConfiguration returns configuration of the installation.
   */
  override suspend
      fun getInstallationConfiguration(request: Installation.GetInstallationConfigurationRequest,
      headers: Headers): ResponseMessage<Installation.GetInstallationConfigurationResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.v1.InstallationService/GetInstallationConfiguration",
      io.devtool.publicapi.v1.Installation.GetInstallationConfigurationRequest::class,
      io.devtool.publicapi.v1.Installation.GetInstallationConfigurationResponse::class,
      StreamType.UNARY,
    ),
  )

}