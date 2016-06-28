Flash partition calculation utility
===================================

This utility can be used to run calculations on flash partition layout.
I stated doing this in a spreadsheet and quickly got frustrated at 
the difficulty in working with hex numbers.  Using Go's text/tabwriter,
the output can be formatted nearly as nicely as a spreadsheet, yet
with the full power of a real programming language to generate the
data.

To run:

* install Go >= 1.6
* go run go-partition-calcs.go

Example output:

```
Old Partitions
  Device       Name       Start     Size(B)  Size(MiB)  Size(Blks)
    mtd0    xloader         0x0     0x80000        0.5           4
    mtd1      uboot     0x80000    0x1c0000       1.75          14
    mtd2  uboot env    0x240000     0x40000       0.25           2
    mtd3      linux    0x280000    0xa00000         10          80
    mtd4     rootfs    0xc80000  0x186a0000    390.625        3125
    mtd5       data  0x19320000  0x26ce0000    620.875        4967
total size = 0x40000000, 1024MiB



New Partitions 500MB
  Device       Name       Start    Size(B)  Size(MiB)  Size(Blks)
    mtd0    xloader         0x0    0x80000        0.5           4
    mtd1      uboot     0x80000   0x500000          5          40
    mtd2  uboot env    0x580000    0x40000       0.25           2
    mtd3     linux1    0x5c0000  0x1400000         20         160
    mtd4     linux2   0x19c0000  0x1400000         20         160
    mtd5    rootfs1   0x2dc0000  0xa000000        160        1280
    mtd6    rootfs2   0xcdc0000  0xa000000        160        1280
    mtd7        log  0x16dc0000   0xa00000         10          80
    mtd8       data  0x177c0000  0x7c40000     124.25         994
total size = 0x1f400000, 500MiB

size of boot partitions:  45.75 MiB


New Partitions 1G
  Device       Name       Start     Size(B)  Size(MiB)  Size(Blks)
    mtd0    xloader         0x0     0x80000        0.5           4
    mtd1      uboot     0x80000    0x500000          5          40
    mtd2  uboot env    0x580000     0x40000       0.25           2
    mtd3     linux1    0x5c0000   0x1400000         20         160
    mtd4     linux2   0x19c0000   0x1400000         20         160
    mtd5    rootfs1   0x2dc0000   0xfa00000        250        2000
    mtd6    rootfs2  0x127c0000   0xfa00000        250        2000
    mtd7        log  0x221c0000   0x3200000         50         400
    mtd8       data  0x253c0000  0x1ac40000     428.25        3426
total size = 0x40000000, 1024MiB

size of boot partitions:  45.75 MiB


New Partitions 2G
  Device       Name       Start     Size(B)  Size(MiB)  Size(Blks)
    mtd0    xloader         0x0     0x80000        0.5           4
    mtd1      uboot     0x80000    0x500000          5          40
    mtd2  uboot env    0x580000     0x40000       0.25           2
    mtd3     linux1    0x5c0000   0x1400000         20         160
    mtd4     linux2   0x19c0000   0x1400000         20         160
    mtd5    rootfs1   0x2dc0000  0x1f400000        500        4000
    mtd6    rootfs2  0x221c0000  0x1f400000        500        4000
    mtd7        log  0x415c0000   0x6400000        100         800
    mtd8       data  0x479c0000  0x38640000     902.25        7218
total size = 0x80000000, 2048MiB

size of boot partitions:  45.75 MiB

```


