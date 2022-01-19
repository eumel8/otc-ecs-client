package main

import (
	"flag"
	"fmt"
	gophercloud "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/rds/v3/instances"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v3"
)

func rdsCreate(client *gophercloud.ServiceClient, opts *servers.ListOpts) {

	createOpts := instances.CreateRdsOpts{
	createResult := instances.Create(client, createOpts)
	r, err := createResult.Extract()
	if err != nil {
		panic(err)
	}

}

func readYaml{

     yfile, err := ioutil.ReadFile("mydb.yaml")

     if err != nil {

          log.Fatal(err)
     }

     data := make(map[interface{}]interface{})

     err2 := yaml.Unmarshal(yfile, &data)

     if err2 != nil {

          log.Fatal(err2)
     }

     for k, v := range data {

          fmt.Printf("%s -> %d\n", k, v)
     }
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

	createOpts := instances.CreateRdsOpts{
		Name:             d.Get("name").(string),
		Datastore:        resourceRDSDataStore(d),
		Ha:               resourceRDSHa(d),
		ConfigurationId:  d.Get("param_group_id").(string),
		Port:             dbPortString,
		Password:         dbInfo["password"].(string),
		BackupStrategy:   resourceRDSBackupStrategy(d),
		DiskEncryptionId: volumeInfo["disk_encryption_id"].(string),
		FlavorRef:        d.Get("flavor").(string),
		Volume:           resourceRDSVolume(d),
		Region:           config.GetRegion(d),
		AvailabilityZone: resourceRDSAvailabilityZones(d),
		VpcId:            d.Get("vpc_id").(string),
		SubnetId:         d.Get("subnet_id").(string),
		SecurityGroupId:  d.Get("security_group_id").(string),
		ChargeInfo:       resourceRDSChangeMode(),
	}


	rds, err := rdsCreate(provider, rdsOptions)
	if err != nil {
		panic(err)
	}
	ecsList(ecs, &servers.ListOpts{Status: *status})
}
