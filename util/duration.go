package util

import (
	"fmt"
	"time"
)

func Duration(start time.Time) {
	fmt.Printf(" in %fs\n", time.Since(start).Seconds())
}
