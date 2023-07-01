package infra

import (
	"github.com/ziollek/cb-perf-tester/pkg/config"
	"github.com/ziollek/cb-perf-tester/pkg/kv"
	kvInterfaces "github.com/ziollek/cb-perf-tester/pkg/kv/interfaces"
	"github.com/ziollek/cb-perf-tester/pkg/tester"
	"github.com/ziollek/cb-perf-tester/pkg/tester/interfaces"
	"github.com/ziollek/cb-perf-tester/pkg/tester/model"
)

func BuildSubDocBenchmark(level model.Difficulty, searchNum int) (interfaces.Benchmark, error) {
	kv, err := BuildKV()
	if err != nil {
		return nil, err
	}
	return &tester.ConfigurableSubDocBenchmark{SubDocLevel: level, SubDocSearchNum: searchNum, KV: kv}, nil
}

func BuildRegularBenchmark(searchNotExistent bool) (interfaces.Benchmark, error) {
	kv, err := BuildKV()
	if err != nil {
		return nil, err
	}
	return &tester.RegularGetBenchmark{SearchForNotExists: searchNotExistent, KV: kv}, nil
}

func BuildKV() (kvInterfaces.KV, error) {
	c, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	return kv.NewKV(c.Couchbase)
}
