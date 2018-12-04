package schemas

import (
	"github.com/hashicorp/terraform/helper/schema"
)

type Link struct {
	Id   string
	Name string
	Ipv6 bool
}

func LoadLink(d *schema.ResourceData) (sch *Link) {
	sch = &Link{
		Id:   d.Id(),
		Name: d.Get("name").(string),
		Ipv6: d.Get("ipv6").(bool),
	}

	return
}
