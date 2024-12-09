# Copyright (c) 2021 Devtool GmbH. All rights reserved.
# Licensed under the GNU Affero General Public License (AGPL).
# See License.AGPL.txt in the project root for license information.

FROM caddy:2.7.6-builder AS builder

WORKDIR /plugins

COPY plugins /plugins

# build caddy
RUN xcaddy build v2.7.6 \
  --output /caddy \
  --with github.com/khulnasoft/devtool/proxy/plugins/corsorigin=/plugins/corsorigin \
  --with github.com/khulnasoft/devtool/proxy/plugins/secwebsocketkey=/plugins/secwebsocketkey \
  --with github.com/khulnasoft/devtool/proxy/plugins/workspacedownload=/plugins/workspacedownload \
  --with github.com/khulnasoft/devtool/proxy/plugins/headlesslogdownload=/plugins/headlesslogdownload \
  --with github.com/khulnasoft/devtool/proxy/plugins/configcat=/plugins/configcat \
  --with github.com/khulnasoft/devtool/proxy/plugins/analytics=/plugins/analytics \
  --with github.com/khulnasoft/devtool/proxy/plugins/logif=/plugins/logif \
  --with github.com/khulnasoft/devtool/proxy/plugins/jsonselect=/plugins/jsonselect \
  --with github.com/khulnasoft/devtool/proxy/plugins/sshtunnel=/plugins/sshtunnel \
  --with github.com/khulnasoft/devtool/proxy/plugins/frontend_dev=/plugins/frontend_dev

FROM cgr.dev/chainguard/wolfi-base:latest@sha256:b3dd9cf08283b959c6a0a3c833e68b2882a50129930215060154b43ae6a3e81c

# Ensure latest packages are present, like security updates.
RUN  apk upgrade --no-cache \
  && apk add --no-cache ca-certificates bash

# Debug convenience
ENV TERM=xterm
ENV SHELL=/bin/bash

COPY --from=builder /caddy /usr/bin/caddy

COPY conf/Caddyfile /etc/caddy/Caddyfile
COPY conf/vhost.empty /etc/caddy/vhosts/vhost.empty
COPY conf/workspace-handler.* /etc/caddy/workspace-handler/

ARG __GIT_COMMIT
ARG VERSION

ENV DEVTOOL_BUILD_GIT_COMMIT=${__GIT_COMMIT}
ENV DEVTOOL_BUILD_VERSION=${VERSION}

CMD [ "caddy", "run", "--adapter", "inject-ssh-tunnel" , "--watch", "--config", "/etc/caddy/Caddyfile" ]
