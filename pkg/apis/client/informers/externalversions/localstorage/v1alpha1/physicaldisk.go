// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	versioned "github.com/hwameistor/local-storage/pkg/apis/client/clientset/versioned"
	internalinterfaces "github.com/hwameistor/local-storage/pkg/apis/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/hwameistor/local-storage/pkg/apis/client/listers/localstorage/v1alpha1"
	localstoragev1alpha1 "github.com/hwameistor/local-storage/pkg/apis/localstorage/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PhysicalDiskInformer provides access to a shared informer and lister for
// PhysicalDisks.
type PhysicalDiskInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PhysicalDiskLister
}

type physicalDiskInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewPhysicalDiskInformer constructs a new informer for PhysicalDisk type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPhysicalDiskInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPhysicalDiskInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredPhysicalDiskInformer constructs a new informer for PhysicalDisk type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPhysicalDiskInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LocalstorageV1alpha1().PhysicalDisks().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LocalstorageV1alpha1().PhysicalDisks().Watch(context.TODO(), options)
			},
		},
		&localstoragev1alpha1.PhysicalDisk{},
		resyncPeriod,
		indexers,
	)
}

func (f *physicalDiskInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPhysicalDiskInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *physicalDiskInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&localstoragev1alpha1.PhysicalDisk{}, f.defaultInformer)
}

func (f *physicalDiskInformer) Lister() v1alpha1.PhysicalDiskLister {
	return v1alpha1.NewPhysicalDiskLister(f.Informer().GetIndexer())
}
