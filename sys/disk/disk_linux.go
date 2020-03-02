package disk

import (
	"strconv"
	"strings"
)

const (
	dfCmd = "df -k -T -t ext4 -t ext3 -t xfs -t zfs"
)

func splitLine(line string) DiskInfo {
	var disk DiskInfo
	str := strings.Split(trimLine(line), " ")
	if len(str) != 7 {
		return disk
	}
	disk.Name = str[0]
	disk.Size, _ = strconv.Atoi(str[2])
	disk.SizeUsed, _ = strconv.Atoi(str[3])
	disk.SizeUnUsed, _ = strconv.Atoi(str[4])
	disk.Proportion = str[5]
	disk.MountPath = str[6]
	return disk
}
