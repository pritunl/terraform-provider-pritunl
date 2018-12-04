package resources

import (
	"fmt"
	"github.com/dropbox/godropbox/errors"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/errortypes"
	"github.com/pritunl/terraform-provider-pritunl/request"
	"github.com/pritunl/terraform-provider-pritunl/schemas"
)

func Organization() *schema.Resource {
	return &schema.Resource{
		Create: organizationCreate,
		Read:   organizationRead,
		Update: organizationUpdate,
		Delete: organizationDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

type organizationPostData struct {
	Name string `json:"name"`
}

type organizationPutData struct {
	Name string `json:"name"`
}

type organizationData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func organizationGet(prvdr *schemas.Provider, sch *schemas.Organization) (
	data *organizationData, err error) {

	req := request.Request{
		Method: "GET",
		Path:   "/tf/organization",
		Query: map[string]string{
			"id":   sch.Id,
			"name": sch.Name,
		},
	}

	data = &organizationData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		data = nil
	}

	return
}

func organizationPut(prvdr *schemas.Provider, sch *schemas.Organization) (
	data *organizationData, err error) {

	req := request.Request{
		Method: "PUT",
		Path:   fmt.Sprintf("/tf/organization/%s", sch.Id),
		Json: &organizationPutData{
			Name: sch.Name,
		},
	}

	data = &organizationData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		data = nil
	}

	return
}

func organizationPost(prvdr *schemas.Provider, sch *schemas.Organization) (
	data *organizationData, err error) {

	req := request.Request{
		Method: "POST",
		Path:   "/tf/organization",
		Json: &organizationPostData{
			Name: sch.Name,
		},
	}

	data = &organizationData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		err = &errortypes.RequestError{
			errors.New("organization: Not found on post"),
		}
		return
	}

	return
}

func organizationDel(prvdr *schemas.Provider, sch *schemas.Organization) (
	err error) {

	req := request.Request{
		Method: "DELETE",
		Path:   fmt.Sprintf("/tf/organization/%s", sch.Id),
	}

	_, err = req.Do(prvdr, nil)
	if err != nil {
		return
	}

	return
}

func organizationCreate(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadOrganization(d)

	data, err := organizationGet(prvdr, sch)
	if err != nil {
		return
	}

	if data != nil {
		sch.Id = data.Id

		data, err = organizationPut(prvdr, sch)
		if err != nil {
			return
		}
	}

	if data == nil {
		data, err = organizationPost(prvdr, sch)
		if err != nil {
			return
		}
	}

	d.SetId(data.Id)

	return
}

func organizationUpdate(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadOrganization(d)

	data, err := organizationPut(prvdr, sch)
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

func organizationRead(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadOrganization(d)

	data, err := organizationGet(prvdr, sch)
	if err != nil {
		return
	}

	if data == nil {
		d.SetId("")
		return
	}

	d.Set("name", data.Name)
	d.SetId(data.Id)

	return
}

func organizationDelete(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadOrganization(d)

	err = organizationDel(prvdr, sch)
	if err != nil {
		return
	}

	d.SetId("")

	return
}
