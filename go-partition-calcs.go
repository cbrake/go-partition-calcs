package main

import (
	"bytes"
	"fmt"
	"text/tabwriter"
)

// to run this: go run go-partition-calcs.go

const (
	nandBlockSize = 128 * 1024
	KiB           = 1024
	MiB           = 1024 * 1024
	GiB           = 1024 * 1024 * 1024
)

type Partition struct {
	Device string
	Name   string
	Start  uint
	Size   uint
}

func (p Partition) String() string {
	sizeMiB := float32(p.Size) / (1024 * 1024)
	sizeBlocks := float32(p.Size) / nandBlockSize
	return fmt.Sprintf("%v\t%v\t0x%x\t0x%x\t%v\t%v\t",
		p.Device, p.Name, p.Start, p.Size, sizeMiB, sizeBlocks)
}

type Partitions []Partition

func (parts Partitions) String() string {
	w := new(tabwriter.Writer)
	buf := new(bytes.Buffer)
	w.Init(buf, 5, 0, 2, ' ', tabwriter.AlignRight)

	fmt.Fprintln(w, "Device\tName\tStart\tSize(B)\tSize(MiB)\tSize(Blks)\t")

	for _, p := range parts {
		fmt.Fprintln(w, p.String())
	}

	w.Flush()

	size := parts.CalcSize()
	sizeMiB := float32(size) / (1024 * 1024)

	fmt.Fprintf(w, "total size = 0x%x, %vMiB\n", size, sizeMiB)

	return buf.String()
}

func (parts Partitions) Get(part int) Partition {
	parts_ := []Partition(parts)
	return parts_[part]
}

func (parts *Partitions) FillIn(deviceSize uint, align uint) {
	curAdr := uint(0)
	parts_ := []Partition(*parts)
	mtdDev := 0

	for i, p := range parts_ {
		// first round size to align size
		if align > 0 {
			parts_[i].Size = (p.Size / nandBlockSize) * nandBlockSize
		}

		parts_[i].Start = curAdr
		curAdr += p.Size
		parts_[i].Device = fmt.Sprintf("mtd%v", mtdDev)
		mtdDev += 1
	}

	lastI := len(parts_) - 1

	if deviceSize > 0 && parts_[lastI].Size == 0 {
		// fill in size of last partition
		parts_[lastI].Size = deviceSize - curAdr
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
	old := Partitions{
		Partition{Name: "xloader", Size: 0x80000},
		Partition{Name: "uboot", Size: 0x1c0000},
		Partition{Name: "uboot env", Size: 0x40000},
		Partition{Name: "linux", Size: 0xa00000},
		Partition{Name: "rootfs", Size: 0x186a0000},
		Partition{Name: "data", Size: 0x26ce0000},
	}

	old.FillIn(0, 0)

	fmt.Printf("Old Partitions\n%v\n", old)

	new500m := Partitions{
		Partition{Name: "xloader", Size: 0x80000},
		Partition{Name: "uboot", Size: 5 * MiB},
		Partition{Name: "uboot env", Size: 0x40000},
		Partition{Name: "linux1", Size: 20 * MiB},
		Partition{Name: "linux2", Size: 20 * MiB},
		Partition{Name: "rootfs1", Size: 160 * MiB},
		Partition{Name: "rootfs2", Size: 160 * MiB},
		Partition{Name: "log", Size: 10 * MiB},
		Partition{Name: "data"},
	}

	new500m.FillIn(500*MiB, nandBlockSize)

	fmt.Printf("\n\nNew Partitions 500MB\n%v\n", new500m)

	bootSize := uint(0)
	for i := 0; i < 5; i++ {
		bootSize += new500m.Get(i).Size
	}

	fmt.Println("size of boot partitions: ", float32(bootSize)/MiB, "MiB")

	new1g := Partitions{
		Partition{Name: "xloader", Size: 0x80000},
		Partition{Name: "uboot", Size: 5 * MiB},
		Partition{Name: "uboot env", Size: 0x40000},
		Partition{Name: "linux1", Size: 20 * MiB},
		Partition{Name: "linux2", Size: 20 * MiB},
		Partition{Name: "rootfs1", Size: 250 * MiB},
		Partition{Name: "rootfs2", Size: 250 * MiB},
		Partition{Name: "log", Size: 50 * MiB},
		Partition{Name: "data"},
	}

	new1g.FillIn(GiB, nandBlockSize)

	fmt.Printf("\n\nNew Partitions 1G\n%v\n", new1g)

	bootSize = uint(0)
	for i := 0; i < 5; i++ {
		bootSize += new1g.Get(i).Size
	}

	fmt.Println("size of boot partitions: ", float32(bootSize)/MiB, "MiB")

	new2g := Partitions{
		Partition{Name: "xloader", Size: 0x80000},
		Partition{Name: "uboot", Size: 5 * MiB},
		Partition{Name: "uboot env", Size: 0x40000},
		Partition{Name: "linux1", Size: 20 * MiB},
		Partition{Name: "linux2", Size: 20 * MiB},
		Partition{Name: "rootfs1", Size: 500 * MiB},
		Partition{Name: "rootfs2", Size: 500 * MiB},
		Partition{Name: "log", Size: 100 * MiB},
		Partition{Name: "data"},
	}

	new2g.FillIn(2*GiB, nandBlockSize)

	fmt.Printf("\n\nNew Partitions 2G\n%v\n", new2g)

	bootSize = uint(0)
	for i := 0; i < 5; i++ {
		bootSize += new2g.Get(i).Size
	}

	fmt.Println("size of boot partitions: ", float32(bootSize)/MiB, "MiB")

}
