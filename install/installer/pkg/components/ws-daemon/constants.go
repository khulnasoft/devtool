// Copyright (c) 2021 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package wsdaemon

import "github.com/khulnasoft/devtool/common-go/baseserver"

const (
	Component               = "ws-daemon"
	ServicePort             = 8080
	HostWorkingAreaMk2      = "/var/devtool/workspaces-mk2"
	ContainerWorkingAreaMk2 = "/mnt/workingarea-mk2"
	HostBackupPath          = "/var/devtool/tmp/backup"
	TLSSecretName           = "ws-daemon-tls"
	VolumeTLSCerts          = "ws-daemon-tls-certs"
	ReadinessPort           = baseserver.BuiltinHealthPort
)
