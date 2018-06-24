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

package cmd

import (
	"fmt"
	"log"

	"github.com/nemca/nemstore/internal/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List of stored files",
	Args:  cobra.NoArgs,
	Run:   lsCmdRun,
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

func lsCmdRun(cmd *cobra.Command, args []string) {
	storageDir := viper.GetString("StorageDir")
	if storageDir == "" {
		log.Fatalf("Option `StorageDir` not set in %s\n", viper.ConfigFileUsed())
	}

	files, err := storage.ReadDir(storageDir)
	if err != nil {
		log.Fatalf("Can't read storage dir: %v\n", err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}
