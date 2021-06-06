package pingdom

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/russellcardullo/go-pingdom/pingdom"
)

// resource function -> returns Resource
func resourcePingdomCheck() *schema.Resource {
	// imlementing functions and providing the required schema for the resource
	return &schema.Resource{
		Create: resourcePingdomCheckCreate,
		Read: resourcePingdomCheckRead,
		Update: resourcePingdomCheckUpdate,
		Delete: resourcePingdomCheckDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			"paused": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
			},

	
			"resolution": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: false,
			},

	
			"url": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
				Default:  "/",
			},

			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: false,
				Computed: true,
			},
		},
	}
}

func resourcePingdomCheckCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pingdom.Client)
	check := &pingdom.HttpCheck{
		Name: d.Get("name").(string),
		Resolution: d.Get("resolution").(int),
		Paused: d.Get("paused").(bool),
		Hostname: d.Get("url").(string),
		Port: d.Get("port").(int),
	}
	resp, _ := client.Checks.Create(check)
	d.SetId(fmt.Sprintf("%d", resp.ID))
	return resourcePingdomCheckRead(d, meta)
}

func resourcePingdomCheckRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pingdom.Client)
	checkId, _ := strconv.Atoi(d.Id())
	check, _ := client.Checks.Read(checkId)
	d.Set("name", check.Name)
	d.Set("resolution", check.Resolution)
	d.Set("port", check.Type.HTTP.Port)
	d.Set("paused", check.Paused)
	d.Set("url", check.Hostname)
	return nil
}

func resourcePingdomCheckDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pingdom.Client)
	// Atoi = convert string/ascii to int
	id, _ := strconv.Atoi(d.Id())
	log.Printf("[INFO] Deleting pingdom Check: %v", id)
	_, err := client.Checks.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting check: %s", err)
	}
	return nil
}

func resourcePingdomCheckUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}
