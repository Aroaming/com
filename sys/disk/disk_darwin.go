package disk

import (
	"strconv"
	"strings"
)

const (
	dfCmd   = "df -k -l -P"
	linkDir = "/private/var"
)

//split to struct
func splitLine(line string) DiskInfo {
	var disk DiskInfo
	str := strings.Split(trimLine(line), " ")
	// ignore path /private/var
	if len(str) != 6 || strings.Contains(str[5], linkDir) {
		return disk
	}
	disk.Name = str[0]
	disk.Size, _ = strconv.Atoi(str[1])
	disk.SizeUsed, _ = strconv.Atoi(str[2])
	disk.SizeUnUsed, _ = strconv.Atoi(str[3])
	disk.Proportion = str[4]
	disk.MountPath = str[5]
	return disk
}
