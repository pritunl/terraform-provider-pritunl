package schemas

import (
	"github.com/hashicorp/terraform/helper/schema"
)

type User struct {
	Id              string
	OrganizationId  string
	Name            string
	Email           string
	AuthType        string
	Groups          []string
	Pin             string
	Disabled        bool
	NetworkLinks    []string
	BypassSecondary bool
	ClientToClient  bool
	DnsServers      []string
	DnsSuffix       string
}

func LoadUser(d *schema.ResourceData) (sch *User) {
	sch = &User{
		Id:              d.Id(),
		OrganizationId:  d.Get("organization_id").(string),
		Name:            d.Get("name").(string),
		Email:           d.Get("email").(string),
		AuthType:        d.Get("auth_type").(string),
		Pin:             d.Get("pin").(string),
		Disabled:        d.Get("disabled").(bool),
		BypassSecondary: d.Get("bypass_secondary").(bool),
		ClientToClient:  d.Get("client_to_client").(bool),
		DnsSuffix:       d.Get("dns_suffix").(string),
	}

	groups := d.Get("groups").([]interface{})
	if groups != nil {
		sch.Groups = []string{}
		for _, group := range groups {
			sch.Groups = append(sch.Groups, group.(string))
		}
	}

	networkLinks := d.Get("network_links").([]interface{})
	if networkLinks != nil {
		sch.NetworkLinks = []string{}
		for _, networkLink := range networkLinks {
			sch.NetworkLinks = append(sch.NetworkLinks, networkLink.(string))
		}
	}

	dnsServers := d.Get("dns_servers").([]interface{})
	if dnsServers != nil {
		sch.DnsServers = []string{}
		for _, dnsServer := range dnsServers {
			sch.DnsServers = append(sch.DnsServers, dnsServer.(string))
		}
	}

	return
}
