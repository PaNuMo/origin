// This file was automatically generated by lister-gen with arguments: --input-dirs=[github.com/openshift/origin/pkg/authorization/api,github.com/openshift/origin/pkg/authorization/api/v1,github.com/openshift/origin/pkg/build/api,github.com/openshift/origin/pkg/build/api/v1,github.com/openshift/origin/pkg/deploy/api,github.com/openshift/origin/pkg/deploy/api/v1,github.com/openshift/origin/pkg/image/api,github.com/openshift/origin/pkg/image/api/v1,github.com/openshift/origin/pkg/oauth/api,github.com/openshift/origin/pkg/oauth/api/v1,github.com/openshift/origin/pkg/project/api,github.com/openshift/origin/pkg/project/api/v1,github.com/openshift/origin/pkg/quota/api,github.com/openshift/origin/pkg/quota/api/v1,github.com/openshift/origin/pkg/route/api,github.com/openshift/origin/pkg/route/api/v1,github.com/openshift/origin/pkg/sdn/api,github.com/openshift/origin/pkg/sdn/api/v1,github.com/openshift/origin/pkg/template/api,github.com/openshift/origin/pkg/template/api/v1,github.com/openshift/origin/pkg/user/api,github.com/openshift/origin/pkg/user/api/v1] --logtostderr=true

package internalversion

import (
	api "github.com/openshift/origin/pkg/template/api"
	pkg_api "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/errors"
	"k8s.io/kubernetes/pkg/client/cache"
	"k8s.io/kubernetes/pkg/labels"
)

// BrokerTemplateInstanceLister helps list BrokerTemplateInstances.
type BrokerTemplateInstanceLister interface {
	// List lists all BrokerTemplateInstances in the indexer.
	List(selector labels.Selector) (ret []*api.BrokerTemplateInstance, err error)
	// Get retrieves the BrokerTemplateInstance from the index for a given name.
	Get(name string) (*api.BrokerTemplateInstance, error)
	BrokerTemplateInstanceListerExpansion
}

// brokerTemplateInstanceLister implements the BrokerTemplateInstanceLister interface.
type brokerTemplateInstanceLister struct {
	indexer cache.Indexer
}

// NewBrokerTemplateInstanceLister returns a new BrokerTemplateInstanceLister.
func NewBrokerTemplateInstanceLister(indexer cache.Indexer) BrokerTemplateInstanceLister {
	return &brokerTemplateInstanceLister{indexer: indexer}
}

// List lists all BrokerTemplateInstances in the indexer.
func (s *brokerTemplateInstanceLister) List(selector labels.Selector) (ret []*api.BrokerTemplateInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*api.BrokerTemplateInstance))
	})
	return ret, err
}

// Get retrieves the BrokerTemplateInstance from the index for a given name.
func (s *brokerTemplateInstanceLister) Get(name string) (*api.BrokerTemplateInstance, error) {
	key := &api.BrokerTemplateInstance{ObjectMeta: pkg_api.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(api.Resource("brokertemplateinstance"), name)
	}
	return obj.(*api.BrokerTemplateInstance), nil
}
