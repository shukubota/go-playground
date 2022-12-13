package main

import (
	"archive/zip"
	"fmt"
	"log"
)

func main() {
	r, err := zip.OpenReader("/Users/shu.kubota/Desktop/upload/ken_all.zip")
	if err != nil {
		log.Fatalf("%+v", err)
	}
	fmt.Println(r)
	defer r.Close()

	for _, f := range r.File {
		fmt.Println(f.Name)
		//if f.Name != csvFilename {
		//	continue
		//}

		//outFile, err := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		//if err != nil {
		//	return err
		//}
		//
		//rc, err := f.Open()
		//if err != nil {
		//	return err
		//}
		//_, err = io.Copy(outFile, rc)
		//
		//outFile.Close()
		//rc.Close()
		//
		//return err
	}
}
