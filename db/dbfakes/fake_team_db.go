// This file was generated by counterfeiter
package dbfakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
)

type FakeTeamDB struct {
	GetTeamStub        func() (db.SavedTeam, bool, error)
	getTeamMutex       sync.RWMutex
	getTeamArgsForCall []struct{}
	getTeamReturns     struct {
		result1 db.SavedTeam
		result2 bool
		result3 error
	}
	getTeamReturnsOnCall map[int]struct {
		result1 db.SavedTeam
		result2 bool
		result3 error
	}
	GetConfigStub        func(pipelineName string) (atc.Config, atc.RawConfig, db.ConfigVersion, error)
	getConfigMutex       sync.RWMutex
	getConfigArgsForCall []struct {
		pipelineName string
	}
	getConfigReturns struct {
		result1 atc.Config
		result2 atc.RawConfig
		result3 db.ConfigVersion
		result4 error
	}
	getConfigReturnsOnCall map[int]struct {
		result1 atc.Config
		result2 atc.RawConfig
		result3 db.ConfigVersion
		result4 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTeamDB) GetTeam() (db.SavedTeam, bool, error) {
	fake.getTeamMutex.Lock()
	ret, specificReturn := fake.getTeamReturnsOnCall[len(fake.getTeamArgsForCall)]
	fake.getTeamArgsForCall = append(fake.getTeamArgsForCall, struct{}{})
	fake.recordInvocation("GetTeam", []interface{}{})
	fake.getTeamMutex.Unlock()
	if fake.GetTeamStub != nil {
		return fake.GetTeamStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getTeamReturns.result1, fake.getTeamReturns.result2, fake.getTeamReturns.result3
}

func (fake *FakeTeamDB) GetTeamCallCount() int {
	fake.getTeamMutex.RLock()
	defer fake.getTeamMutex.RUnlock()
	return len(fake.getTeamArgsForCall)
}

func (fake *FakeTeamDB) GetTeamReturns(result1 db.SavedTeam, result2 bool, result3 error) {
	fake.GetTeamStub = nil
	fake.getTeamReturns = struct {
		result1 db.SavedTeam
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeamDB) GetTeamReturnsOnCall(i int, result1 db.SavedTeam, result2 bool, result3 error) {
	fake.GetTeamStub = nil
	if fake.getTeamReturnsOnCall == nil {
		fake.getTeamReturnsOnCall = make(map[int]struct {
			result1 db.SavedTeam
			result2 bool
			result3 error
		})
	}
	fake.getTeamReturnsOnCall[i] = struct {
		result1 db.SavedTeam
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeamDB) GetConfig(pipelineName string) (atc.Config, atc.RawConfig, db.ConfigVersion, error) {
	fake.getConfigMutex.Lock()
	ret, specificReturn := fake.getConfigReturnsOnCall[len(fake.getConfigArgsForCall)]
	fake.getConfigArgsForCall = append(fake.getConfigArgsForCall, struct {
		pipelineName string
	}{pipelineName})
	fake.recordInvocation("GetConfig", []interface{}{pipelineName})
	fake.getConfigMutex.Unlock()
	if fake.GetConfigStub != nil {
		return fake.GetConfigStub(pipelineName)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	return fake.getConfigReturns.result1, fake.getConfigReturns.result2, fake.getConfigReturns.result3, fake.getConfigReturns.result4
}

func (fake *FakeTeamDB) GetConfigCallCount() int {
	fake.getConfigMutex.RLock()
	defer fake.getConfigMutex.RUnlock()
	return len(fake.getConfigArgsForCall)
}

func (fake *FakeTeamDB) GetConfigArgsForCall(i int) string {
	fake.getConfigMutex.RLock()
	defer fake.getConfigMutex.RUnlock()
	return fake.getConfigArgsForCall[i].pipelineName
}

func (fake *FakeTeamDB) GetConfigReturns(result1 atc.Config, result2 atc.RawConfig, result3 db.ConfigVersion, result4 error) {
	fake.GetConfigStub = nil
	fake.getConfigReturns = struct {
		result1 atc.Config
		result2 atc.RawConfig
		result3 db.ConfigVersion
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeTeamDB) GetConfigReturnsOnCall(i int, result1 atc.Config, result2 atc.RawConfig, result3 db.ConfigVersion, result4 error) {
	fake.GetConfigStub = nil
	if fake.getConfigReturnsOnCall == nil {
		fake.getConfigReturnsOnCall = make(map[int]struct {
			result1 atc.Config
			result2 atc.RawConfig
			result3 db.ConfigVersion
			result4 error
		})
	}
	fake.getConfigReturnsOnCall[i] = struct {
		result1 atc.Config
		result2 atc.RawConfig
		result3 db.ConfigVersion
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeTeamDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getTeamMutex.RLock()
	defer fake.getTeamMutex.RUnlock()
	fake.getConfigMutex.RLock()
	defer fake.getConfigMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTeamDB) recordInvocation(key string, args []interface{}) {
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

var _ db.TeamDB = new(FakeTeamDB)
