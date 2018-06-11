// https://app.pluralsight.com/player?course=go&author=john-sonmez&name=go-m4-functions&clip=4&mode=live

// cd ~/go/src/learning_go && go install && ../../bin/learning_go

// In more basic setups you can:
// run using "go run hello_world.go"
// or "go build hello_world.go" and ./hello_world
// OR got src, put file in repo/project, "go install ./repo/project", thing executeable will be generated and put into a bin folder

package main

import (
	"fmt"
	"reflect"
	"runtime"

	"learning_go/chapter"
)

func main() {
	runChapter(chapter.VarsTypesPointers)
	runChapter(chapter.Functions)
	runChapter(chapter.Branching)
	runChapter(chapter.Loops)
	runChapter(chapter.Maps)
	runChapter(chapter.Slices)
}

type chapterFunc func()

func runChapter(thisChapter chapterFunc) {
	fmt.Println("------------ [" + getFunctionName(thisChapter) + "] -------------")
	thisChapter()
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// https://golang.org/ref/spec
// https://github.com/golang/go/wiki/CodeReviewComments
