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
		Schema: map[string]*schema.Schema{
			"api_token": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			// the name of the resource
			"pingdom_check": resourcePingdomCheck(),
		},
		DataSourcesMap: map[string]*schema.Resource{
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := PingdomConfig{
		APIToken: d.Get("api_token").(string),
	}
	client, err := config.getPingdomClient()
	if err != nil {
		return nil, err
	}

	return client, nil
}

type PingdomConfig struct {
	APIToken	string
}

func (c *PingdomConfig) getPingdomClient() (*pingdom.Client, error) {
	if v := os.Getenv("PINGDOM_API_TOKEN"); v != "" {
		c.APIToken = v
	}
	client, _ := pingdom.NewClientWithConfig(pingdom.ClientConfig{
		APIToken: c.APIToken,
	})

	log.Printf("[INFO] Pingdom Client configured.")

	return client, nil
}