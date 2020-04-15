package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

/*
Reference:
https://ninedegreesbelow.com/photography/exiftool-commands.html
https://www.linux.com/training-tutorials/how-sort-and-remove-duplicate-photos-linux/
*/

//movePic moves a file between directories
func movePic(src, dest string) {
	fmt.Printf("source: %s, destination: %s\n", src, dest)
	absSrc, _ := filepath.Abs(src)
	absDest, _ := filepath.Abs(dest)
	err := os.Rename(absSrc, absDest)
	check(err)
}

//getPath returns a full path in the form of %Y/%m as string after processing
//the filename
func getPath(input string) string {
	reYearMonth := regexp.MustCompile(`^\d*-\d*`)                //regex to obtain year-month
	reSplit := regexp.MustCompile(`-`)                           //regex to split year-month
	pathList := reSplit.Split(reYearMonth.FindString(input), -1) //splitting into a list of [year, month]
	path := fmt.Sprintf("./%s/%s/", pathList[0], pathList[1])    //formatting "year/month"
	//fmt.Printf("year: %s, month: %s, path: %s\n", pathList[0], pathList[1], path)
	return path
}

func createDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		check(err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Move() {
	files, err := ioutil.ReadFile("./error.log")                  //reading files
	check(err)                                                    //Checking for errors
	scanner := bufio.NewScanner(strings.NewReader(string(files))) //obtaining a scanner
	var fileName, src, dest, path string
	//scanning line by line
	for scanner.Scan() {
		fileName = strings.ReplaceAll(scanner.Text(), "'", "")
		src = "./" + fileName
		if _, err := os.Stat(src); os.IsNotExist(err) { //if src does not exist
			log.Println(src)
			continue
		} 		
		//path = getPath(strings.ReplaceAll(fileName, "'", ""))
		path = getPath(fileName)
		dest = path + fileName
		createDir(path)
		//fmt.Print(src, dest)
		movePic(src, dest)
	}
}
