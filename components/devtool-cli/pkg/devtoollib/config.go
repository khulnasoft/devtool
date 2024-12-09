// Copyright (c) 2020 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package devtoollib

type DevtoolImage struct {
	File    string
	Context string `yaml:"context,omitempty"`
}

type devtoolPort struct {
	Number int32 `yaml:"port"`
}

type devtoolTask struct {
	Command string `yaml:"command,omitempty"`
	Init    string `yaml:"init,omitempty"`
}

type DevtoolFile struct {
	Image             interface{}  `yaml:"image,omitempty"`
	Ports             []devtoolPort `yaml:"ports,omitempty"`
	Tasks             []devtoolTask `yaml:"tasks,omitempty"`
	CheckoutLocation  string       `yaml:"checkoutLocation,omitempty"`
	WorkspaceLocation string       `yaml:"workspaceLocation,omitempty"`
}

// SetImageName configures a pre-built docker image by name
func (cfg *DevtoolFile) SetImageName(name string) {
	if name == "" {
		return
	}
	cfg.Image = name
}

// SetImage configures a Dockerfile as workspace image
func (cfg *DevtoolFile) SetImage(img DevtoolImage) {
	cfg.Image = img
}

// AddPort adds a port to the list of exposed ports
func (cfg *DevtoolFile) AddPort(port int32) {
	cfg.Ports = append(cfg.Ports, devtoolPort{
		Number: port,
	})
}

// AddTask adds a workspace startup task
func (cfg *DevtoolFile) AddTask(task ...string) {
	if len(task) > 1 {
		cfg.Tasks = append(cfg.Tasks, devtoolTask{
			Command: task[0],
			Init:    task[1],
		})
	} else {
		cfg.Tasks = append(cfg.Tasks, devtoolTask{
			Command: task[0],
		})
	}
}
