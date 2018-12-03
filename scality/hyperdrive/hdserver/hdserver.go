package hdserver

import (
	"fmt"
	"sync"

	"github.com/fferrandis/simu/scality/hyperdrive/disk"
	"github.com/fferrandis/simu/scality/hyperdrive/diskset"
	. "github.com/fferrandis/simu/scality/hyperdrive/group"
)

type HDSrv struct {
	dset *diskset.DiskSet

	group Group

	nrdata     int
	nrcoding   int
	extentsize uint64
	sync.Mutex
}

func (hdsrv *HDSrv) HDSrvGroupInit(ts uint64) bool {
	datadisk := make([]*disk.Disk, hdsrv.nrdata)
	codingdisk := make([]*disk.Disk, hdsrv.nrcoding)
	hdsrv.dset.Select(datadisk, codingdisk, hdsrv.nrdata, hdsrv.nrcoding)
	hdsrv.group.GroupInit(datadisk, codingdisk, hdsrv.extentsize, hdsrv.nrdata, hdsrv.nrcoding, ts)

	return true
}

func NewHDSrv(nrdata int,
	nrcoding int,
	extentsize uint64,
	data *disk.Disk,
	numberof_disk int,
	ts uint64) *HDSrv {

	h := &HDSrv{
		nrdata:     nrdata,
		nrcoding:   nrcoding,
		extentsize: extentsize,
		dset:       diskset.New(numberof_disk, data),
	}

	h.HDSrvGroupInit(ts)
	return h
}

func (hdsrv *HDSrv) PutData(datalen uint64, ts uint64) (bool, uint64) {
	r := true
	load := uint64(0)
	hdsrv.Lock()
	defer hdsrv.Unlock()
	r, load = hdsrv.group.PutData(datalen, ts)
	if r == false {
		fmt.Println("Group full, create new one")
		hdsrv.HDSrvGroupInit(ts)
		r, load = hdsrv.group.PutData(datalen, ts)
	}

	return r, load
}
