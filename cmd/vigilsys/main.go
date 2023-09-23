package main

import (
	"fmt"
	"time"

	"github.com/mkdemir/vigilsys/pkg/cpu"
)

func main() {
	interval := 2 * time.Second

	cpuInfoChan := cpu.StartCPUInfoCollector(interval, 1)

	for cpuInfo := range cpuInfoChan {
		fmt.Println("CPU Percentages:  ", cpuInfo.CPUPercentages)
	}
}
