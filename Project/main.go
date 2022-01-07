package main

import (
	"time"

	"github.com/absoran/goproject/internal"
)

func main() {
	defer internal.CalculationTime(time.Now(), "")
	data, err := internal.GetInputData()
	internal.CheckError(err)
	internal.ProcessFlag(data)
}
