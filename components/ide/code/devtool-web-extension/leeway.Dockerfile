# Copyright (c) 2020 Devtool GmbH. All rights reserved.
# Licensed under the GNU Affero General Public License (AGPL).
# See License.AGPL.txt in the project root for license information.
FROM node:18 as builder

ARG CODE_EXTENSION_COMMIT

RUN apt update -y \
    && apt install jq --no-install-recommends -y

RUN mkdir /devtool-code-web \
    && cd /devtool-code-web \
    && git init \
    && git remote add origin https://github.com/khulnasoft/devtool-code \
    && git fetch origin $CODE_EXTENSION_COMMIT --depth=1 \
    && git reset --hard FETCH_HEAD
WORKDIR /devtool-code-web
RUN yarn --frozen-lockfile --network-timeout 180000

# update package.json
RUN cd devtool-web && \
    setSegmentKey="setpath([\"segmentKey\"]; \"untrusted-dummy-key\")" && \
    jqCommands="${setSegmentKey}" && \
    cat package.json | jq "${jqCommands}" > package.json.tmp && \
    mv package.json.tmp package.json
RUN yarn build:devtool-web && yarn --cwd devtool-web/ inject-commit-hash


FROM scratch

COPY --from=builder --chown=33333:33333 /devtool-code-web/devtool-web/out /ide/extensions/devtool-web/out/
COPY --from=builder --chown=33333:33333 /devtool-code-web/devtool-web/public /ide/extensions/devtool-web/public/
COPY --from=builder --chown=33333:33333 /devtool-code-web/devtool-web/resources /ide/extensions/devtool-web/resources/
COPY --from=builder --chown=33333:33333 /devtool-code-web/devtool-web/package.json /devtool-code-web/devtool-web/package.nls.json /devtool-code-web/devtool-web/README.md /devtool-code-web/devtool-web/LICENSE.txt /ide/extensions/devtool-web/
