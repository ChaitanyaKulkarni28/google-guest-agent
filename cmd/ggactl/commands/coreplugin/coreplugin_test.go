//  Copyright 2024 Google LLC
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package coreplugin

import (
	"context"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/google-guest-agent/cmd/ggactl/commands"
	"github.com/GoogleCloudPlatform/google-guest-agent/cmd/ggactl/commands/testhelper"
	"github.com/GoogleCloudPlatform/google-guest-agent/internal/cfg"
)

func TestCorePluginCommands(t *testing.T) {
	if err := cfg.Load(nil); err != nil {
		t.Fatalf("cfg.Load(nil) failed unexpectedly with error: %v", err)
	}
	ctx := context.WithValue(context.Background(), commands.TestOverrideKey, true)
	cmd := New()
	cmd.SetContext(ctx)

	tests := []struct {
		name       string
		args       []string
		wantErr    string
		shouldFail bool
	}{
		{
			name:       "no_subcommand_error",
			wantErr:    "no subcommand",
			shouldFail: true,
		},
		{
			name:       "invalid_args",
			args:       []string{"restart", "invalid_arg"},
			wantErr:    "unknown command",
			shouldFail: true,
		},
		{
			name:       "no_plugin_found",
			args:       []string{"restart"},
			wantErr:    "unable to restart core plugin",
			shouldFail: true,
		},
		{
			name:       "stop_success",
			args:       []string{"stop"},
			shouldFail: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out, err := testhelper.ExecuteCommand(ctx, cmd, test.args)
			if (err != nil) != test.shouldFail {
				t.Errorf("ExecuteCommand(%s, %v) = %v, want error: %t", cmd.Name(), test.args, err, test.shouldFail)
			}
			if test.wantErr == "" {
				return
			}
			if !strings.Contains(out, test.wantErr) {
				t.Errorf("ExecuteCommand(%s, %v) = %q, want error containing %q", cmd.Name(), test.args, out, test.wantErr)
			}
		})
	}
}
