package schemas

import (
	"github.com/hashicorp/terraform/helper/schema"
)

type Provider struct {
	PritunlHost   string
	PritunlToken  string
	PritunlSecret string
}

func LoadProvider(d *schema.ResourceData) (sch *Provider) {
	sch = &Provider{
		PritunlHost:   d.Get("pritunl_host").(string),
		PritunlToken:  d.Get("pritunl_token").(string),
		PritunlSecret: d.Get("pritunl_secret").(string),
	}

	return
}
