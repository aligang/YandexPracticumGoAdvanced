//go:build ignore

package main // generate.go

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const fileName = "build_info.go"

func main() {
	buildParams := map[string]string{}
	year, month, day := time.Now().Date()
	buildParams["buildDate"] = fmt.Sprintf("%d-%s-%d", year, month.String(), day)

	gitTag, err := exec.Command("git", "rev-parse", "HEAD").Output()
	if err != nil {
		panic(err)
	}
	buildParams["buildTag"] = string(gitTag[:len(gitTag)-1])

	buildParams["buildVersion"] = updateBuildVersion()
	var sb strings.Builder
	
	fmt.Fprintf(&sb, "%s", `
// Code generated by go generate; DO NOT EDIT.
// This file was generated by genconstants.go

package main

import "fmt"

var (
`)

	for k, v := range buildParams {
		fmt.Fprintf(&sb, "%s = \"%s\"\n", k, v)
	}
	fmt.Fprintf(&sb, `)

func printBuildInfo() {
`)
	fmt.Fprintf(&sb, "fmt.Println(\"Build version: ")
	fmt.Fprintf(&sb, buildParams["buildVersion"])
	fmt.Fprintf(&sb, "\")\n")

	fmt.Fprintf(&sb, "fmt.Println(\"Build date: ")
	fmt.Fprintf(&sb, buildParams["buildDate"])
	fmt.Fprintf(&sb, "\")\n")

	fmt.Fprintf(&sb, "fmt.Println(\"Build tag: ")
	fmt.Fprintf(&sb, buildParams["buildTag"])
	fmt.Fprintf(&sb, "\")\n")

	fmt.Fprintf(&sb, "}")

	//fmt.Println(sb.String())
	generated := []byte(sb.String())
	formatted, err := format.Source(generated)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(fileName, formatted, 0644)
	if err != nil {
		panic(err)
	}

}

func updateBuildVersion() string {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fileName, nil, 0)
	if err != nil {
		panic(err)
	}
	var buildVersion int64 = 0

	ast.Inspect(f, func(n ast.Node) bool {
		// проверяем, какой конкретный тип лежит в узле
		if x, ok := n.(*ast.ValueSpec); ok {
			for _, name := range x.Names {
				if name.Name == "buildVersion" {
					for _, value := range x.Values {
						if v, ok := value.(*ast.BasicLit); ok {
							r, err := strconv.ParseInt(strings.Trim(v.Value, "\""), 10, 64)
							if err == nil {
								buildVersion = r + 1
							}
						}
					}
				}
			}
		}
		return true
	})

	return strconv.FormatInt(buildVersion, 10)
}
