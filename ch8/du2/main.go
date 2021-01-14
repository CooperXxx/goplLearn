package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose message")

//du2 增加了对扫描的实时显示
func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(50 * time.Millisecond)
	}

	var nfiles, nbytes int64
	//for size := range fileSizes {
	//	nfiles++
	//	nbytes+=size
	//}
loop:
	for {
		select {
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles int64, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir(root string, sizes chan int64) {
	//time.Sleep(time.Millisecond*10)
	for _, entry := range directs(root) {
		if entry.IsDir() {
			subDir := filepath.Join(root, entry.Name())
			walkDir(subDir, sizes)
		} else {
			sizes <- entry.Size()
		}
	}
}

func directs(dir string) []os.FileInfo {
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1:%v\n", err)
	}
	return infos
}
