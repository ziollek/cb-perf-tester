package kv

import (
	"github.com/couchbase/gocb/v2"
	"github.com/ziollek/cb-perf-tester/pkg/config"
	"github.com/ziollek/cb-perf-tester/pkg/kv/interfaces"
	"time"
)

const GenericTimeout = 15 * time.Second

type Couchbase struct {
	collection *gocb.Collection
	uri        string
}

func buildConnectionOptions(c *config.Couchbase) gocb.ClusterOptions {
	options := gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: c.User,
			Password: c.Password,
		},
	}
	// TODO: make timeouts configurable
	options.TimeoutsConfig.ConnectTimeout = GenericTimeout
	options.TimeoutsConfig.KVTimeout = GenericTimeout
	return options
}

func NewKV(c *config.Couchbase) (interfaces.KV, error) {
	cluster, err := gocb.Connect(c.URI, buildConnectionOptions(c))
	if err != nil {
		return nil, err
	}
	bucket := cluster.Bucket(c.Bucket)
	err = bucket.WaitUntilReady(GenericTimeout, nil)
	if err != nil {
		return nil, err
	}
	return &Couchbase{collection: bucket.DefaultCollection(), uri: c.URI}, nil
}

func (c *Couchbase) Upsert(key string, value interface{}, expiry time.Duration) error {
	_, err := c.collection.Upsert(key, value, &gocb.UpsertOptions{Expiry: expiry})
	return err
}

func (c *Couchbase) Get(key string, value interface{}) error {
	_, err := c.collection.Get(key, nil)
	if err != nil {
		return err
	}
	return nil
	//err = result.Content(value)
	return err
}

func (c *Couchbase) LookupIn(key string, fields []string, values ...interface{}) error {
	ops := []gocb.LookupInSpec{}
	for _, field := range fields {
		ops = append(ops, gocb.GetSpec(field, &gocb.GetSpecOptions{}))
	}
	result, err := c.collection.LookupIn(key, ops, &gocb.LookupInOptions{})
	if err != nil {
		return err
	}
	for i, value := range values {
		err = result.ContentAt(uint(i), value)
		if err != nil {
			break
		}
	}
	return err
}
