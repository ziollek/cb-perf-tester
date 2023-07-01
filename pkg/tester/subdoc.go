package tester

import (
	"fmt"
	"github.com/ziollek/cb-perf-tester/pkg/kv/interfaces"
	"github.com/ziollek/cb-perf-tester/pkg/tester/model"
	"sync"
	"sync/atomic"
	"time"
)

type ConfigurableSubDocBenchmark struct {
	SubDocLevel     model.Difficulty
	SubDocSearchNum int
	KV              interfaces.KV
	okCnt           atomic.Int32
	errCnt          atomic.Int32
}

func (tester *ConfigurableSubDocBenchmark) DoExperiment(doc *model.Document, parallel int, repeat int) model.Report {
	var wg sync.WaitGroup
	tester.KV.Upsert(doc.Key, doc, 10*time.Minute)
	subDocs := doc.GetSubKeysByDifficulty(tester.SubDocLevel, tester.SubDocSearchNum)
	fmt.Printf("search for subkeys [level=%s]: %s\n\n", tester.SubDocLevel.ToString(), subDocs)
	start := time.Now()
	for goroutine := 0; goroutine < parallel; goroutine++ {
		wg.Add(1)
		go func(myId int) {
			defer wg.Done()
			result := prepareResultPointers(tester.SubDocSearchNum)
			d := time.Duration(0)
			for operation := 0; operation < repeat; operation++ {
				t := time.Now()
				err := tester.KV.LookupIn(doc.Key, subDocs, result...)
				if err != nil {
					// we do not want printing expected errors
					if tester.SubDocLevel != model.Impossible {
						fmt.Printf("error: %s\n", err)
					}
					tester.errCnt.Add(1)
				} else {
					tester.okCnt.Add(1)
				}
				d += time.Since(t)
			}
			//fmt.Printf("goroutine %d) duration = %s\n", myId, d)
		}(goroutine)
	}
	wg.Wait()
	return model.Report{
		Errors:    int(tester.errCnt.Load()),
		Successes: int(tester.okCnt.Load()),
		Duration:  time.Since(start),
	}
}

func prepareResultPointers(size int) (result []interface{}) {
	data := make([]string, size)
	for i := range data {
		result = append(result, &data[i])
	}
	return
}
