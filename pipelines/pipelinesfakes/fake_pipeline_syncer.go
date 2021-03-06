// This file was generated by counterfeiter
package pipelinesfakes

import (
	"sync"

	"github.com/concourse/atc/pipelines"
)

type FakePipelineSyncer struct {
	SyncStub         func()
	syncMutex        sync.RWMutex
	syncArgsForCall  []struct{}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePipelineSyncer) Sync() {
	fake.syncMutex.Lock()
	fake.syncArgsForCall = append(fake.syncArgsForCall, struct{}{})
	fake.recordInvocation("Sync", []interface{}{})
	fake.syncMutex.Unlock()
	if fake.SyncStub != nil {
		fake.SyncStub()
	}
}

func (fake *FakePipelineSyncer) SyncCallCount() int {
	fake.syncMutex.RLock()
	defer fake.syncMutex.RUnlock()
	return len(fake.syncArgsForCall)
}

func (fake *FakePipelineSyncer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.syncMutex.RLock()
	defer fake.syncMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePipelineSyncer) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ pipelines.PipelineSyncer = new(FakePipelineSyncer)
