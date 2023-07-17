package utils

import (
	"fmt"
	"time"
)

func GetDate() {
	date := time.Now()
	fmt.Println(date.Format("02/Jan/2006 15:04:05 "))
}