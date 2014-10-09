//  Copyright 2013 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var exceptions = []string{
	"asm/errors/errors.go",
	"asm/lexer/acttab.go",
	"asm/lexer/lexer.go",
	"asm/lexer/transitiontable.go",
	"asm/parser/action.go",
	"asm/parser/actiontable.go",
	"asm/parser/gototable.go",
	"asm/parser/productionstable.go",
	"asm/test/person.proto",
	"asm/token/token.go",
	"asm/test/srctree.proto",
	"asm/test/puddingmilkshake.proto",
	"asm/test/taxonomy.proto",
	"asm/test/treeregister.proto",
	"asm/test/typewriterprison.proto",
	"list_of_functions.txt",
	"proto/tokens/test.proto",
	"asm/parser/parser.go",
}

func main() {
	if len(os.Args) == 0 {
		fmt.Fprintf(os.Stderr, "no folder specified\n")
		os.Exit(1)
	}
	files := []string{}
	err := filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		for _, filepart := range exceptions {
			if strings.HasSuffix(path, filepart) {
				return nil
			}
		}
		if strings.Contains(path, ".git") {
			return nil
		}
		base := filepath.Base(path)
		if base == "LICENSE" {
			return nil
		}
		if base == ".gitignore" {
			return nil
		}
		if strings.HasPrefix(strings.ToLower(base), "readme") {
			return nil
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		datas := string(data)
		if strings.HasPrefix(datas, "// Code generated by") {
			return nil
		}
		datas = strings.Replace(datas, "//", "", 1)
		datas = strings.Replace(datas, "#", "", 1)
		datas = strings.TrimSpace(datas)
		if !strings.HasPrefix(datas, "Copyright") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if len(files) == 0 {
		return
	}
	fmt.Fprintf(os.Stderr, "ERROR The following files still needs a LICENSE: [%s]\n", strings.Join(files, ", "))
	os.Exit(1)
}
