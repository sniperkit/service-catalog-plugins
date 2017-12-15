// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset"
	servicecatalogv1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset/typed/servicecatalog/v1beta1"
	"k8s.io/client-go/discovery"
)

type FakeInterface struct {
	DiscoveryStub        func() discovery.DiscoveryInterface
	discoveryMutex       sync.RWMutex
	discoveryArgsForCall []struct{}
	discoveryReturns     struct {
		result1 discovery.DiscoveryInterface
	}
	discoveryReturnsOnCall map[int]struct {
		result1 discovery.DiscoveryInterface
	}
	ServicecatalogV1beta1Stub        func() servicecatalogv1beta1.ServicecatalogV1beta1Interface
	servicecatalogV1beta1Mutex       sync.RWMutex
	servicecatalogV1beta1ArgsForCall []struct{}
	servicecatalogV1beta1Returns     struct {
		result1 servicecatalogv1beta1.ServicecatalogV1beta1Interface
	}
	servicecatalogV1beta1ReturnsOnCall map[int]struct {
		result1 servicecatalogv1beta1.ServicecatalogV1beta1Interface
	}
	ServicecatalogStub        func() servicecatalogv1beta1.ServicecatalogV1beta1Interface
	servicecatalogMutex       sync.RWMutex
	servicecatalogArgsForCall []struct{}
	servicecatalogReturns     struct {
		result1 servicecatalogv1beta1.ServicecatalogV1beta1Interface
	}
	servicecatalogReturnsOnCall map[int]struct {
		result1 servicecatalogv1beta1.ServicecatalogV1beta1Interface
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInterface) Discovery() discovery.DiscoveryInterface {
	fake.discoveryMutex.Lock()
	ret, specificReturn := fake.discoveryReturnsOnCall[len(fake.discoveryArgsForCall)]
	fake.discoveryArgsForCall = append(fake.discoveryArgsForCall, struct{}{})
	fake.recordInvocation("Discovery", []interface{}{})
	fake.discoveryMutex.Unlock()
	if fake.DiscoveryStub != nil {
		return fake.DiscoveryStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.discoveryReturns.result1
}

func (fake *FakeInterface) DiscoveryCallCount() int {
	fake.discoveryMutex.RLock()
	defer fake.discoveryMutex.RUnlock()
	return len(fake.discoveryArgsForCall)
}

func (fake *FakeInterface) DiscoveryReturns(result1 discovery.DiscoveryInterface) {
	fake.DiscoveryStub = nil
	fake.discoveryReturns = struct {
		result1 discovery.DiscoveryInterface
	}{result1}
}

func (fake *FakeInterface) DiscoveryReturnsOnCall(i int, result1 discovery.DiscoveryInterface) {
	fake.DiscoveryStub = nil
	if fake.discoveryReturnsOnCall == nil {
		fake.discoveryReturnsOnCall = make(map[int]struct {
			result1 discovery.DiscoveryInterface
		})
	}
	fake.discoveryReturnsOnCall[i] = struct {
		result1 discovery.DiscoveryInterface
	}{result1}
}

func (fake *FakeInterface) ServicecatalogV1beta1() servicecatalogv1beta1.ServicecatalogV1beta1Interface {
	fake.servicecatalogV1beta1Mutex.Lock()
	ret, specificReturn := fake.servicecatalogV1beta1ReturnsOnCall[len(fake.servicecatalogV1beta1ArgsForCall)]
	fake.servicecatalogV1beta1ArgsForCall = append(fake.servicecatalogV1beta1ArgsForCall, struct{}{})
	fake.recordInvocation("ServicecatalogV1beta1", []interface{}{})
	fake.servicecatalogV1beta1Mutex.Unlock()
	if fake.ServicecatalogV1beta1Stub != nil {
		return fake.ServicecatalogV1beta1Stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.servicecatalogV1beta1Returns.result1
}

func (fake *FakeInterface) ServicecatalogV1beta1CallCount() int {
	fake.servicecatalogV1beta1Mutex.RLock()
	defer fake.servicecatalogV1beta1Mutex.RUnlock()
	return len(fake.servicecatalogV1beta1ArgsForCall)
}

func (fake *FakeInterface) ServicecatalogV1beta1Returns(result1 servicecatalogv1beta1.ServicecatalogV1beta1Interface) {
	fake.ServicecatalogV1beta1Stub = nil
	fake.servicecatalogV1beta1Returns = struct {
		result1 servicecatalogv1beta1.ServicecatalogV1beta1Interface
	}{result1}
}

func (fake *FakeInterface) ServicecatalogV1beta1ReturnsOnCall(i int, result1 servicecatalogv1beta1.ServicecatalogV1beta1Interface) {
	fake.ServicecatalogV1beta1Stub = nil
	if fake.servicecatalogV1beta1ReturnsOnCall == nil {
		fake.servicecatalogV1beta1ReturnsOnCall = make(map[int]struct {
			result1 servicecatalogv1beta1.ServicecatalogV1beta1Interface
		})
	}
	fake.servicecatalogV1beta1ReturnsOnCall[i] = struct {
		result1 servicecatalogv1beta1.ServicecatalogV1beta1Interface
	}{result1}
}

func (fake *FakeInterface) Servicecatalog() servicecatalogv1beta1.ServicecatalogV1beta1Interface {
	fake.servicecatalogMutex.Lock()
	ret, specificReturn := fake.servicecatalogReturnsOnCall[len(fake.servicecatalogArgsForCall)]
	fake.servicecatalogArgsForCall = append(fake.servicecatalogArgsForCall, struct{}{})
	fake.recordInvocation("Servicecatalog", []interface{}{})
	fake.servicecatalogMutex.Unlock()
	if fake.ServicecatalogStub != nil {
		return fake.ServicecatalogStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.servicecatalogReturns.result1
}

func (fake *FakeInterface) ServicecatalogCallCount() int {
	fake.servicecatalogMutex.RLock()
	defer fake.servicecatalogMutex.RUnlock()
	return len(fake.servicecatalogArgsForCall)
}

func (fake *FakeInterface) ServicecatalogReturns(result1 servicecatalogv1beta1.ServicecatalogV1beta1Interface) {
	fake.ServicecatalogStub = nil
	fake.servicecatalogReturns = struct {
		result1 servicecatalogv1beta1.ServicecatalogV1beta1Interface
	}{result1}
}

func (fake *FakeInterface) ServicecatalogReturnsOnCall(i int, result1 servicecatalogv1beta1.ServicecatalogV1beta1Interface) {
	fake.ServicecatalogStub = nil
	if fake.servicecatalogReturnsOnCall == nil {
		fake.servicecatalogReturnsOnCall = make(map[int]struct {
			result1 servicecatalogv1beta1.ServicecatalogV1beta1Interface
		})
	}
	fake.servicecatalogReturnsOnCall[i] = struct {
		result1 servicecatalogv1beta1.ServicecatalogV1beta1Interface
	}{result1}
}

func (fake *FakeInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.discoveryMutex.RLock()
	defer fake.discoveryMutex.RUnlock()
	fake.servicecatalogV1beta1Mutex.RLock()
	defer fake.servicecatalogV1beta1Mutex.RUnlock()
	fake.servicecatalogMutex.RLock()
	defer fake.servicecatalogMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInterface) recordInvocation(key string, args []interface{}) {
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

var _ clientset.Interface = new(FakeInterface)