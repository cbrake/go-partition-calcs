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
     Device          Name          Start           Size
       mtd0       xloader            0x0        0x80000
       mtd1         uboot        0x80000       0x1c0000
       mtd2     uboot env       0x240000        0x40000
       mtd3         linux       0x280000       0xa00000
       mtd4        rootfs       0xc80000     0x186a0000
       mtd5          data     0x19320000     0x26ce0000
total size = 0x40000000, 1024MiB
```


