package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/buaazp/diskutil"
)

var (
	megaPath     string
	adapterCount int
)

func init() {
	flag.StringVar(&megaPath, "mega-path", "/opt/MegaRAID/MegaCli/MegaCli64", "megaCli binary path")
	flag.IntVar(&adapterCount, "adapter-count", 0, "adapter count in your server")
}

func main() {
	flag.Parse()
	ds, err := diskutil.NewDiskStatus(megaPath, adapterCount)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "DiskStatus New error: %v\n", err)
		return
	}

	err = ds.Get()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "DiskStatus Get error: %v\n", err)
		return
	}

	for i, ads := range ds.AdapterStats {
		fmt.Printf("adapter #%d \n", i)
		for j, pds := range ads.PhysicalDriveStats {
			pdStatus := pds.FirmwareState
			pdName := []string{pds.Brand, pds.Model, pds.SerialNumber}
			pdSN := strings.Join(pdName, " ")
			fmt.Printf("PD%d: %s status: %s size %s\n", j, pdSN, pdStatus, pds.RawSize)
		}
		fmt.Printf("\n")
	}

	brokenVds, brokenPds, err := ds.ListBrokenDrive()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "DiskStatus ListBrokenDrive error: %v\n", err)
		return
	}
	for _, bvd := range brokenVds {
		fmt.Println(bvd)
	}
	for _, bpd := range brokenPds {
		fmt.Println(bpd)
	}

	jsonStatus, err := ds.ToJson()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "DiskStatus ToJson error: %v\n", err)
		return
	}
	fmt.Println(jsonStatus)
}
