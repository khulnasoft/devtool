#!/bin/sh
# Copyright (c) 2022 Devtool GmbH. All rights reserved.
# Licensed under the GNU Affero General Public License (AGPL).
# See License.AGPL.txt in the project root for license information.

set -e

touch logs.txt

# Set Domain to `preview.devtool-self-hosted.com` if not set
if [ -z "${DOMAIN}" ]; then
  export DOMAIN="preview.devtool-self-hosted.com"
fi

# Create a USER_ID to be used everywhere
USER_ID="$(od -x /dev/urandom | head -1 | awk '{OFS="-"; print $2$3,$4,$5,$6,$7$8$9}')"
export USER_ID

if [ "$1" != "logging" ]; then
  $0 logging > logs.txt 2>&1 &
  /prettylog
  exit
fi

# check for minimum requirements
REQUIRED_MEM_KB=$((6 * 1024 * 1024))
total_mem_kb=$(awk '/MemTotal:/ {print $2}' /proc/meminfo)
if [ "${total_mem_kb}" -lt "${REQUIRED_MEM_KB}" ]; then
    echo "Devtool local preview requires a system with at least 6GB of memory"
    exit 1
fi

REQUIRED_CORES=4
total_cores=$(nproc)
if [ "${total_cores}" -lt "${REQUIRED_CORES}" ]; then
    echo "Devtool local preview requires a system with at least 4 CPU Cores"
    exit 1
fi

echo "Devtool Domain: $DOMAIN"

# With cgroupv2, We need to move the k3s processes into the
# init group when we override the entrypoint in the container
if [ -f /sys/fs/cgroup/cgroup.controllers ]; then
  mkdir -p /sys/fs/cgroup/init
  busybox xargs -rn1 < /sys/fs/cgroup/cgroup.procs > /sys/fs/cgroup/init/cgroup.procs || :
  sed -e 's/ / +/g' -e 's/^/+/' <"/sys/fs/cgroup/cgroup.controllers" >"/sys/fs/cgroup/cgroup.subtree_control"
fi


mount --make-shared /sys/fs/cgroup
mount --make-shared /proc
mount --make-shared /var/devtool

# install in local store
mkcert -install
cat "${HOME}"/.local/share/mkcert/rootCA.pem >> /etc/ssl/certs/ca-certificates.crt
# also send root cert into a volume
cat "${HOME}"/.local/share/mkcert/rootCA.pem > /var/devtool/devtool-ca.crt

FN_CACERT="./ca.pem"
FN_SSLCERT="./ssl.crt"
FN_SSLKEY="./ssl.key"

cat "${HOME}"/.local/share/mkcert/rootCA.pem > "$FN_CACERT"
mkcert -cert-file "$FN_SSLCERT" \
  -key-file "$FN_SSLKEY" \
  "*.ws.${DOMAIN}" "*.${DOMAIN}" "${DOMAIN}" "reg.${DOMAIN}" "registry.default.svc.cluster.local" "devtool.default" "ws-manager.default.svc" "ws-manager" "ws-manager-dev" "registry-facade" "server" "ws-manager-bridge" "ws-proxy" "ws-manager" "ws-daemon.default.svc" "ws-daemon" "wsdaemon"

CACERT=$(base64 -w0 < "$FN_CACERT")
SSLCERT=$(base64 -w0 < "$FN_SSLCERT")
SSLKEY=$(base64 -w0 < "$FN_SSLKEY")

mkdir -p /var/lib/rancher/k3s/server/manifests/devtool

cat << EOF > /var/lib/rancher/k3s/server/manifests/devtool/customCA-cert.yaml
---
apiVersion: v1
kind: Secret
metadata:
  name: ca-key-pair
  labels:
    app: devtool
data:
  ca.crt: $CACERT
EOF

cat << EOF > /var/lib/rancher/k3s/server/manifests/devtool/https-cert.yaml
---
apiVersion: v1
kind: Secret
metadata:
  name: https-certificates
  labels:
    app: devtool
type: kubernetes.io/tls
data:
  tls.crt: $SSLCERT
  tls.key: $SSLKEY
EOF

cat << EOF > /var/lib/rancher/k3s/server/manifests/devtool/builtin-registry-certs.yaml
---
apiVersion: v1
kind: Secret
metadata:
  name: builtin-registry-certs
  labels:
    app: devtool
type: kubernetes.io/tls
data:
  ca.crt: $CACERT
  tls.crt: $SSLCERT
  tls.key: $SSLKEY
EOF

cat << EOF > /var/lib/rancher/k3s/server/manifests/devtool/ws-manager-tls.yaml
---
apiVersion: v1
kind: Secret
metadata:
  name: ws-manager-tls
  labels:
    app: devtool
type: kubernetes.io/tls
data:
  ca.crt: $CACERT
  tls.crt: $SSLCERT
  tls.key: $SSLKEY
EOF

cat << EOF > /var/lib/rancher/k3s/server/manifests/devtool/ws-manager-client-tls.yaml
---
apiVersion: v1
kind: Secret
metadata:
  name: ws-manager-client-tls
  labels:
    app: devtool
type: kubernetes.io/tls
data:
  ca.crt: $CACERT
  tls.crt: $SSLCERT
  tls.key: $SSLKEY
EOF

cat << EOF > /var/lib/rancher/k3s/server/manifests/devtool/ws-daemon-tls.yaml
---
apiVersion: v1
kind: Secret
metadata:
  name: ws-daemon-tls
  labels:
    app: devtool
type: kubernetes.io/tls
data:
  ca.crt: $CACERT
  tls.crt: $SSLCERT
  tls.key: $SSLKEY
EOF

cat << EOF > /var/lib/rancher/k3s/server/manifests/devtool/builtin-registry-facade-cert.yaml
---
apiVersion: v1
kind: Secret
metadata:
  name: builtin-registry-facade-cert
  labels:
    app: devtool
type: kubernetes.io/tls
data:
  ca.crt: $CACERT
  tls.crt: $SSLCERT
  tls.key: $SSLKEY
EOF

/devtool-installer init > config.yaml
yq e -i '.domain = "'"${DOMAIN}"'"' config.yaml
yq e -i '.certificate.name = "https-certificates"' config.yaml
yq e -i '.certificate.kind = "secret"' config.yaml
yq e -i '.customCACert.name = "ca-key-pair"' config.yaml
yq e -i '.customCACert.kind = "secret"' config.yaml
yq e -i '.observability.logLevel = "debug"' config.yaml
yq e -i '.workspace.runtime.containerdSocket = "/run/k3s/containerd/containerd.sock"' config.yaml
yq e -i '.workspace.runtime.containerdRuntimeDir = "/var/lib/rancher/k3s/agent/containerd/io.containerd.runtime.v2.task/k8s.io/"' config.yaml
yq e -i '.workspace.pvc.size = "10Gi"' config.yaml
yq e -i '.workspace.resources.requests.memory = "500Mi"' config.yaml
yq e -i '.workspace.resources.requests.cpu = "500m"' config.yaml
yq e -i '.experimental.telemetry.data.platform = "local-preview"' config.yaml

echo "extracting images to download ahead..."
/devtool-installer render --use-experimental-config --config config.yaml | grep 'image:' | sed 's/ *//g' | sed 's/image://g' | sed 's/\"//g' | sed 's/^-//g' | sort | uniq > /devtool-images.txt
echo "downloading images..."
while read -r image "$(cat /devtool-images.txt)"; do
   # shellcheck disable=SC2154
   ctr images pull "$image" >/dev/null &
done

ctr images pull "docker.io/devtool/workspace-full:latest" >/dev/null &

echo "images pulled"
/devtool-installer render --use-experimental-config --config config.yaml --output-split-files /var/lib/rancher/k3s/server/manifests/devtool

# store files in `devtool.debug` for debugging purposes
for f in /var/lib/rancher/k3s/server/manifests/devtool/*.yaml; do (cat "$f"; echo) >> /var/lib/rancher/k3s/server/devtool.debug; done
# remove unused resources
rm /var/lib/rancher/k3s/server/manifests/devtool/*NetworkPolicy*
rm /var/lib/rancher/k3s/server/manifests/devtool/*Certificate*
rm /var/lib/rancher/k3s/server/manifests/devtool/*Issuer*
# update PersistentVolumeClaim's to use k3s's `local-path` storage class
for f in /var/lib/rancher/k3s/server/manifests/devtool/*PersistentVolumeClaim*.yaml; do yq e -i '.spec.storageClassName="local-path"' "$f"; done
# Set `volumeClassTemplate` so that each replica creates its own PVC
# update Statefulset's to use k3s's `local-path` storage class
for f in /var/lib/rancher/k3s/server/manifests/devtool/*StatefulSet*.yaml; do yq e -i '.spec.volumeClaimTemplates[0].spec.storageClassName="local-path"' "$f"; done

# removing init container from ws-daemon (systemd and Ubuntu)
yq eval-all -i 'del(.spec.template.spec.initContainers[0])' /var/lib/rancher/k3s/server/manifests/devtool/*_DaemonSet_ws-daemon.yaml

# set lower requirements
yq eval-all -i '.spec.template.spec.containers[0].resources.requests.memory="250Mi"' /var/lib/rancher/k3s/server/manifests/devtool/*_DaemonSet_ws-daemon.yaml
yq eval-all -i '.spec.template.spec.containers[0].resources.requests.cpu="250m"' /var/lib/rancher/k3s/server/manifests/devtool/*_DaemonSet_ws-daemon.yaml
yq eval-all -i '.spec.template.spec.containers[0].resources.requests.memory="250Mi"' /var/lib/rancher/k3s/server/manifests/devtool/*_Deployment_minio.yaml

# set storage requests to be lower
for f in /var/lib/rancher/k3s/server/manifests/devtool/*PersistentVolumeClaim*.yaml; do
    yq e -i '.spec.resources.requests.storage="1Gi"' "$f";
done

for f in /var/lib/rancher/k3s/server/manifests/devtool/*StatefulSet*.yaml; do
    yq e -i '.spec.volumeClaimTemplates[0].spec.resources.requests.storage="1Gi"' "$f";
done

touch /var/lib/rancher/k3s/server/manifests/coredns.yaml.skip
mv -f /app/manifests/coredns.yaml /var/lib/rancher/k3s/server/manifests/custom-coredns.yaml

for f in /var/lib/rancher/k3s/server/manifests/devtool/*.yaml; do (cat "$f"; echo) >> /var/lib/rancher/k3s/server/manifests/devtool.yaml; done
rm -rf /var/lib/rancher/k3s/server/manifests/devtool

echo "manifests generated"
# waits for devtool pods to be ready, and manually runs the `devtool-telemetry` cronjob
run_telemetry(){
  # wait for the k3s cluster to be ready and Devtool workloads are added
  sleep 100
  # indefinitely wait for Devtool pods to be ready
  kubectl wait --timeout=-1s --for=condition=ready pod -l app=devtool,component!=migrations
  echo "Devtool pods are ready"
  # honour DO_NOT_TRACK if set
  if [ -n "${DO_NOT_TRACK}" ] && [ "${DO_NOT_TRACK}" -eq 1 ]; then
    # suspend the cronjob
    kubectl patch cronjobs devtool-telemetry -p '{"spec" : {"suspend" : true }}'
  else
    # manually run the cronjob
    kubectl create job devtool-telemetry-init --from=cronjob/devtool-telemetry
  fi
}

run_telemetry 2>&1 &

/bin/k3s server --disable traefik \
  --node-label devtool.io/workload_meta=true \
  --node-label devtool.io/workload_ide=true \
  --node-label devtool.io/workload_workspace_services=true \
  --node-label devtool.io/workload_services=true \
  --node-label devtool.io/workload_workspace_regular=true \
  --node-label devtool.io/workload_workspace_headless=true
