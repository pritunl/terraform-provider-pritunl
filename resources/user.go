package resources

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/request"
	"github.com/pritunl/terraform-provider-pritunl/schemas"
	"github.com/pritunl/terraform-provider-pritunl/errortypes"
	"github.com/dropbox/godropbox/errors"
)

func User() *schema.Resource {
	return &schema.Resource{
		Create: userCreate,
		Read:   userRead,
		Update: userUpdate,
		Delete: userDelete,
		Schema: map[string]*schema.Schema{
			"organization_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"pin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"disabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"network_links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"bypass_secondary": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"client_to_client": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"dns_servers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"dns_suffix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

type userPostData struct {
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	AuthType        string   `json:"auth_type"`
	Groups          []string `json:"groups"`
	Pin             string   `json:"pin"`
	Disabled        bool     `json:"disabled"`
	NetworkLinks    []string `json:"network_links"`
	BypassSecondary bool     `json:"bypass_secondary"`
	ClientToClient  bool     `json:"client_to_client"`
	DnsServers      []string `json:"dns_servers"`
	DnsSuffix       string   `json:"dns_suffix"`
}

type userPutData struct {
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	AuthType        string   `json:"auth_type"`
	Groups          []string `json:"groups"`
	Pin             string   `json:"pin"`
	Disabled        bool     `json:"disabled"`
	NetworkLinks    []string `json:"network_links"`
	BypassSecondary bool     `json:"bypass_secondary"`
	ClientToClient  bool     `json:"client_to_client"`
	DnsServers      []string `json:"dns_servers"`
	DnsSuffix       string   `json:"dns_suffix"`
}

type userData struct {
	Id              string   `json:"id"`
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	AuthType        string   `json:"auth_type"`
	Groups          []string `json:"groups"`
	Disabled        bool     `json:"disabled"`
	NetworkLinks    []string `json:"network_links"`
	BypassSecondary bool     `json:"bypass_secondary"`
	ClientToClient  bool     `json:"client_to_client"`
	DnsServers      []string `json:"dns_servers"`
	DnsSuffix       string   `json:"dns_suffix"`
}

func userGet(prvdr *schemas.Provider, sch *schemas.User) (
	data *userData, err error) {

	req := request.Request{
		Method: "GET",
		Path:   fmt.Sprintf("/tf/user/%s", sch.OrganizationId),
		Query: map[string]string{
			"id":   sch.Id,
			"name": sch.Name,
		},
	}

	data = &userData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		data = nil
	}

	return
}

func userPut(prvdr *schemas.Provider, sch *schemas.User) (
	data *userData, err error) {

	req := request.Request{
		Method: "PUT",
		Path:   fmt.Sprintf("/tf/user/%s/%s", sch.OrganizationId, sch.Id),
		Json: &userPutData{
			Name:            sch.Name,
			Email:           sch.Email,
			AuthType:        sch.AuthType,
			Groups:          sch.Groups,
			Pin:             sch.Pin,
			Disabled:        sch.Disabled,
			NetworkLinks:    sch.NetworkLinks,
			BypassSecondary: sch.BypassSecondary,
			ClientToClient:  sch.ClientToClient,
			DnsServers:      sch.DnsServers,
			DnsSuffix:       sch.DnsSuffix,
		},
	}

	data = &userData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		data = nil
	}

	return
}

func userPost(prvdr *schemas.Provider, sch *schemas.User) (
	data *userData, err error) {

	req := request.Request{
		Method: "POST",
		Path:   fmt.Sprintf("/tf/user/%s", sch.OrganizationId),
		Json: &userPostData{
			Name:            sch.Name,
			Email:           sch.Email,
			AuthType:        sch.AuthType,
			Groups:          sch.Groups,
			Pin:             sch.Pin,
			Disabled:        sch.Disabled,
			NetworkLinks:    sch.NetworkLinks,
			BypassSecondary: sch.BypassSecondary,
			ClientToClient:  sch.ClientToClient,
			DnsServers:      sch.DnsServers,
			DnsSuffix:       sch.DnsSuffix,
		},
	}

	data = &userData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		err = &errortypes.RequestError{
			errors.New("user: Not found on post"),
		}
		return
	}

	return
}

func userDel(prvdr *schemas.Provider, sch *schemas.User) (
	err error) {

	req := request.Request{
		Method: "DELETE",
		Path:   fmt.Sprintf("/tf/user/%s/%s", sch.OrganizationId, sch.Id),
	}

	_, err = req.Do(prvdr, nil)
	if err != nil {
		return
	}

	return
}

func userCreate(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadUser(d)

	data, err := userGet(prvdr, sch)
	if err != nil {
		return
	}

	if data != nil {
		sch.Id = data.Id

		data, err = userPut(prvdr, sch)
		if err != nil {
			return
		}
	}

	if data == nil {
		data, err = userPost(prvdr, sch)
		if err != nil {
			return
		}
	}

	d.SetId(data.Id)

	return
}

func userUpdate(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadUser(d)

	data, err := userPut(prvdr, sch)
	if err != nil {
		return
	}

	if data == nil {
		d.SetId("")
		return
	}

	d.SetId(data.Id)

	return
}

func userRead(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadUser(d)

	data, err := userGet(prvdr, sch)
	if err != nil {
		return
	}

	if data == nil {
		d.SetId("")
		return
	}

	d.Set("name", data.Name)
	d.Set("email", data.Email)
	d.Set("auth_type", data.AuthType)
	d.Set("groups", data.Groups)
	d.Set("disabled", data.Disabled)
	d.Set("network_links", data.NetworkLinks)
	d.Set("bypass_secondary", data.BypassSecondary)
	d.Set("client_to_client", data.ClientToClient)
	d.Set("dns_servers", data.DnsServers)
	d.Set("dns_suffix", data.DnsSuffix)
	d.SetId(data.Id)

	return
}

func userDelete(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadUser(d)

	err = userDel(prvdr, sch)
	if err != nil {
		return
	}

	d.SetId("")

	return
}
