/**
 * Copyright (c) 2024 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

// @generated by protoc-gen-connect-es v1.1.2 with parameter "target=ts"
// @generated from file devtool/v1/ssh.proto (package devtool.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateSSHPublicKeyRequest, CreateSSHPublicKeyResponse, DeleteSSHPublicKeyRequest, DeleteSSHPublicKeyResponse, ListSSHPublicKeysRequest, ListSSHPublicKeysResponse } from "./ssh_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service devtool.v1.SSHService
 */
export const SSHService = {
  typeName: "devtool.v1.SSHService",
  methods: {
    /**
     * ListSSHPublicKeys returns all the ssh public keys for the
     * authenticated user.
     *
     * @generated from rpc devtool.v1.SSHService.ListSSHPublicKeys
     */
    listSSHPublicKeys: {
      name: "ListSSHPublicKeys",
      I: ListSSHPublicKeysRequest,
      O: ListSSHPublicKeysResponse,
      kind: MethodKind.Unary,
    },
    /**
     * CreateSSHPublicKeys creates an ssh public key for the
     * authenticated user.
     *
     * @generated from rpc devtool.v1.SSHService.CreateSSHPublicKey
     */
    createSSHPublicKey: {
      name: "CreateSSHPublicKey",
      I: CreateSSHPublicKeyRequest,
      O: CreateSSHPublicKeyResponse,
      kind: MethodKind.Unary,
    },
    /**
     * DeleteSSHPublicKeys deletes an ssh public key for the
     * authenticated user.
     *
     * @generated from rpc devtool.v1.SSHService.DeleteSSHPublicKey
     */
    deleteSSHPublicKey: {
      name: "DeleteSSHPublicKey",
      I: DeleteSSHPublicKeyRequest,
      O: DeleteSSHPublicKeyResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;