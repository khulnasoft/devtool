# Copyright (c) 2020 Devtool GmbH. All rights reserved.
# Licensed under the GNU Affero General Public License (AGPL).
# See License.AGPL.txt in the project root for license information.

FROM cgr.dev/chainguard/wolfi-base:latest@sha256:b3dd9cf08283b959c6a0a3c833e68b2882a50129930215060154b43ae6a3e81c

COPY components-node-labeler--app/node-labeler /app/node-labeler

ARG __GIT_COMMIT
ARG VERSION

ENV DEVTOOL_BUILD_GIT_COMMIT=${__GIT_COMMIT}
ENV DEVTOOL_BUILD_VERSION=${VERSION}
ENTRYPOINT [ "/app/node-labeler" ]
CMD [ "-v", "help" ]
