package local

import (
	"fmt"
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/vishvananda/netlink"
)

// DisconnectNetwork disconnects a container from a network
func (c *localController) DisconnectNetwork(networkName string, containerPid int) error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	logrus.Debugf("disconnecting %d from networks %s", containerPid, networkName)

	localPeerName := getLocalPeerName(networkName, containerPid)
	iface, err := netlink.LinkByName(localPeerName)
	if err != nil {
		return fmt.Errorf("error getting local peer link: %s", err)
	}

	if err := c.ipam.ReleasePeersForPid(networkName, containerPid); err != nil {
		return err
	}

	if err := netlink.LinkSetDown(iface); err != nil {
		return fmt.Errorf("error downing interface: %s", err)
	}

	if err := netlink.LinkDel(iface); err != nil {
		return err
	}

	return nil
}
