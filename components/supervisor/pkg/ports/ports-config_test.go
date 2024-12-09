// Copyright (c) 2020 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package ports

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	devtool "github.com/khulnasoft/devtool/devtool-protocol"
)

func TestPortsConfig(t *testing.T) {
	tests := []struct {
		Desc         string
		DevtoolConfig *devtool.DevtoolConfig
		Expectation  *PortConfigTestExpectations
	}{
		{
			Desc:        "no configs",
			Expectation: &PortConfigTestExpectations{},
		},
		{
			Desc: "instance port config",
			DevtoolConfig: &devtool.DevtoolConfig{
				Ports: []*devtool.PortsItems{
					{
						Port:        9229,
						OnOpen:      "ignore",
						Visibility:  "public",
						Name:        "Nice Port Name",
						Description: "Nice Port Description",
					},
				},
			},
			Expectation: &PortConfigTestExpectations{
				InstancePortConfigs: []*devtool.PortConfig{
					{
						Port:        9229,
						OnOpen:      "ignore",
						Visibility:  "public",
						Name:        "Nice Port Name",
						Description: "Nice Port Description",
					},
				},
			},
		},
		{
			Desc: "instance range config",
			DevtoolConfig: &devtool.DevtoolConfig{
				Ports: []*devtool.PortsItems{
					{
						Port:        "9229-9339",
						OnOpen:      "ignore",
						Visibility:  "public",
						Name:        "Nice Port Name",
						Description: "Nice Port Description",
					},
				},
			},
			Expectation: &PortConfigTestExpectations{
				InstanceRangeConfigs: []*RangeConfig{
					{
						PortsItems: devtool.PortsItems{
							Port:        "9229-9339",
							OnOpen:      "ignore",
							Visibility:  "public",
							Description: "Nice Port Description",
							Name:        "Nice Port Name",
						},
						Start: 9229,
						End:   9339,
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.Desc, func(t *testing.T) {
			configService := &testDevtoolConfigService{
				configs: make(chan *devtool.DevtoolConfig),
				changes: make(chan *struct{}),
			}
			defer close(configService.configs)
			defer close(configService.changes)

			context, cancel := context.WithCancel(context.Background())
			defer cancel()

			workspaceID := "test"

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := NewConfigService(workspaceID, configService)
			updates, errors := service.Observe(context)

			actual := &PortConfigTestExpectations{}

			if test.DevtoolConfig != nil {
				go func() {
					configService.configs <- test.DevtoolConfig
				}()
				select {
				case err := <-errors:
					t.Fatal(err)
				case change := <-updates:
					actual.InstanceRangeConfigs = change.instanceRangeConfigs
					for _, config := range change.instancePortConfigs {
						actual.InstancePortConfigs = append(actual.InstancePortConfigs, &config.PortConfig)
					}
				}
			}

			if diff := cmp.Diff(test.Expectation, actual); diff != "" {
				t.Errorf("unexpected output (-want +got):\n%s", diff)
			}
		})
	}
}

type PortConfigTestExpectations struct {
	InstancePortConfigs  []*devtool.PortConfig
	InstanceRangeConfigs []*RangeConfig
}

type testDevtoolConfigService struct {
	configs chan *devtool.DevtoolConfig
	changes chan *struct{}
}

func (service *testDevtoolConfigService) Watch(ctx context.Context) {
}

func (service *testDevtoolConfigService) Observe(ctx context.Context) <-chan *devtool.DevtoolConfig {
	return service.configs
}

func (service *testDevtoolConfigService) ObserveImageFile(ctx context.Context) <-chan *struct{} {
	return service.changes
}
