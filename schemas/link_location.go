package schemas

import (
	"github.com/hashicorp/terraform/helper/schema"
)

type LinkLocation struct {
	Id     string
	LinkId string
	Name   string
}

func LoadLinkLocation(d *schema.ResourceData) (sch *LinkLocation) {
	sch = &LinkLocation{
		Id:     d.Id(),
		LinkId: d.Get("link_id").(string),
		Name:   d.Get("name").(string),
	}

	return
}
