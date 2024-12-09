// Copyright (c) 2022 Devtool GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package config

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/khulnasoft/devtool/common-go/log"
	devtool "github.com/khulnasoft/devtool/devtool-protocol"
)

func TestAnalyzeDevtoolConfig(t *testing.T) {
	tests := []struct {
		Desc    string
		Prev    *devtool.DevtoolConfig
		Current *devtool.DevtoolConfig
		Fields  []string
	}{
		{
			Desc: "change",
			Prev: &devtool.DevtoolConfig{
				CheckoutLocation: "foo",
			},
			Current: &devtool.DevtoolConfig{
				CheckoutLocation: "bar",
			},
			Fields: []string{"CheckoutLocation"},
		},
		{
			Desc: "add",
			Prev: &devtool.DevtoolConfig{},
			Current: &devtool.DevtoolConfig{
				CheckoutLocation: "bar",
			},
			Fields: []string{"CheckoutLocation"},
		},
		{
			Desc: "remove",
			Prev: &devtool.DevtoolConfig{
				CheckoutLocation: "bar",
			},
			Current: &devtool.DevtoolConfig{},
			Fields:  []string{"CheckoutLocation"},
		},
		{
			Desc: "none",
			Prev: &devtool.DevtoolConfig{
				CheckoutLocation: "bar",
			},
			Current: &devtool.DevtoolConfig{
				CheckoutLocation: "bar",
			},
			Fields: nil,
		},
		{
			Desc:    "fie created",
			Current: &devtool.DevtoolConfig{},
			Fields:  nil,
		},
	}
	for _, test := range tests {
		t.Run(test.Desc, func(t *testing.T) {
			var fields []string
			analyzer := NewConfigAnalyzer(log.Log, 100*time.Millisecond, func(field string) {
				fields = append(fields, field)
			}, test.Prev)
			<-analyzer.Analyse(test.Current)
			if diff := cmp.Diff(test.Fields, fields); diff != "" {
				t.Errorf("unexpected output (-want +got):\n%s", diff)
			}
		})
	}
}
