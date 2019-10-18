// +build linux
package disk

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

const (
	cmd = "df -k -T -t ext4 -t ext3 -t xfs -t zfs"
)

func init() {
	cmd := exec.Command("bash", "-c", cmd)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	//read title
	out.ReadString('\n')
	for {
		line, err := out.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		disks = append(disks, splitLine(line))
	}
}
