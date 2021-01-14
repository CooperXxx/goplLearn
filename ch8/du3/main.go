package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose message")
var sema = make(chan struct{}, 20)

//du3并发执行walkDir,并且使用信号量控制并发数
func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	//go func() {
	//	for _, root := range roots {
	//		walkDir(root,fileSizes)
	//	}
	//	close(fileSizes)
	//}()

	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}

	go func() {
		n.Wait()
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

func walkDir(root string, n *sync.WaitGroup, sizes chan int64) {
	//time.Sleep(time.Millisecond*10)

	defer n.Done()
	for _, entry := range directs(root) {
		if entry.IsDir() {
			subDir := filepath.Join(root, entry.Name())
			n.Add(1)
			go walkDir(subDir, n, sizes)
		} else {
			sizes <- entry.Size()
		}
	}
}

func directs(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1:%v\n", err)
	}
	return infos
}
