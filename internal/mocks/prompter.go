// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	big "math/big"

	ids "github.com/ava-labs/avalanchego/ids"
	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	models "github.com/ava-labs/avalanche-cli/pkg/models"

	prompts "github.com/ava-labs/avalanche-cli/pkg/prompts"

	time "time"

	url "net/url"
)

// Prompter is an autogenerated mock type for the Prompter type
type Prompter struct {
	mock.Mock
}

// CaptureAddress provides a mock function with given fields: promptStr
func (_m *Prompter) CaptureAddress(promptStr string) (common.Address, error) {
	ret := _m.Called(promptStr)

	if len(ret) == 0 {
		panic("no return value specified for CaptureAddress")
	}

	var r0 common.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (common.Address, error)); ok {
		return rf(promptStr)
	}
	if rf, ok := ret.Get(0).(func(string) common.Address); ok {
		r0 = rf(promptStr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(promptStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureAddresses provides a mock function with given fields: promptStr
func (_m *Prompter) CaptureAddresses(promptStr string) ([]common.Address, error) {
	ret := _m.Called(promptStr)

	if len(ret) == 0 {
		panic("no return value specified for CaptureList")
	}

	var r0 []common.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]common.Address, error)); ok {
		return rf(promptStr)
	}
	if rf, ok := ret.Get(0).(func(string) []common.Address); ok {
		r0 = rf(promptStr)
	} else {
		r0 = ret.Get(0).([]common.Address)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(promptStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureDate provides a mock function with given fields: promptStr
func (_m *Prompter) CaptureDate(promptStr string) (time.Time, error) {
	ret := _m.Called(promptStr)

	if len(ret) == 0 {
		panic("no return value specified for CaptureDate")
	}

	var r0 time.Time
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (time.Time, error)); ok {
		return rf(promptStr)
	}
	if rf, ok := ret.Get(0).(func(string) time.Time); ok {
		r0 = rf(promptStr)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(promptStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureEmail provides a mock function with given fields: promptStr
func (_m *Prompter) CaptureEmail(promptStr string) (string, error) {
	ret := _m.Called(promptStr)

	if len(ret) == 0 {
		panic("no return value specified for CaptureEmail")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(promptStr)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(promptStr)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(promptStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureExistingFilepath provides a mock function with given fields: promptStr
func (_m *Prompter) CaptureExistingFilepath(promptStr string) (string, error) {
	ret := _m.Called(promptStr)

	if len(ret) == 0 {
		panic("no return value specified for CaptureExistingFilepath")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(promptStr)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(promptStr)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(promptStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureFloat provides a mock function with given fields: promptStr, validator
func (_m *Prompter) CaptureFloat(promptStr string, validator func(float64) error) (float64, error) {
	ret := _m.Called(promptStr, validator)

	if len(ret) == 0 {
		panic("no return value specified for CaptureFloat")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(string, func(float64) error) (float64, error)); ok {
		return rf(promptStr, validator)
	}
	if rf, ok := ret.Get(0).(func(string, func(float64) error) float64); ok {
		r0 = rf(promptStr, validator)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(string, func(float64) error) error); ok {
		r1 = rf(promptStr, validator)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureFujiDuration provides a mock function with given fields: promptStr
func (_m *Prompter) CaptureFujiDuration(promptStr string) (time.Duration, error) {
	ret := _m.Called(promptStr)

	if len(ret) == 0 {
		panic("no return value specified for CaptureFujiDuration")
	}

	var r0 time.Duration
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (time.Duration, error)); ok {
		return rf(promptStr)
	}
	if rf, ok := ret.Get(0).(func(string) time.Duration); ok {
		r0 = rf(promptStr)
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(promptStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureFutureDate provides a mock function with given fields: promptStr, minDate
func (_m *Prompter) CaptureFutureDate(promptStr string, minDate time.Time) (time.Time, error) {
	ret := _m.Called(promptStr, minDate)

	if len(ret) == 0 {
		panic("no return value specified for CaptureFutureDate")
	}

	var r0 time.Time
	var r1 error
	if rf, ok := ret.Get(0).(func(string, time.Time) (time.Time, error)); ok {
		return rf(promptStr, minDate)
	}
	if rf, ok := ret.Get(0).(func(string, time.Time) time.Time); ok {
		r0 = rf(promptStr, minDate)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	if rf, ok := ret.Get(1).(func(string, time.Time) error); ok {
		r1 = rf(promptStr, minDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureGitURL provides a mock function with given fields: promptStr
func (_m *Prompter) CaptureGitURL(promptStr string) (*url.URL, error) {
	ret := _m.Called(promptStr)

	if len(ret) == 0 {
		panic("no return value specified for CaptureGitURL")
	}

	var r0 *url.URL
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*url.URL, error)); ok {
		return rf(promptStr)
	}
	if rf, ok := ret.Get(0).(func(string) *url.URL); ok {
		r0 = rf(promptStr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*url.URL)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(promptStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureID provides a mock function with given fields: promptStr
func (_m *Prompter) CaptureID(promptStr string) (ids.ID, error) {
	ret := _m.Called(promptStr)

	if len(ret) == 0 {
		panic("no return value specified for CaptureID")
	}

	var r0 ids.ID
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (ids.ID, error)); ok {
		return rf(promptStr)
	}
	if rf, ok := ret.Get(0).(func(string) ids.ID); ok {
		r0 = rf(promptStr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ids.ID)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(promptStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureIndex provides a mock function with given fields: promptStr, options
func (_m *Prompter) CaptureIndex(promptStr string, options []interface{}) (int, error) {
	ret := _m.Called(promptStr, options)

	if len(ret) == 0 {
		panic("no return value specified for CaptureIndex")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []interface{}) (int, error)); ok {
		return rf(promptStr, options)
	}
	if rf, ok := ret.Get(0).(func(string, []interface{}) int); ok {
		r0 = rf(promptStr, options)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string, []interface{}) error); ok {
		r1 = rf(promptStr, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureInt provides a mock function with given fields: promptStr
func (_m *Prompter) CaptureInt(promptStr string) (int, error) {
	ret := _m.Called(promptStr)

	if len(ret) == 0 {
		panic("no return value specified for CaptureInt")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int, error)); ok {
		return rf(promptStr)
	}
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(promptStr)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(promptStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureList provides a mock function with given fields: promptStr, options
func (_m *Prompter) CaptureList(promptStr string, options []string) (string, error) {
	ret := _m.Called(promptStr, options)

	if len(ret) == 0 {
		panic("no return value specified for CaptureList")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []string) (string, error)); ok {
		return rf(promptStr, options)
	}
	if rf, ok := ret.Get(0).(func(string, []string) string); ok {
		r0 = rf(promptStr, options)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(promptStr, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureListWithSize provides a mock function with given fields: promptStr, options, size
func (_m *Prompter) CaptureListWithSize(promptStr string, options []string, size int) (string, error) {
	ret := _m.Called(promptStr, options, size)

	if len(ret) == 0 {
		panic("no return value specified for CaptureListWithSize")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []string, int) (string, error)); ok {
		return rf(promptStr, options, size)
	}
	if rf, ok := ret.Get(0).(func(string, []string, int) string); ok {
		r0 = rf(promptStr, options, size)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, []string, int) error); ok {
		r1 = rf(promptStr, options, size)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureMainnetDuration provides a mock function with given fields: promptStr
func (_m *Prompter) CaptureMainnetDuration(promptStr string) (time.Duration, error) {
	ret := _m.Called(promptStr)

	if len(ret) == 0 {
		panic("no return value specified for CaptureMainnetDuration")
	}

	var r0 time.Duration
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (time.Duration, error)); ok {
		return rf(promptStr)
	}
	if rf, ok := ret.Get(0).(func(string) time.Duration); ok {
		r0 = rf(promptStr)
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(promptStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureNewFilepath provides a mock function with given fields: promptStr
func (_m *Prompter) CaptureNewFilepath(promptStr string) (string, error) {
	ret := _m.Called(promptStr)

	if len(ret) == 0 {
		panic("no return value specified for CaptureNewFilepath")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(promptStr)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(promptStr)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(promptStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureNoYes provides a mock function with given fields: promptStr
func (_m *Prompter) CaptureNoYes(promptStr string) (bool, error) {
	ret := _m.Called(promptStr)

	if len(ret) == 0 {
		panic("no return value specified for CaptureNoYes")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(promptStr)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(promptStr)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(promptStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureNodeID provides a mock function with given fields: promptStr
func (_m *Prompter) CaptureNodeID(promptStr string) (ids.NodeID, error) {
	ret := _m.Called(promptStr)

	if len(ret) == 0 {
		panic("no return value specified for CaptureNodeID")
	}

	var r0 ids.NodeID
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (ids.NodeID, error)); ok {
		return rf(promptStr)
	}
	if rf, ok := ret.Get(0).(func(string) ids.NodeID); ok {
		r0 = rf(promptStr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ids.NodeID)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(promptStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CapturePChainAddress provides a mock function with given fields: promptStr, network
func (_m *Prompter) CapturePChainAddress(promptStr string, network models.Network) (string, error) {
	ret := _m.Called(promptStr, network)

	if len(ret) == 0 {
		panic("no return value specified for CapturePChainAddress")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, models.Network) (string, error)); ok {
		return rf(promptStr, network)
	}
	if rf, ok := ret.Get(0).(func(string, models.Network) string); ok {
		r0 = rf(promptStr, network)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, models.Network) error); ok {
		r1 = rf(promptStr, network)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CapturePositiveBigInt provides a mock function with given fields: promptStr
func (_m *Prompter) CapturePositiveBigInt(promptStr string) (*big.Int, error) {
	ret := _m.Called(promptStr)

	if len(ret) == 0 {
		panic("no return value specified for CapturePositiveBigInt")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*big.Int, error)); ok {
		return rf(promptStr)
	}
	if rf, ok := ret.Get(0).(func(string) *big.Int); ok {
		r0 = rf(promptStr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(promptStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CapturePositiveInt provides a mock function with given fields: promptStr, comparators
func (_m *Prompter) CapturePositiveInt(promptStr string, comparators []prompts.Comparator) (int, error) {
	ret := _m.Called(promptStr, comparators)

	if len(ret) == 0 {
		panic("no return value specified for CapturePositiveInt")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []prompts.Comparator) (int, error)); ok {
		return rf(promptStr, comparators)
	}
	if rf, ok := ret.Get(0).(func(string, []prompts.Comparator) int); ok {
		r0 = rf(promptStr, comparators)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string, []prompts.Comparator) error); ok {
		r1 = rf(promptStr, comparators)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureRepoBranch provides a mock function with given fields: promptStr, repo
func (_m *Prompter) CaptureRepoBranch(promptStr string, repo string) (string, error) {
	ret := _m.Called(promptStr, repo)

	if len(ret) == 0 {
		panic("no return value specified for CaptureRepoBranch")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(promptStr, repo)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(promptStr, repo)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(promptStr, repo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureRepoFile provides a mock function with given fields: promptStr, repo, branch
func (_m *Prompter) CaptureRepoFile(promptStr string, repo string, branch string) (string, error) {
	ret := _m.Called(promptStr, repo, branch)

	if len(ret) == 0 {
		panic("no return value specified for CaptureRepoFile")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (string, error)); ok {
		return rf(promptStr, repo, branch)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) string); ok {
		r0 = rf(promptStr, repo, branch)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(promptStr, repo, branch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureString provides a mock function with given fields: promptStr
func (_m *Prompter) CaptureString(promptStr string) (string, error) {
	ret := _m.Called(promptStr)

	if len(ret) == 0 {
		panic("no return value specified for CaptureString")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(promptStr)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(promptStr)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(promptStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureStringAllowEmpty provides a mock function with given fields: promptStr
func (_m *Prompter) CaptureStringAllowEmpty(promptStr string) (string, error) {
	ret := _m.Called(promptStr)

	if len(ret) == 0 {
		panic("no return value specified for CaptureStringAllowEmpty")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(promptStr)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(promptStr)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(promptStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureURL provides a mock function with given fields: promptStr, validateConnection
func (_m *Prompter) CaptureURL(promptStr string, validateConnection bool) (string, error) {
	ret := _m.Called(promptStr, validateConnection)

	if len(ret) == 0 {
		panic("no return value specified for CaptureURL")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, bool) (string, error)); ok {
		return rf(promptStr, validateConnection)
	}
	if rf, ok := ret.Get(0).(func(string, bool) string); ok {
		r0 = rf(promptStr, validateConnection)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(promptStr, validateConnection)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureUint32 provides a mock function with given fields: promptStr
func (_m *Prompter) CaptureUint32(promptStr string) (uint32, error) {
	ret := _m.Called(promptStr)

	if len(ret) == 0 {
		panic("no return value specified for CaptureUint32")
	}

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (uint32, error)); ok {
		return rf(promptStr)
	}
	if rf, ok := ret.Get(0).(func(string) uint32); ok {
		r0 = rf(promptStr)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(promptStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureUint64 provides a mock function with given fields: promptStr
func (_m *Prompter) CaptureUint64(promptStr string) (uint64, error) {
	ret := _m.Called(promptStr)

	if len(ret) == 0 {
		panic("no return value specified for CaptureUint64")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (uint64, error)); ok {
		return rf(promptStr)
	}
	if rf, ok := ret.Get(0).(func(string) uint64); ok {
		r0 = rf(promptStr)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(promptStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureUint64Compare provides a mock function with given fields: promptStr, comparators
func (_m *Prompter) CaptureUint64Compare(promptStr string, comparators []prompts.Comparator) (uint64, error) {
	ret := _m.Called(promptStr, comparators)

	if len(ret) == 0 {
		panic("no return value specified for CaptureUint64Compare")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []prompts.Comparator) (uint64, error)); ok {
		return rf(promptStr, comparators)
	}
	if rf, ok := ret.Get(0).(func(string, []prompts.Comparator) uint64); ok {
		r0 = rf(promptStr, comparators)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(string, []prompts.Comparator) error); ok {
		r1 = rf(promptStr, comparators)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureValidatedString provides a mock function with given fields: promptStr, validator
func (_m *Prompter) CaptureValidatedString(promptStr string, validator func(string) error) (string, error) {
	ret := _m.Called(promptStr, validator)

	if len(ret) == 0 {
		panic("no return value specified for CaptureValidatedString")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, func(string) error) (string, error)); ok {
		return rf(promptStr, validator)
	}
	if rf, ok := ret.Get(0).(func(string, func(string) error) string); ok {
		r0 = rf(promptStr, validator)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, func(string) error) error); ok {
		r1 = rf(promptStr, validator)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureVersion provides a mock function with given fields: promptStr
func (_m *Prompter) CaptureVersion(promptStr string) (string, error) {
	ret := _m.Called(promptStr)

	if len(ret) == 0 {
		panic("no return value specified for CaptureVersion")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(promptStr)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(promptStr)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(promptStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureWeight provides a mock function with given fields: promptStr
func (_m *Prompter) CaptureWeight(promptStr string) (uint64, error) {
	ret := _m.Called(promptStr)

	if len(ret) == 0 {
		panic("no return value specified for CaptureWeight")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (uint64, error)); ok {
		return rf(promptStr)
	}
	if rf, ok := ret.Get(0).(func(string) uint64); ok {
		r0 = rf(promptStr)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(promptStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureXChainAddress provides a mock function with given fields: promptStr, network
func (_m *Prompter) CaptureXChainAddress(promptStr string, network models.Network) (string, error) {
	ret := _m.Called(promptStr, network)

	if len(ret) == 0 {
		panic("no return value specified for CaptureXChainAddress")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, models.Network) (string, error)); ok {
		return rf(promptStr, network)
	}
	if rf, ok := ret.Get(0).(func(string, models.Network) string); ok {
		r0 = rf(promptStr, network)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, models.Network) error); ok {
		r1 = rf(promptStr, network)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptureYesNo provides a mock function with given fields: promptStr
func (_m *Prompter) CaptureYesNo(promptStr string) (bool, error) {
	ret := _m.Called(promptStr)

	if len(ret) == 0 {
		panic("no return value specified for CaptureYesNo")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(promptStr)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(promptStr)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(promptStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChooseKeyOrLedger provides a mock function with given fields: goal
func (_m *Prompter) ChooseKeyOrLedger(goal string) (bool, error) {
	ret := _m.Called(goal)

	if len(ret) == 0 {
		panic("no return value specified for ChooseKeyOrLedger")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(goal)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(goal)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(goal)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPrompter creates a new instance of Prompter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPrompter(t interface {
	mock.TestingT
	Cleanup(func())
}) *Prompter {
	mock := &Prompter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
