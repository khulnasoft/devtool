// Copyright (c) 2024 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by connect-kotlin. DO NOT EDIT.
//
// Source: devtool/v1/verification.proto
//
package io.devtool.publicapi.v1

import com.connectrpc.Headers
import com.connectrpc.MethodSpec
import com.connectrpc.ProtocolClientInterface
import com.connectrpc.ResponseMessage
import com.connectrpc.StreamType

public class VerificationServiceClient(
  private val client: ProtocolClientInterface,
) : VerificationServiceClientInterface {
  /**
   *  SendPhoneNumberVerificationToken sends a verification token to the
   *  specified phone number.
   */
  override suspend
      fun sendPhoneNumberVerificationToken(request: Verification.SendPhoneNumberVerificationTokenRequest,
      headers: Headers): ResponseMessage<Verification.SendPhoneNumberVerificationTokenResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.v1.VerificationService/SendPhoneNumberVerificationToken",
      io.devtool.publicapi.v1.Verification.SendPhoneNumberVerificationTokenRequest::class,
      io.devtool.publicapi.v1.Verification.SendPhoneNumberVerificationTokenResponse::class,
      StreamType.UNARY,
    ),
  )


  /**
   *  VerifyPhoneNumberVerificationToken verifies the specified verification
   *  token.
   */
  override suspend
      fun verifyPhoneNumberVerificationToken(request: Verification.VerifyPhoneNumberVerificationTokenRequest,
      headers: Headers): ResponseMessage<Verification.VerifyPhoneNumberVerificationTokenResponse> =
      client.unary(
    request,
    headers,
    MethodSpec(
    "devtool.v1.VerificationService/VerifyPhoneNumberVerificationToken",
      io.devtool.publicapi.v1.Verification.VerifyPhoneNumberVerificationTokenRequest::class,
      io.devtool.publicapi.v1.Verification.VerifyPhoneNumberVerificationTokenResponse::class,
      StreamType.UNARY,
    ),
  )

}
