package main

import (
	"flag"
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
	ColorWhite = "\x1b[37m"
	AppVersion = "0.0.2"
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
	t.AppendFooter(table.Row{white(tenant), white(region), white(project)})
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

func white(s string) string {
	return fmt.Sprintf("%s%s%s", ColorWhite, s, ColorDefault)
}

func main() {

	status := flag.String("status", "ACTIVE", "ecs status (ACTIVE|SHUTOFF|FAILURE)")
	version := flag.Bool("version", false, "app version")
	help := flag.Bool("help", false, "print out the help")

	flag.Parse()

	if *help {
		fmt.Println("Provide ENV variable to connect OTC: OS_PROJECT_NAME, OS_REGION_NAME, OS_AUTH_URL, OS_IDENTITY_API_VERSION, OS_USER_DOMAIN_NAME, OS_USERNAME, OS_PASSWORD")
		os.Exit(0)
	}

	if *version {
		fmt.Println("version", AppVersion)
		os.Exit(0)
	}

	if os.Getenv("OS_AUTH_URL") == "" {
		os.Setenv("OS_AUTH_URL", "https://iam.eu-de.otc.t-systems.com:443/v3")
	}

	if os.Getenv("OS_IDENTITY_API_VERSION") == "" {
		os.Setenv("OS_IDENTITY_API_VERSION", "3")
	}

	if os.Getenv("OS_REGION_NAME") == "" {
		os.Setenv("OS_REGION_NAME", "eu-de")
	}

	if os.Getenv("OS_PROJECT_NAME") == "" {
		os.Setenv("OS_PROJECT_NAME", "eu-de")
	}

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
	ecsList(ecs, &servers.ListOpts{Status: *status})
}
