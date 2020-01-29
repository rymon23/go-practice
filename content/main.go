package content

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

var PathBase, PathErr = filepath.Abs(filepath.Dir("go-practice/frontend/"))
var PathPages = PathBase + "/"
var PathContent = PathBase + "/content/"
var SavedContent []string

func Get()  {
	files, err := ioutil.ReadDir(PathContent)
	if err != nil {
		log.Fatal(err)
	}

	SavedContent = make([]string, len(files))

	for i, f := range files {
		// fmt.Println(f.Name())
		SavedContent[i] = f.Name()
	}
	fmt.Printf("Loaded Content: \n\n %v \n\n", SavedContent)
}