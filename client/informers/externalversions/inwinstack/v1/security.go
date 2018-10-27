/*
Copyright © 2018 inwinSTACK.inc

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

package v1

import (
	time "time"

	inwinstackv1 "github.com/inwinstack/blended/apis/inwinstack/v1"
	versioned "github.com/inwinstack/blended/client/clientset/versioned"
	internalinterfaces "github.com/inwinstack/blended/client/informers/externalversions/internalinterfaces"
	v1 "github.com/inwinstack/blended/client/listers/inwinstack/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SecurityInformer provides access to a shared informer and lister for
// Securities.
type SecurityInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.SecurityLister
}

type securityInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSecurityInformer constructs a new informer for Security type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSecurityInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSecurityInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSecurityInformer constructs a new informer for Security type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSecurityInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InwinstackV1().Securities(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InwinstackV1().Securities(namespace).Watch(options)
			},
		},
		&inwinstackv1.Security{},
		resyncPeriod,
		indexers,
	)
}

func (f *securityInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSecurityInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *securityInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&inwinstackv1.Security{}, f.defaultInformer)
}

func (f *securityInformer) Lister() v1.SecurityLister {
	return v1.NewSecurityLister(f.Informer().GetIndexer())
}
