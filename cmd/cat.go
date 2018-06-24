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
	"path/filepath"

	"github.com/nemca/nemstore/internal/crypto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat file",
	Short: "Print encrypted file",
	Args:  cobra.ExactArgs(1),
	Run:   catCmdRun,
}

func init() {
	rootCmd.AddCommand(catCmd)
}

func catCmdRun(cmd *cobra.Command, args []string) {
	passphrase, err := crypto.ReadPassphrase()
	if err != nil {
		log.Fatal(err)
	}

	storageDir := viper.GetString("StorageDir")
	if storageDir == "" {
		log.Fatalf("option `StorageDir` not set in %s\n", viper.ConfigFileUsed())
	}

	path := filepath.Join(storageDir, args[0])
	data, err := crypto.DecryptFile(path, passphrase)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(data))
}
