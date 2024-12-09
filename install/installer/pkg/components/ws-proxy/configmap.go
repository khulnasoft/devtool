// Copyright (c) 2021 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package wsproxy

import (
	"fmt"
	"time"

	"github.com/khulnasoft/devtool/installer/pkg/components/workspace"
	wsmanagermk2 "github.com/khulnasoft/devtool/installer/pkg/components/ws-manager-mk2"
	configv1 "github.com/khulnasoft/devtool/installer/pkg/config/v1"
	"github.com/khulnasoft/devtool/installer/pkg/config/v1/experimental"

	"github.com/khulnasoft/devtool/common-go/baseserver"
	"github.com/khulnasoft/devtool/common-go/util"
	"github.com/khulnasoft/devtool/installer/pkg/common"
	"github.com/khulnasoft/devtool/ws-proxy/pkg/config"
	"github.com/khulnasoft/devtool/ws-proxy/pkg/proxy"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func configmap(ctx *common.RenderContext) ([]runtime.Object, error) {
	header := HostHeader
	blobServeHost := fmt.Sprintf("ide.%s", ctx.Config.Domain)
	devtoolInstallationHostName := ctx.Config.Domain

	installationShortNameSuffix := ""
	if ctx.Config.Metadata.InstallationShortname != "" && ctx.Config.Metadata.InstallationShortname != configv1.InstallationShortNameOldDefault {
		installationShortNameSuffix = "-" + ctx.Config.Metadata.InstallationShortname
	}

	devtoolInstallationWorkspaceHostSuffix := fmt.Sprintf(".ws%s.%s", installationShortNameSuffix, ctx.Config.Domain)
	devtoolInstallationWorkspaceHostSuffixRegex := fmt.Sprintf("\\.ws[^\\.]*\\.%s", ctx.Config.Domain)

	wsManagerConfig := &config.WorkspaceManagerConn{
		Addr: fmt.Sprintf("ws-manager-mk2:%d", wsmanagermk2.RPCPort),
		TLS: struct {
			CA   string "json:\"ca\""
			Cert string "json:\"crt\""
			Key  string "json:\"key\""
		}{
			CA:   "/ws-manager-client-tls-certs/ca.crt",
			Cert: "/ws-manager-client-tls-certs/tls.crt",
			Key:  "/ws-manager-client-tls-certs/tls.key",
		},
	}

	ctx.WithExperimental(func(ucfg *experimental.Config) error {
		if ucfg.Workspace == nil {
			return nil
		}
		if ucfg.Workspace.WSProxy.IngressHeader != "" {
			header = ucfg.Workspace.WSProxy.IngressHeader
		}
		if ucfg.Workspace.WSProxy.BlobServeHost != "" {
			blobServeHost = ucfg.Workspace.WSProxy.BlobServeHost
		}
		if ucfg.Workspace.WSProxy.DevtoolInstallationHostName != "" {
			devtoolInstallationHostName = ucfg.Workspace.WSProxy.DevtoolInstallationHostName
		}
		if ucfg.Workspace.WSProxy.DevtoolInstallationWorkspaceHostSuffix != "" {
			devtoolInstallationWorkspaceHostSuffix = ucfg.Workspace.WSProxy.DevtoolInstallationWorkspaceHostSuffix
		}
		if ucfg.Workspace.WSProxy.DevtoolInstallationWorkspaceHostSuffixRegex != "" {
			devtoolInstallationWorkspaceHostSuffixRegex = ucfg.Workspace.WSProxy.DevtoolInstallationWorkspaceHostSuffixRegex
		}

		return nil
	})

	// todo(sje): wsManagerProxy seems to be unused
	wspcfg := config.Config{
		Namespace: ctx.Namespace,
		Ingress: proxy.HostBasedIngressConfig{
			HTTPAddress:  fmt.Sprintf("0.0.0.0:%d", HTTPProxyPort),
			HTTPSAddress: fmt.Sprintf("0.0.0.0:%d", HTTPSProxyPort),
			Header:       header,
		},
		Proxy: proxy.Config{
			HTTPS: struct {
				Key         string `json:"key"`
				Certificate string `json:"crt"`
			}{
				Key:         "/mnt/certificates/tls.key",
				Certificate: "/mnt/certificates/tls.crt",
			},
			TransportConfig: &proxy.TransportConfig{
				ConnectTimeout:      util.Duration(time.Second * 10),
				IdleConnTimeout:     util.Duration(time.Minute),
				MaxIdleConns:        0,
				MaxIdleConnsPerHost: 100,
			},
			BlobServer: &proxy.BlobServerConfig{
				Scheme:     "https",
				Host:       blobServeHost,
				PathPrefix: "/blobserve",
			},
			DevtoolInstallation: &proxy.DevtoolInstallation{
				Scheme:                   "https",
				HostName:                 devtoolInstallationHostName,
				WorkspaceHostSuffix:      devtoolInstallationWorkspaceHostSuffix,
				WorkspaceHostSuffixRegex: devtoolInstallationWorkspaceHostSuffixRegex,
			},
			WorkspacePodConfig: &proxy.WorkspacePodConfig{
				TheiaPort:               workspace.ContainerPort,
				IDEDebugPort:            workspace.IDEDebugPort,
				SupervisorPort:          workspace.SupervisorPort,
				SupervisorDebugPort:     workspace.SupervisorDebugPort,
				DebugWorkspaceProxyPort: workspace.DebugWorkspaceProxyPort,
				SupervisorImage:         ctx.ImageName(ctx.Config.Repository, workspace.SupervisorImage, ctx.VersionManifest.Components.Workspace.Supervisor.Version),
			},
			BuiltinPages: proxy.BuiltinPagesConfig{
				Location: "/app/public",
			},
		},
		PProfAddr:          common.LocalhostAddressFromPort(baseserver.BuiltinDebugPort),
		PrometheusAddr:     common.LocalhostPrometheusAddr(),
		ReadinessProbeAddr: fmt.Sprintf(":%v", ReadinessPort),
		WorkspaceManager:   wsManagerConfig,
	}

	if ctx.Config.SSHGatewayCAKey != nil {
		wspcfg.Proxy.SSHGatewayCAKeyFile = "/mnt/ca-key/ca.key"
	}

	fc, err := common.ToJSONString(wspcfg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ws-proxy config: %w", err)
	}

	return []runtime.Object{
		&corev1.ConfigMap{
			TypeMeta: common.TypeMetaConfigmap,
			ObjectMeta: metav1.ObjectMeta{
				Name:        Component,
				Namespace:   ctx.Namespace,
				Labels:      common.CustomizeLabel(ctx, Component, common.TypeMetaConfigmap),
				Annotations: common.CustomizeAnnotation(ctx, Component, common.TypeMetaConfigmap),
			},
			Data: map[string]string{
				"config.json": string(fc),
			},
		},
	}, nil
}
