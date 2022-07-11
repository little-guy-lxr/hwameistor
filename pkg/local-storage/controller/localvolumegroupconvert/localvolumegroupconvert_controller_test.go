package localvolumegroupconvert

import (
	"context"
	"fmt"
	"testing"
	"time"

	ldmv1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/hwameistor/local-disk-manager/v1alpha1"
	apisv1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/hwameistor/local-storage/v1alpha1"
	"github.com/hwameistor/hwameistor/pkg/local-storage/member"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var (
	fakeLocalVolumeGroupName              = "local-volume-group-example"
	fakeLocalVolumeGroupConvertName       = "local-volume-group-convert-example"
	fakeLocalVolumeGroupConvertUID        = "local-volume-group-convert-uid"
	fakeNamespace                         = "local-volume-group-test"
	fakeNodename                          = "10-6-118-10"
	fakeStorageIp                         = "10.6.118.11"
	fakeZone                              = "zone-test"
	fakeRegion                            = "region-test"
	fakeVgType                            = "LocalStorage_PoolHDD"
	fakeVgName                            = "vg-test"
	fakePoolClass                         = "HDD"
	fakePoolType                          = "REGULAR"
	fakeTotalCapacityBytes          int64 = 10 * 1024 * 1024 * 1024
	fakeFreeCapacityBytes           int64 = 8 * 1024 * 1024 * 1024
	fakeDiskCapacityBytes           int64 = 2 * 1024 * 1024 * 1024

	apiversion                  = "hwameistor.io/v1alpha1"
	LocalVolumeGroupConvertKind = "LocalVolumeGroupConvert"
	fakeRecorder                = record.NewFakeRecorder(100)
)

func TestNewLocalVolumeGroupConvertController(t *testing.T) {

	cli, s := CreateFakeClient()
	// Create a Reconcile for LocalVolumeGroupConvert
	r := ReconcileLocalVolumeGroupConvert{
		client:        cli,
		scheme:        s,
		storageMember: member.Member().ConfigureController(s),
	}

	// Create LocalVolumeGroupConvert
	lsn := GenFakeLocalVolumeGroupConvertObject()
	err := r.client.Create(context.Background(), lsn)
	if err != nil {
		t.Errorf("Create LocalVolumeGroupConvert fail %v", err)
	}
	defer r.DeleteFakeLocalVolumeGroupConvert(t, lsn)

	// Get lsn
	err = r.client.Get(context.Background(), types.NamespacedName{Namespace: lsn.GetNamespace(), Name: lsn.GetName()}, lsn)
	if err != nil {
		t.Errorf("Get lsn fail %v", err)
	}
	fmt.Printf("lsn = %+v", lsn)
	fmt.Printf("r.storageMember = %+v", r.storageMember)

	// Mock LocalVolumeGroupConvert request
	req := reconcile.Request{NamespacedName: types.NamespacedName{Namespace: lsn.GetNamespace(), Name: lsn.GetName()}}
	_, err = r.Reconcile(req)
	if err != nil {
		t.Errorf("Reconcile fail %v", err)
	}

}

// DeleteFakeLocalVolumeGroupConvert
func (r *ReconcileLocalVolumeGroupConvert) DeleteFakeLocalVolumeGroupConvert(t *testing.T, lsn *apisv1alpha1.LocalVolumeGroupConvert) {
	if err := r.client.Delete(context.Background(), lsn); err != nil {
		t.Errorf("Delete LocalVolumeGroupConvert %v fail %v", lsn.GetName(), err)
	}
}

// GenFakeLocalVolumeGroupConvertObject Create lsn request
func GenFakeLocalVolumeGroupConvertObject() *apisv1alpha1.LocalVolumeGroupConvert {
	lsn := &apisv1alpha1.LocalVolumeGroupConvert{}

	TypeMeta := metav1.TypeMeta{
		Kind:       LocalVolumeGroupConvertKind,
		APIVersion: apiversion,
	}

	ObjectMata := metav1.ObjectMeta{
		Name:              fakeLocalVolumeGroupConvertName,
		Namespace:         fakeNamespace,
		ResourceVersion:   "",
		UID:               types.UID(fakeLocalVolumeGroupConvertUID),
		CreationTimestamp: metav1.Time{Time: time.Now()},
	}

	Spec := apisv1alpha1.LocalVolumeGroupConvertSpec{
		LocalVolumeGroupName: fakeLocalVolumeGroupName,
		ReplicaNumber:        1,
		Abort:                true,
	}

	disks := make([]apisv1alpha1.LocalDisk, 0, 10)
	var localdisk1 apisv1alpha1.LocalDisk
	localdisk1.DevPath = "/dev/sdf"
	localdisk1.State = apisv1alpha1.DiskStateAvailable
	localdisk1.Class = fakePoolClass
	localdisk1.CapacityBytes = fakeDiskCapacityBytes
	disks = append(disks, localdisk1)

	volumes := make([]string, 0, 5)
	volumes = append(volumes, "volume-test1")

	pools := make(map[string]apisv1alpha1.LocalPool)
	pools[fakeVgType] = apisv1alpha1.LocalPool{
		Name:                     fakeVgName,
		Class:                    fakePoolClass,
		Type:                     fakePoolType,
		TotalCapacityBytes:       int64(fakeTotalCapacityBytes),
		UsedCapacityBytes:        int64(fakeTotalCapacityBytes) - int64(fakeFreeCapacityBytes),
		FreeCapacityBytes:        int64(fakeFreeCapacityBytes),
		VolumeCapacityBytesLimit: int64(fakeTotalCapacityBytes),
		TotalVolumeCount:         apisv1alpha1.LVMVolumeMaxCount,
		UsedVolumeCount:          int64(len(volumes)),
		FreeVolumeCount:          apisv1alpha1.LVMVolumeMaxCount - int64(len(volumes)),
		Disks:                    disks,
		Volumes:                  volumes,
	}

	lsn.ObjectMeta = ObjectMata
	lsn.TypeMeta = TypeMeta
	lsn.Spec = Spec
	lsn.Status.State = apisv1alpha1.VolumeStateCreating
	return lsn
}

// CreateFakeClient Create LocalVolumeGroupConvert resource
func CreateFakeClient() (client.Client, *runtime.Scheme) {
	lsn := GenFakeLocalVolumeGroupConvertObject()
	lsnList := &apisv1alpha1.LocalVolumeGroupConvertList{
		TypeMeta: metav1.TypeMeta{
			Kind:       LocalVolumeGroupConvertKind,
			APIVersion: apiversion,
		},
	}

	s := scheme.Scheme
	s.AddKnownTypes(ldmv1alpha1.SchemeGroupVersion, lsn)
	s.AddKnownTypes(ldmv1alpha1.SchemeGroupVersion, lsnList)
	return fake.NewFakeClientWithScheme(s), s
}
