// Copyright (c) 2023 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package utils

import (
	"errors"
	"os"
	"path/filepath"

	devtool "github.com/khulnasoft/devtool/devtool-protocol"
	yaml "gopkg.in/yaml.v2"
)

func ParseDevtoolConfig(repoRoot string) (*devtool.DevtoolConfig, error) {
	if repoRoot == "" {
		return nil, errors.New("repoRoot is empty")
	}
	data, err := os.ReadFile(filepath.Join(repoRoot, ".devtool.yml"))
	if err != nil {
		// .devtool.yml not exist is ok
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, errors.New("read .devtool.yml file failed: " + err.Error())
	}
	var config *devtool.DevtoolConfig
	if err = yaml.Unmarshal(data, &config); err != nil {
		return nil, errors.New("unmarshal .devtool.yml file failed" + err.Error())
	}
	return config, nil
}
