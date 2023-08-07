package lsblk

import (
	"encoding/json"
	"strconv"
)

var LvsDeviceCmd = "lvs --units b -o +devices -o -pool_lv -o -origin -o -lv_attr -o -data_percent -o -metadata_percent -o -move_pv -o -mirror_log -o -copy_percent -o -convert_lv"

type MemoryB struct {
	Mem int64
}

func (s *MemoryB) UnmarshalJSON(b []byte) error {
	b = b[1 : len(b)-2]
	i, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}
	s.Mem = i
	return nil
}

type Report struct {
	Report []struct {
		Lv []LvsDevice `json:"lv"`
	} `json:"report"`
}

type LvsDevice struct {
	LVName  string   `json:"lv_name,omitempty"`
	VGName  string   `json:"vg_name,omitempty"`
	LVSize  *MemoryB `json:"lv_size,omitempty"`
	Devices string   `json:"devices,omitempty"`
}

// const (
// 	VGName              = "vg_name"
// 	VGUUID              = "vg_uuid"
// 	VGPVvount           = "pv_count"
// 	VGLvCount           = "lv_count"
// 	VGMaxLv             = "max_lv"
// 	VGMaxPv             = "max_pv"
// 	VGSnapCount         = "snap_count"
// 	VGMissingPvCount    = "vg_missing_pv_count"
// 	VGMetadataCount     = "vg_mda_count"
// 	VGMetadataUsedCount = "vg_mda_used_count"
// 	VGSize              = "vg_size"
// 	VGFreeSize          = "vg_free"
// 	VGMetadataSize      = "vg_mda_size"
// 	VGMetadataFreeSize  = "vg_mda_free"
// 	VGPermissions       = "vg_permissions"
// 	VGAllocationPolicy  = "vg_allocation_policy"

// 	LVName            = "lv_name"
// 	LVFullName        = "lv_full_name"
// 	LVUUID            = "lv_uuid"
// 	LVPath            = "lv_path"
// 	LVDmPath          = "lv_dm_path"
// 	LVActive          = "lv_active"
// 	LVSize            = "lv_size"
// 	LVMetadataSize    = "lv_metadata_size"
// 	LVSegtype         = "segtype"
// 	LVHost            = "lv_host"
// 	LVPool            = "pool_lv"
// 	LVPermissions     = "lv_permissions"
// 	LVWhenFull        = "lv_when_full"
// 	LVHealthStatus    = "lv_health_status"
// 	RaidSyncAction    = "raid_sync_action"
// 	LVDataPercent     = "data_percent"
// 	LVMetadataPercent = "metadata_percent"
// 	LVSnapPercent     = "snap_percent"

// 	PVName             = "pv_name"
// 	PVUUID             = "pv_uuid"
// 	PVInUse            = "pv_in_use"
// 	PVAllocatable      = "pv_allocatable"
// 	PVMissing          = "pv_missing"
// 	PVSize             = "pv_size"
// 	PVFreeSize         = "pv_free"
// 	PVUsedSize         = "pv_used"
// 	PVMetadataSize     = "pv_mda_size"
// 	PVMetadataFreeSize = "pv_mda_free"
// 	PVDeviceSize       = "dev_size"
// )

func LvsReport() (report *Report, err error) {
	output, err := runCmd(LvsDeviceCmd)
	if err != nil {
		return nil, err
	}

	lvsRsp := &Report{}
	err = json.Unmarshal(output, &lvsRsp)
	if err != nil {
		return nil, err
	}

	return lvsRsp, nil
}
