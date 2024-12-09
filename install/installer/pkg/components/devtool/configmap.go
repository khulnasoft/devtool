// Copyright (c) 2021 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package devtool

import (
	"fmt"

	"github.com/khulnasoft/devtool/installer/pkg/common"
	"github.com/khulnasoft/devtool/installer/pkg/config"
	"github.com/khulnasoft/devtool/installer/pkg/config/versions"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type Devtool struct {
	VersionManifest versions.Manifest `json:"versions"`
}

func configmap(ctx *common.RenderContext) ([]runtime.Object, error) {
	gpversions := Devtool{
		VersionManifest: ctx.VersionManifest,
	}

	cfg, err := config.Marshal(config.CurrentVersion, ctx.Config)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Devtool config: %w", err)
	}

	versions, err := common.ToJSONString(gpversions)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Devtool versions: %w", err)
	}

	return []runtime.Object{
		&corev1.ConfigMap{
			TypeMeta: common.TypeMetaConfigmap,
			ObjectMeta: metav1.ObjectMeta{
				Name:      Component,
				Namespace: ctx.Namespace,
				Labels:    common.DefaultLabels(Component),
			},
			Data: map[string]string{
				"config.yaml":   string(cfg),
				"versions.json": string(versions),
			},
		},
	}, nil
}
