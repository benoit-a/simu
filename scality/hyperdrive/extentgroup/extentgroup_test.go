package extentgroup

import (
	"fmt"
	"testing"

	"github.com/fferrandis/simu/scality/hyperdrive/disk"
	"github.com/fferrandis/simu/scality/hyperdrive/diskset"
)

func TestCreateGroup(t *testing.T) {
	var model = disk.New(2000, 10, 5, 0)
	var group ExtentDataGroup
	var coding_g ExtentCodingGroup

	set := diskset.New(8, model)

	data := make([]*disk.Disk, 4)
	coding := make([]*disk.Disk, 2)

	r := set.DiskSetSelect(data, coding, 4, 2)
	if r != true {
		t.Error("expected success")
	}

	group.ExtentDataGroupInit(data, 4, 2000, 0)
	coding_g.ExtentCodingGroupInit(coding, 2, 2000, 0)

	/* extent id 0 should be at the end of our group of extent after */
	group.ExtentDataGroupPutData(1000, 0)
	if group.list[3].id != 0 {
		t.Error("extent id 0 should be at the end after data insertion")
	}

	group.ExtentDataGroupPutData(1000, 0)
	if group.list[2].id != 1 {
		t.Error("extent id 1 should be just before  the end after data insertion")
	}

	group.ExtentDataGroupClose(0)
	coding_g.ExtentCodingGroupClose(0)

	for i := 0; i < 4; i++ {
		fmt.Println("data : ", data[i])
	}

	for i := 0; i < 2; i++ {
		fmt.Println("coding : ", coding[i])
	}

}
