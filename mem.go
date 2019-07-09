package mem

import (
	"fmt"
	"os"
	"runtime"
	"text/tabwriter"
)

// PrintMem ...
func PrintMem() {
	var ms runtime.MemStats
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "Alloc\t", "TotalAlloc\t", "Sys\t", "Mallocs\t", "Frees\t", "LiveObjects\t", "PauseTotalNs\t", "NumGC\t", "NumGoroutine\t")
	fmt.Fprintln(w, "-----\t", "----------\t", "---\t", "-------\t", "-----\t", "-----------\t", "------------\t", "-----\t", "------------\t")

	runtime.ReadMemStats(&ms)

	fmt.Fprintln(w, BtoMB(ms.Alloc),
		fmt.Sprintf("%v MiB\t", BtoMB(ms.TotalAlloc)),
		fmt.Sprintf("%v MiB\t", BtoMB(ms.Sys)),
		fmt.Sprintf("%v MiB\t", BtoMB(ms.Mallocs)),
		fmt.Sprintf("%v MiB\t", BtoMB(ms.Frees)),
		fmt.Sprintf("%v MiB\t", BtoMB(ms.Mallocs-ms.Frees)),
		fmt.Sprintf("%v\t", NStoMS(ms.PauseTotalNs)),
		fmt.Sprintf("%v\t", ms.NumGC),
		fmt.Sprintf("%v\t", runtime.NumGoroutine()))

	w.Flush()
}

// BtoMB ...
func BtoMB(b uint64) uint64 {
	return b / 1024 / 1024
}

// NStoMS ...
func NStoMS(ns uint64) uint64 {
	return ns * (1 / 1000000)
}
