package cpu

import (
	"log"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// CPUInfo struct is used to store CPU information.
type CPUInfo struct {
	CPUPercentages []float64
}

// StartCPUInfoCollector is used to collect CPU statistics at a specified time interval.
func StartCPUInfoCollector(interval time.Duration, numCPU int) <-chan CPUInfo {
	infoChan := make(chan CPUInfo)

	go func() {
		for {
			percentages, err := cpu.Percent(interval, true)
			if err != nil {
				log.Fatal(err)
			}

			cpuInfo := CPUInfo{
				CPUPercentages: percentages[:numCPU],
			}

			infoChan <- cpuInfo

			time.Sleep(interval)
		}
	}()

	return infoChan
}
