package main

import (
	"os"
	"log"
	"io"
	"sync"
)

func main() {

	destinationsPaths := []string{"/Users/gfalace/Documents/", "/Users/gfalace/Downloads/"}
	fileOrigin := "/Users/gfalace/"
	fileName := "profile.jpg"

	var w sync.WaitGroup

	w.Add(2)

	for _, destination := range destinationsPaths {
		go func(mountPoint, fileOrigin, fileName string) {

			defer w.Done()

			originalFile, err := os.Open(fileOrigin + fileName)
			if err != nil {
				log.Fatal(err)
			}
			defer originalFile.Close()

			os.Mkdir(mountPoint +"copies", 0777);

			newFile, err := os.Create(mountPoint + "copies/copy_of_" + fileName)
			if err != nil {
				log.Fatal(err)
			}
			defer newFile.Close()

			// Copy the bytes to destination from source
			bytesWritten, err := io.Copy(newFile, originalFile)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Copied %d bytes.", bytesWritten)

			// Commit (flushes) the file contents to disk
			err = newFile.Sync()
			if err != nil {
				log.Fatal(err)
			}
		}(destination, fileOrigin, fileName)
	}

	w.Wait()
}
