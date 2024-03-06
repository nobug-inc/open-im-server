// Copyright © 2024 OpenIM. All rights reserved.
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

package genutil

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/OpenIMSDK/tools/errs"
)

// OutDir creates the absolute path name from path and checks path exists.
// Returns absolute path including trailing '/' or error if path does not exist.
func OutDir(path string) (string, error) {
	outDir, err := filepath.Abs(path)
	if err != nil {
		return "", errs.Wrap(err, "output directory %s does not exist", path)
	}

	stat, err := os.Stat(outDir)
	if err != nil {
		return "", errs.Wrap(err, "output directory %s does not exist", outDir)
	}

	if !stat.IsDir() {
		return "", errs.Wrap(err, "output directory %s is not a directory", outDir)
	}
	outDir += "/"
	return outDir, nil
}

func ExitWithError(err error) {
	progName := filepath.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "%s exit -1: %+v\n", progName, err)
	os.Exit(-1)
}

func SIGTERMExit() {
	progName := filepath.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "Warning %s receive process terminal SIGTERM exit 0\n", progName)
}
