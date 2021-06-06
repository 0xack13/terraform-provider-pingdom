package pingdom

import (
	"log"
	"os"

	// "github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/russellcardullo/go-pingdom/pingdom"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		// schema contains the APITOKEN that will be sent to Pingdmo
		Schema: map[string]*schema.Schema{
			"api_token": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			// the name of the resource in the tf file
			"pingdom_check": resourcePingdomCheck(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
		// configurefunc is used to return  a client
		ConfigureFunc: providerConfigure,
	}
}

type PingdomConfig struct {
	APIToken	string
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	// struct PingdomConfig
	config := PingdomConfig{
		APIToken: d.Get("api_token").(string),
	}
	client, err := config.getPingdomClient()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *PingdomConfig) getPingdomClient() (*pingdom.Client, error) {
	// checks the env variable
	if v := os.Getenv("PINGDOM_API_TOKEN"); v != "" {
		c.APIToken = v
	}
	// Create a new Pingdom client with the API token
	client, _ := pingdom.NewClientWithConfig(pingdom.ClientConfig{
		APIToken: c.APIToken,
	})

	log.Printf("[INFO] Pingdom Client configured.")

	return client, nil
}