package schemas

import (
	"github.com/hashicorp/terraform/helper/schema"
)

type Organization struct {
	Id   string
	Name string
}

func LoadOrganization(d *schema.ResourceData) (sch *Organization) {
	sch = &Organization{
		Id:   d.Id(),
		Name: d.Get("name").(string),
	}

	return
}
