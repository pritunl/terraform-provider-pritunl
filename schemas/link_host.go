package schemas

import (
	"github.com/hashicorp/terraform/helper/schema"
)

type LinkHost struct {
	Id            string
	LocationId    string
	Name          string
	Timeout       int
	Priority      int
	Static        bool
	PublicAddress string
	LocalAddress  string
	Address6      string
}

func LoadLinkHost(d *schema.ResourceData) (sch *LinkHost) {
	sch = &LinkHost{
		Id:            d.Id(),
		LocationId:    d.Get("location_id").(string),
		Name:          d.Get("name").(string),
		Timeout:       d.Get("timeout").(int),
		Priority:      d.Get("priority").(int),
		Static:        d.Get("static").(bool),
		PublicAddress: d.Get("public_address").(string),
		LocalAddress:  d.Get("local_address").(string),
		Address6:      d.Get("address6").(string),
	}

	return
}
