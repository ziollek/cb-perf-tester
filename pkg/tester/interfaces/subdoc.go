package interfaces

import "github.com/ziollek/cb-perf-tester/pkg/tester/model"

type Benchmark interface {
	DoExperiment(doc *model.Document, parallel int, repeat int) model.Report
}
