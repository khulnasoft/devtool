/**
 * Copyright (c) 2024 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

// @generated by protoc-gen-es v1.3.3 with parameter "target=ts"
// @generated from file devtool/v1/token.proto (package devtool.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message devtool.v1.CreateTemporaryAccessTokenRequest
 */
export class CreateTemporaryAccessTokenRequest extends Message<CreateTemporaryAccessTokenRequest> {
  /**
   * user_id is the identifier of the user for which the token is created.
   *
   * @generated from field: string user_id = 1;
   */
  userId = "";

  /**
   * expiry_seconds is the number of seconds the token is valid for.
   * value should in the range [1, 600]
   *
   * @generated from field: int32 expiry_seconds = 2;
   */
  expirySeconds = 0;

  constructor(data?: PartialMessage<CreateTemporaryAccessTokenRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devtool.v1.CreateTemporaryAccessTokenRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "expiry_seconds", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateTemporaryAccessTokenRequest {
    return new CreateTemporaryAccessTokenRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateTemporaryAccessTokenRequest {
    return new CreateTemporaryAccessTokenRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateTemporaryAccessTokenRequest {
    return new CreateTemporaryAccessTokenRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateTemporaryAccessTokenRequest | PlainMessage<CreateTemporaryAccessTokenRequest> | undefined, b: CreateTemporaryAccessTokenRequest | PlainMessage<CreateTemporaryAccessTokenRequest> | undefined): boolean {
    return proto3.util.equals(CreateTemporaryAccessTokenRequest, a, b);
  }
}

/**
 * @generated from message devtool.v1.CreateTemporaryAccessTokenResponse
 */
export class CreateTemporaryAccessTokenResponse extends Message<CreateTemporaryAccessTokenResponse> {
  /**
   * cookie_name is the name of the cookie to use for the token.
   *
   * @generated from field: string cookie_name = 1;
   */
  cookieName = "";

  /**
   * @generated from field: string token = 2;
   */
  token = "";

  constructor(data?: PartialMessage<CreateTemporaryAccessTokenResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "devtool.v1.CreateTemporaryAccessTokenResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "cookie_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateTemporaryAccessTokenResponse {
    return new CreateTemporaryAccessTokenResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateTemporaryAccessTokenResponse {
    return new CreateTemporaryAccessTokenResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateTemporaryAccessTokenResponse {
    return new CreateTemporaryAccessTokenResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateTemporaryAccessTokenResponse | PlainMessage<CreateTemporaryAccessTokenResponse> | undefined, b: CreateTemporaryAccessTokenResponse | PlainMessage<CreateTemporaryAccessTokenResponse> | undefined): boolean {
    return proto3.util.equals(CreateTemporaryAccessTokenResponse, a, b);
  }
}
