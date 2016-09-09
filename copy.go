package main

import (
	"os"
	"log"
	"io"
	"sync"
	"io/ioutil"
	"errors"
	"runtime"
)

func main() {

	//TODO read these from input
	destinationsPaths := []string{"/Users/gfalace/Documents/destination", "/Users/gfalace/Downloads/destination"}
	fileOrigin := "/Users/gfalace/origin"

	processors := min(len(destinationsPaths), runtime.NumCPU())
	log.Printf("Runtime using %d processors", processors)
	runtime.GOMAXPROCS(processors)

	var w sync.WaitGroup

	w.Add(2)

	for _, destination := range destinationsPaths {
		go func(origin, dest string) {
			log.Printf("in the go routine with params origin: %s  -->  dest: %s", origin, dest)
			err := copyDir(origin, dest)
			if err != nil {
				log.Print(err)
			}
			log.Printf("Finished copying to destination %s", dest)
			w.Done()
		}(fileOrigin, destination)
	}

	w.Wait()
}

func copyDir(source, dest string) error {

	fileInfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	if !fileInfo.IsDir() {
		return errors.New("source is not a folder!")
	}

	_, err = os.Open(dest)
	if !os.IsNotExist(err) {
		return errors.New("destination folder already exists! " + dest)
	}

	err = os.MkdirAll(dest, fileInfo.Mode())
	log.Print("created dir: " + dest)
	if err != nil {
		return errors.New("could not create detination folder!")
	}

	elementInDirectory, err := ioutil.ReadDir(source)

	for _, element := range elementInDirectory {

		sourceFilePath := source + "/" + element.Name()
		destinationPathFile := dest + "/" + element.Name()

		if element.IsDir() {
			err = copyDir(sourceFilePath, destinationPathFile)
			if err != nil {
				log.Println(err)
			}
		} else {
			err = copyFile(sourceFilePath, destinationPathFile)
			if err != nil {
				log.Println(err)
			}
		}

	}
	return nil
}

func copyFile(source, dest string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return errors.New("error opening source file: " + source)
	}
	defer sourceFile.Close()
	destFile, err := os.Create(dest)
	if err != nil {
		return errors.New("error creating destination file: " + dest)
	}
	defer destFile.Close()
	_, err = io.Copy(destFile, sourceFile)
	if err == nil {
		fileInfo, err := os.Stat(source)
		if err == nil {
			err = os.Chmod(dest, fileInfo.Mode())
		} else {
			errors.New("error getting source stats (while trying to set same permissions as source): " + source)
		}
	} else {
		errors.New("error copying file: " + source)
	}
	log.Print("created: " + dest)
	return nil
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
