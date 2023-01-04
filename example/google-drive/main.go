package main

import (
	"context"
	"fmt"
	"google.golang.org/api/drive/v3"
	"log"
)

func main() {
	ctx := context.Background()
	s, err := drive.NewService(ctx)
	if err != nil {
		log.Fatalf("failed to instantiate google drive service: %+v", err)
	}

	fl := s.Files.List().Fields("files(id, name)").Context(ctx)

	fmt.Printf("filelist: %+v\n", fl)

	r, err := fl.Do()

	if err != nil {
		log.Fatalf("error do")
	}

	fmt.Println(r.Files)
	fmt.Println("===========================b")
	for _, f := range r.Files {
		fmt.Println(f.Name, f.Id)
	}

	//r, err := s.Files.List().PageSize(1000).
	//	Fields("files(id, name)").
	//	Context(ctx).Do()
	//if err != nil {
	//	log.Fatalf("Unable to retrieve files: %v", err)
	//}

	fmt.Println("===========================aa")
	fmt.Println(s)
}
