/*
Copyright © 2019 Portworx

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
package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func executeCli(cli string) []string {
	fmt.Println(cli)
	so, _, r := pxTestSetupCli(cli)

	// Defer to cleanup state
	defer r()

	// Start the CLI
	Execute()

	return strings.Split(so.String(), "\n")
}

// Takes a volume name and size. Returns the created volume id.
// For some reason our test container only recoganizes id and not name for some calls.
func testCreateVolume(t *testing.T, volName string, size uint64) string {
	cli := "px create volume --name " + volName + " --size " + strconv.FormatUint(size, 10)
	lines := executeCli(cli)
	assert.Equal(t, 2, len(lines), "Output does not match")
	assert.Contains(t, lines[0], "Volume "+volName+" created with id", "expected message not received")
	words := strings.Split(lines[0], " ")
	assert.Equal(t, len(words), 6, "expected message not received")
	// The last item in the message is the id
	return words[len(words)-1]
}

// Takes a volume id and returns true if the volume is found
func testGetVolume(t *testing.T, volId string, volName string) bool {
	cli := "px get volume -o wide " + volId
	lines := executeCli(cli)
	if strings.Contains(lines[0], "Error: Failed to get volume: Volume id "+volId+" not found") {
		return false
	}
	w := strings.Split(lines[2], " ")
	// Remove all blanks
	words := make([]string, 0, len(w))
	for _, item := range w {
		if item != "" {
			words = append(words, item)
		}
	}
	// Verify the first item is the id of the volume
	assert.Equal(t, volId, words[0])
	assert.Equal(t, volName, words[1])
	return true
}

//Returns a list of all volume ids
func testGetAllVolumes(t *testing.T) ([]string, []string) {
	cli := "px get volume -o wide"
	lines := executeCli(cli)
	lines = lines[2:]
	volIds := make([]string, 0, len(lines))
	volNames := make([]string, 0, len(lines))
	for _, l := range lines {
		if l == "" {
			continue
		}
		x := strings.Split(l, " ")
		volIds = append(volIds, x[0])
		volNames = append(volNames, x[1])
	}
	return volIds, volNames
}

// Deletes specified volume
func testDeleteVolume(t *testing.T, volName string) {
	cli := "px delete volume --name " + volName
	lines := executeCli(cli)
	fmt.Println(len(lines), lines)
}

// Takes a volume name and snapshot name. Returns the created snapshot's volume id.
// For some reason our test container only recoganizes id and not name for some calls.
func testCreateSnapshot(t *testing.T, volId string, snapName string) string {
	cli := "px create snapshot --name " + snapName + " --volume " + volId
	lines := executeCli(cli)
	assert.Equal(t, 2, len(lines), "Output does not match")
	assert.Contains(t, lines[0], "Snapshot of "+volId+" created with id", "expected message not received")
	fmt.Println(lines[0])
	words := strings.Split(lines[0], " ")
	assert.Equal(t, len(words), 7, "expected message not received")
	// The last item in the message is the id
	return words[len(words)-1]
}
