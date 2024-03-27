//go:build linux
// +build linux

// SPDX-License-Identifier: Apache-2.0
/*
Copyright (C) 2023 The Falco Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package syscall

import (
	"path/filepath"

	"github.com/falcosecurity/event-generator/events"

	"os"
)

var _ = events.Register(
	ReadSshInformation,
	// events.WithDisabled(), // the rule is not enabled by default, so disable the action too
)

func ReadSshInformation(h events.Helper) error {
	// Also creates .ssh directory inside tempDirectory
	tempDirectoryName, err := CreateSshDirectoryUnderHome()
	if err != nil {
		return err
	}
	sshDir := filepath.Join(tempDirectoryName, ".ssh")
	defer os.RemoveAll(tempDirectoryName)

	// Create known_hosts file. os.Create is enough to trigger the rule
	filename := filepath.Join(sshDir, "known_hosts")
	if _, err := os.Create(filename); err != nil {
		return err
	}

	h.Log().Info("attempting to simulate SSH information read")

	return nil
}
