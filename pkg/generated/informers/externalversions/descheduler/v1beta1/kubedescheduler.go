/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	time "time"

	deschedulerv1beta1 "github.com/openshift/cluster-kube-descheduler-operator/pkg/apis/descheduler/v1beta1"
	versioned "github.com/openshift/cluster-kube-descheduler-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/openshift/cluster-kube-descheduler-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/openshift/cluster-kube-descheduler-operator/pkg/generated/listers/descheduler/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KubeDeschedulerInformer provides access to a shared informer and lister for
// KubeDeschedulers.
type KubeDeschedulerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.KubeDeschedulerLister
}

type kubeDeschedulerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKubeDeschedulerInformer constructs a new informer for KubeDescheduler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKubeDeschedulerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKubeDeschedulerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKubeDeschedulerInformer constructs a new informer for KubeDescheduler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKubeDeschedulerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubedeschedulersV1beta1().KubeDeschedulers(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubedeschedulersV1beta1().KubeDeschedulers(namespace).Watch(options)
			},
		},
		&deschedulerv1beta1.KubeDescheduler{},
		resyncPeriod,
		indexers,
	)
}

func (f *kubeDeschedulerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKubeDeschedulerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kubeDeschedulerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&deschedulerv1beta1.KubeDescheduler{}, f.defaultInformer)
}

func (f *kubeDeschedulerInformer) Lister() v1beta1.KubeDeschedulerLister {
	return v1beta1.NewKubeDeschedulerLister(f.Informer().GetIndexer())
}
