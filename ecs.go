package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	gophercloud "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/compute/v2/servers"
	"os"
)

const (
	ColorDefault = "\x1b[39m"

	ColorRed   = "\x1b[91m"
	ColorGreen = "\x1b[32m"
	ColorBlue  = "\x1b[94m"
	ColorGray  = "\x1b[90m"
)

func ecsList(client *gophercloud.ServiceClient, opts *servers.ListOpts) {
	pages, err := servers.List(client, opts).AllPages()
	if err != nil {
		fmt.Printf("nova list failed, err:%v\n", err)
		panic(err)
	}
	tenant := os.Getenv("OS_USER_DOMAIN_NAME")
	project := os.Getenv("OS_PROJECT_NAME")
	region := os.Getenv("OS_REGION_NAME")
	rst, err := servers.ExtractServers(pages)
	if err != nil {
		panic(err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "STATUS", "NAME"})

	for _, li := range rst {
		if li.Status == "ACTIVE" {
			t.AppendRows([]table.Row{{li.ID, green(li.Status), li.Name}})
		} else if li.Status == "SHUTOFF" {
			t.AppendRows([]table.Row{{li.ID, red(li.Status), li.Name}})
		} else {
			t.AppendRows([]table.Row{{li.ID, gray(li.Status), li.Name}})
		}
	}
	t.AppendFooter(table.Row{blue(tenant), blue(region), blue(project)})
	t.SetStyle(table.StyleColoredBright)
	t.Render()
}

func red(s string) string {
	return fmt.Sprintf("%s%s%s", ColorRed, s, ColorDefault)
}

func green(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, s, ColorDefault)
}

func blue(s string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, s, ColorDefault)
}

func gray(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGray, s, ColorDefault)
}

func main() {

	opts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		panic(err)
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		panic(err)
	}

	ecs, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		panic(err)
	}
	ecsList(ecs, &servers.ListOpts{})
}
