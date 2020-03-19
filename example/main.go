package main

import (
	"fmt"
	"sync"
	"time"
	"github.com/hhq163/snow_flake/gen_id"
)

func main() {
	node, err := gen_id.New(13)
	if err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan int64, 10000)
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			id := node.GenId()
			ch <- id
		}()
	}
	wg.Wait()
	close(ch)

	ids := make(map[int64]bool)
	for id := range ch {
		if _, ok := ids[id]; ok {
			fmt.Println("ID is not unique!")
			return
		}
		ids[id] = true
		timestamp, nodeid, idNumber, err := node.PraseId(id)
		if err != nil {
			fmt.Println("node.GetAttrs error id=", id)
			continue
		}
		tm := time.Unix(timestamp/1000, 0)

		fmt.Println("id=", id, "nodeid=", nodeid, "tm=", tm.Format("2006-01-02 15:04:05"), "idNumber=", idNumber)
	}
	
	fmt.Println("the length of ids is ", len(ids))
}
