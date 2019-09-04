package main

import (
	"archive/tar"
	"bytes"
	"flag"
	"github.com/badeadan/starlingx-vbox-installer/pkg/lab"
	"log"
	"os"
)

type TarWriter struct {
	*tar.Writer
}

func (tw *TarWriter) WriteFileBytes(name string, mode int64, buffer *bytes.Buffer) error {
	err := tw.WriteHeader(&tar.Header{
		Name: name,
		Mode: mode,
		Size: int64(buffer.Len()),
	})
	if err == nil {
		_, err = tw.Write(buffer.Bytes())
	}
	return err
}

func main() {
	sl := lab.StandardLab{SystemMode: "duplex"}
	flag.StringVar(&sl.Name, "name", "standard", "group name")
	flag.StringVar(&sl.NatNet, "nat-net", "nat4", "nat network name")
	flag.StringVar(&sl.LoopBackPrefix, "loop-prefix", "127.0.4", "nat loopback prefix")
	flag.StringVar(&sl.Oam.Network, "oam-network", "10.10.10.0/24", "oam network address")
	flag.StringVar(&sl.Oam.Gateway, "oam-gateway", "10.10.10.1", "oam gateway")
	flag.StringVar(&sl.Oam.FloatAddr, "oam-float", "10.10.10.2", "oam floating ip")
	flag.StringVar(&sl.Oam.Controller0, "oam-ctrl-0", "10.10.10.3", "oam controller-0 ip")
	flag.StringVar(&sl.Oam.Controller1, "oam-ctrl-1", "10.10.10.4", "oam controller-1 ip")
	flag.StringVar(&sl.IntNetPrefix, "intnet-prefix", "intnet", "internal network  prefix")
	flag.UintVar(&sl.ControllerCpus, "controller-cpus", 4, "controller cpu count")
	flag.UintVar(&sl.ControllerMemory, "controller-memory", 16, "controller ram size")
	flag.UintVar(&sl.ControllerDiskSize, "controller-disk-size", 520, "controller disk size")
	flag.UintVar(&sl.ComputeCount, "compute-count", 2, "number of compute hosts")
	flag.UintVar(&sl.ComputeCpus, "compute-cpus", 4, "compute cpu count")
	flag.UintVar(&sl.ComputeMemory, "compute-memory", 10, "compute ram size")
	flag.UintVar(&sl.ComputeDiskSize, "compute-disk", 520, "compute disk size")

	flag.Parse()
	err := lab.MakeStandardInstaller(sl, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
