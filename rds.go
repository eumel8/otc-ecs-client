package main

import (
	"flag"
	"fmt"
	gophercloud "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/instances"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

const (
	AppVersion = "0.0.1"
)

type conf struct {
	Name             string `yaml:"name"`
	Datastore        string `yaml:"datastore"`
	Ha               string `yaml:"ha"`
	Port             string `yaml:"port"`
	Password         string `yaml:"passord"`
	BackupStrategy   string `yaml:"backupstrategy"`
	FlavorRef        string `yaml:"flavorref"`
	Volume           string `yaml:"volume"`
	Region           string `yaml:"region"`
	AvailabilityZone string `yaml:"availabilityzone"`
	VpcId            string `yaml:"vpcid"`
	SubnetId         string `yaml:"subnetid"`
	SecurityGroupId  string `yaml:"securitygroupid"`
}

func rdsCreate(client *gophercloud.ServiceClient, opts *c) r {

	createOpts := instances.CreateRdsOpts{
		Name:             opts.Name,
		Datastore:        opts.Datastore,
		Ha:               opts.Ha,
		Port:             opts.Port,
		Password:         opts.Password,
		BackupStrategy:   opts.BackupStrategy,
		FlavorRef:        opts.FlavorRef,
		Volume:           opts.Volume,
		Region:           opts.Region,
		AvailabilityZone: opts.AvailabilityZone,
		VpcId:            opts.VpcId,
		SubnetId:         opts.SubnetId,
		SecurityGroupId:  opts.SecurityGroupId,
	}
	createResult := instances.Create(client, createOpts)
	b, err := createResult.Extract()
	if err != nil {
		panic(r.Err)
	}
	return
}

func (c *conf) getConf() *conf {

	yfile, err := ioutil.ReadFile("mydb.yaml")

	if err != nil {
		panic(err)
	}

	// data := make(map[interface{}]interface{})
	err = yaml.Unmarshal(yfile, c)

	if err != nil {
		panic(err)
	}

	return c
}

func main() {

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

	var c conf
	c.getConf()

	rds, err := rdsCreate(provider, c)
	if err != nil {
		panic(err)
	}
}
