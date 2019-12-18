package disk

import (
	"regexp"
	"strconv"
	"strings"
	"syscall"
)

type DiskInfo struct {
	Name       string
	Type       string
	Size       int
	SizeUsed   int
	SizeUnUsed int
	Proportion string
	MountPath  string
}

var (
	disks []DiskInfo
)

func All() []DiskInfo {
	return disks
}

func Max() DiskInfo {
	var (
		disk DiskInfo
		max  int
	)
	for k, v := range disks {
		if max < v.Size {
			max = v.Size
			disk = disks[k]
		}
	}
	return disk
}

func splitLine(line string) DiskInfo {
	var disk DiskInfo
	str := strings.Split(deleteExtraSpace(line), " ")
	if len(str) != 7 {
		return disk
	}
	disk.Name = str[0]
	disk.Type = str[1]
	disk.Size, _ = strconv.Atoi(str[2])
	disk.SizeUsed, _ = strconv.Atoi(str[3])
	disk.SizeUnUsed, _ = strconv.Atoi(str[4])
	disk.Proportion = str[5]
	disk.MountPath = str[6]
	return disk
}

func deleteExtraSpace(s string) string {
	s1 := strings.Replace(s, "  ", " ", -1)
	regstr := "\\s{2,}"
	reg, _ := regexp.Compile(regstr)
	s2 := make([]byte, len(s1))
	copy(s2, s1)
	spc_index := reg.FindStringIndex(string(s2))
	for len(spc_index) > 0 {
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...)
		spc_index = reg.FindStringIndex(string(s2))
	}
	return string(s2)
}

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

// disk usage of path/disk
func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}
