//  Copyright 2024 Google LLC
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

//go:build linux

// Package main is the entry point for the google authorized keys compat. It is
// responsible for enabling either the new authorized keys system or that in the
// legacy guest agent.
package main

import (
	"context"

	"github.com/GoogleCloudPlatform/galog"
	"github.com/GoogleCloudPlatform/google-guest-agent/internal/metadata"
)

func launchAuthorizedKeys(ctx context.Context, mdsClient metadata.MDSClientInterface, username string) error {
	galog.Infof("Not implemented for Linux")
	return nil
}
