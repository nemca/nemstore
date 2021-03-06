// Copyright © 2018 Michael Bruskov <mixanemca@yandex.ru>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

// RunEditor creates a temporary file and launches $EDITOR
func RunEditor(tmpFile *os.File) ([]byte, error) {
	editor := os.Getenv("EDITOR")
	editorPath, err := exec.LookPath(editor)
	if err != nil {
		return nil, err
	}

	editorCmd := exec.Command(editorPath, tmpFile.Name())
	editorCmd.Stdin = os.Stdin
	editorCmd.Stdout = os.Stdout
	editorCmd.Stderr = os.Stderr
	err = editorCmd.Start()
	if err != nil {
		return nil, err
	}
	err = editorCmd.Wait()
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		return nil, err
	}

	return data, err
}

// RunLess run `less` command with file
func RunLess(input io.Reader) error {
	lessPath, err := exec.LookPath("less")
	if err != nil {
		return err
	}

	lessCmd := exec.Command(lessPath)
	lessCmd.Stdin = input
	lessCmd.Stdout = os.Stdout
	lessCmd.Stderr = os.Stderr
	err = lessCmd.Start()
	if err != nil {
		return err
	}
	err = lessCmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

// CreateTempFile create new temporary file
func CreateTempFile() (*os.File, error) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "nemstore_")
	if err != nil {
		return nil, err
	}
	return tmpFile, nil
}

// ReadDir read storage directory and return files
func ReadDir(path string) ([]string, error) {
	var retFiles []string

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			retFiles = append(retFiles, file.Name())
		}
	}

	return retFiles, nil
}
