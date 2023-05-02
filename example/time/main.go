package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2020, 1, 1, 10, 0, 0, 0, time.UTC)

	formatted := t.Format(time.RFC3339)
	formatted = t.Format("2006-01-02 15:04:05")

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	formatted = t.In(jst).Format("2006-01-02 15:04:05")

	fmt.Println(formatted)
}
