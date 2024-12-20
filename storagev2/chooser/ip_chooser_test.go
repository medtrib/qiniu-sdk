//go:build unit
// +build unit

package chooser_test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/medtrib/qiniu-sdk/v7/storagev2/chooser"
)

func TestIPChooser(t *testing.T) {
	ipc := chooser.NewIPChooser(&chooser.IPChooserConfig{
		FreezeDuration: 2 * time.Second,
	})

	ips := ipc.Choose(context.Background(), nil, &chooser.ChooseOptions{
		Domain: "www.qiniu.com",
	})
	assertIPs(t, ips, []net.IP{})

	ips = ipc.Choose(context.Background(), []net.IP{net.IPv4(1, 2, 3, 4), net.IPv4(5, 6, 7, 8)}, &chooser.ChooseOptions{
		Domain: "www.qiniu.com",
	})
	assertIPs(t, ips, []net.IP{net.IPv4(1, 2, 3, 4), net.IPv4(5, 6, 7, 8)})

	ipc.FeedbackBad(context.Background(), []net.IP{net.IPv4(1, 2, 3, 4), net.IPv4(5, 6, 7, 8)}, &chooser.FeedbackOptions{
		Domain: "www.qiniu.com",
	})
	ips = ipc.Choose(context.Background(), []net.IP{net.IPv4(1, 2, 3, 4), net.IPv4(5, 6, 7, 8)}, &chooser.ChooseOptions{
		Domain: "www.qiniu.com",
	})
	assertIPs(t, ips, []net.IP{net.IPv4(1, 2, 3, 4), net.IPv4(5, 6, 7, 8)})
	ipc.FeedbackGood(context.Background(), []net.IP{net.IPv4(1, 2, 3, 4)}, &chooser.FeedbackOptions{
		Domain: "www.qiniu.com",
	})
	ips = ipc.Choose(context.Background(), []net.IP{net.IPv4(1, 2, 3, 4), net.IPv4(5, 6, 7, 8)}, &chooser.ChooseOptions{
		Domain: "www.qiniu.com",
	})
	assertIPs(t, ips, []net.IP{net.IPv4(1, 2, 3, 4)})

	time.Sleep(2 * time.Second)
	ips = ipc.Choose(context.Background(), []net.IP{net.IPv4(1, 2, 3, 4), net.IPv4(5, 6, 7, 8)}, &chooser.ChooseOptions{
		Domain: "www.qiniu.com",
	})
	assertIPs(t, ips, []net.IP{net.IPv4(1, 2, 3, 4), net.IPv4(5, 6, 7, 8)})

	ipc.FeedbackBad(context.Background(), []net.IP{net.IPv4(1, 2, 3, 4)}, &chooser.FeedbackOptions{
		Domain: "www.qiniu.com",
	})
	ips = ipc.Choose(context.Background(), []net.IP{net.IPv4(1, 2, 3, 4), net.IPv4(5, 6, 7, 8)}, &chooser.ChooseOptions{
		Domain: "www.qiniu.com",
	})
	assertIPs(t, ips, []net.IP{net.IPv4(5, 6, 7, 8)})

	ipc.FeedbackGood(context.Background(), []net.IP{net.IPv4(1, 2, 3, 4)}, &chooser.FeedbackOptions{
		Domain: "www.qiniu.com",
	})
	ips = ipc.Choose(context.Background(), []net.IP{net.IPv4(1, 2, 3, 4), net.IPv4(5, 6, 7, 8)}, &chooser.ChooseOptions{
		Domain: "www.qiniu.com",
	})
	assertIPs(t, ips, []net.IP{net.IPv4(1, 2, 3, 4), net.IPv4(5, 6, 7, 8)})
}
