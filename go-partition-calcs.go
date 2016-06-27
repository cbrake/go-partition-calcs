package main

import (
	"bytes"
	"fmt"
	"text/tabwriter"
)

// to run this: go run thmi-nand-partitions.go

/*
mtd0: 00080000 00020000 "xloader"
mtd1: 001c0000 00020000 "uboot"
mtd2: 00040000 00020000 "uboot environment"
mtd3: 00a00000 00020000 "linux"
mtd4: 186a0000 00020000 "rootfs"
mtd5: 26ce0000 00020000 "data"

Nand block size is 128KB

*/

type Partition struct {
	Device string
	Name   string
	Start  uint
	Size   uint
}

func (p Partition) String() string {
	return fmt.Sprintf("%v\t%v\t0x%x\t0x%x\t", p.Device, p.Name, p.Start, p.Size)
}

type Partitions []Partition

func (parts Partitions) String() string {
	w := new(tabwriter.Writer)
	buf := new(bytes.Buffer)
	w.Init(buf, 10, 0, 5, ' ', tabwriter.AlignRight)

	fmt.Fprintln(w, "Device\tName\tStart\tSize\t")

	for _, p := range parts {
		fmt.Fprintln(w, p.String())
	}

	w.Flush()

	size := parts.CalcSize()
	sizeMiB := float32(size) / (1024 * 1024)

	fmt.Fprintf(w, "total size = 0x%x, %vMiB\n", size, sizeMiB)

	return buf.String()
}

func (parts *Partitions) CalcStart() {
	curAdr := uint(0)
	parts_ := []Partition(*parts)
	for i, p := range parts_ {
		parts_[i].Start = curAdr
		curAdr += p.Size
	}
}

func (parts Partitions) CalcSize() uint {
	size := uint(0)
	parts_ := []Partition(parts)
	for _, p := range parts_ {
		size += p.Size
	}

	return size
}

func main() {
	fmt.Println("THMI partitions")

	old := Partitions{
		Partition{Device: "mtd0", Name: "xloader", Size: 0x80000},
		Partition{Device: "mtd1", Name: "uboot", Size: 0x1c0000},
		Partition{Device: "mtd2", Name: "uboot env", Size: 0x40000},
		Partition{Device: "mtd3", Name: "linux", Size: 0xa00000},
		Partition{Device: "mtd4", Name: "rootfs", Size: 0x186a0000},
		Partition{Device: "mtd5", Name: "data", Size: 0x26ce0000},
	}

	old.CalcStart()

	fmt.Printf("Old Partitions\n%v\n", old)
}
