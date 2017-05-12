package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.Parse("2006-01-02 15:04:05.000", "2016-11-22 20:14:15.833")
	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println(t.Format(time.RFC3339))
}
