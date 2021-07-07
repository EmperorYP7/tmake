/*
Package cmd tmake
Copyright © 2021 Yash Pandey

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
	"bufio"
	"io"
	"log"
	"os"
)

// Copy : copying a file src to dst
func Copy(src, dst string) error {
    in, err := os.Open(src)
    if err != nil {
        return err
    }
    defer in.Close()

    out, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, in)
    if err != nil {
        return err
    }
    return out.Close()
}

func ReadRemoteFile(src string) ([]string, error) {
    readFile, err := os.Open(src)
 
    if err != nil {
        log.Fatalf("failed to open file: %s", err)
        return nil, err
    }
 
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileTextLines []string
 
    for fileScanner.Scan() {
        fileTextLines = append(fileTextLines, fileScanner.Text())
    }
 
    readFile.Close()
    return fileTextLines, nil
}

func GetRemoteNames(src string) ([]string, error) {
    fileTextLines, err := ReadRemoteFile(src)
    if err != nil {
        return nil, err
    }
    var remoteNames []string

    for _, eachline := range fileTextLines {
        remoteNames = append(remoteNames, eachline)
    }

    return remoteNames, nil
}
