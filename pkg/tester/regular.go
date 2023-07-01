package tester

import (
	"fmt"
	"github.com/ziollek/cb-perf-tester/pkg/kv/interfaces"
	"github.com/ziollek/cb-perf-tester/pkg/tester/model"
	"sync"
	"sync/atomic"
	"time"
)

type RegularGetBenchmark struct {
	KV                 interfaces.KV
	SearchForNotExists bool
	okCnt              atomic.Int32
	errCnt             atomic.Int32
}

func (benchmark *RegularGetBenchmark) DoExperiment(doc *model.Document, parallel int, repeat int) model.Report {
	var wg sync.WaitGroup
	var resultDoc model.Document
	benchmark.KV.Upsert(doc.Key, doc, 10*time.Minute)
	askKey := doc.Key
	if benchmark.SearchForNotExists {
		askKey = "not-exists"
	}
	fmt.Printf("search for key: %s\n\n", askKey)
	start := time.Now()

	for goroutine := 0; goroutine < parallel; goroutine++ {
		wg.Add(1)
		go func(myId int) {
			defer wg.Done()
			d := time.Duration(0)
			for operation := 0; operation < repeat; operation++ {
				t := time.Now()
				//err := benchmark.KV.LookupIn(doc.Key, subDocs, result...)
				err := benchmark.KV.Get(askKey, &resultDoc)
				if err != nil {
					// we do not want printing expected errors
					if !benchmark.SearchForNotExists {
						fmt.Printf("error: %s\n", err)
					}
					benchmark.errCnt.Add(1)
				} else {
					benchmark.okCnt.Add(1)
				}
				d += time.Since(t)
			}
			//fmt.Printf("goroutine %d) duration = %s\n", myId, d)
		}(goroutine)
	}
	wg.Wait()
	return model.Report{
		Errors:    int(benchmark.errCnt.Load()),
		Successes: int(benchmark.okCnt.Load()),
		Duration:  time.Since(start),
	}
}
