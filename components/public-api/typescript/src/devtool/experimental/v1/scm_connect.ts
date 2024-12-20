/**
 * Copyright (c) 2024 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

// @generated by protoc-gen-connect-es v1.1.2 with parameter "target=ts"
// @generated from file devtool/experimental/v1/scm.proto (package devtool.experimental.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetSuggestedRepoURLsRequest, GetSuggestedRepoURLsResponse } from "./scm_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service devtool.experimental.v1.SCMService
 */
export const SCMService = {
  typeName: "devtool.experimental.v1.SCMService",
  methods: {
    /**
     * GetSuggestedRepoURLs returns a list of suggested repositories to open for
     * the user.
     *
     * @generated from rpc devtool.experimental.v1.SCMService.GetSuggestedRepoURLs
     */
    getSuggestedRepoURLs: {
      name: "GetSuggestedRepoURLs",
      I: GetSuggestedRepoURLsRequest,
      O: GetSuggestedRepoURLsResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
