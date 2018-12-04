package resources

import (
	"fmt"
	"github.com/dropbox/godropbox/errors"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/errortypes"
	"github.com/pritunl/terraform-provider-pritunl/request"
	"github.com/pritunl/terraform-provider-pritunl/schemas"
)

func Link() *schema.Resource {
	return &schema.Resource{
		Create: linkCreate,
		Read:   linkRead,
		Update: linkUpdate,
		Delete: linkDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

type linkPostData struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Ipv6   bool   `json:"ipv6"`
}

type linkPutData struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Ipv6   bool   `json:"ipv6"`
}

type linkData struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Ipv6   bool   `json:"ipv6"`
}

func linkGet(prvdr *schemas.Provider, sch *schemas.Link) (
	data *linkData, err error) {

	req := request.Request{
		Method: "GET",
		Path:   "/tf/link",
		Query: map[string]string{
			"id":   sch.Id,
			"name": sch.Name,
		},
	}

	data = &linkData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		data = nil
	}

	return
}

func linkPut(prvdr *schemas.Provider, sch *schemas.Link) (
	data *linkData, err error) {

	req := request.Request{
		Method: "PUT",
		Path:   fmt.Sprintf("/tf/link/%s", sch.Id),
		Json: &linkPutData{
			Name: sch.Name,
			Ipv6: sch.Ipv6,
		},
	}

	data = &linkData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		data = nil
	}

	return
}

func linkPost(prvdr *schemas.Provider, sch *schemas.Link) (
	data *linkData, err error) {

	req := request.Request{
		Method: "POST",
		Path:   "/tf/link",
		Json: &linkPostData{
			Name: sch.Name,
			Ipv6: sch.Ipv6,
		},
	}

	data = &linkData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		err = &errortypes.RequestError{
			errors.New("link: Not found on post"),
		}
		return
	}

	return
}

func linkDel(prvdr *schemas.Provider, sch *schemas.Link) (
	err error) {

	req := request.Request{
		Method: "DELETE",
		Path:   fmt.Sprintf("/tf/link/%s", sch.Id),
	}

	_, err = req.Do(prvdr, nil)
	if err != nil {
		return
	}

	return
}

func linkCreate(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadLink(d)

	data, err := linkGet(prvdr, sch)
	if err != nil {
		return
	}

	if data != nil {
		sch.Id = data.Id

		data, err = linkPut(prvdr, sch)
		if err != nil {
			return
		}
	}

	if data == nil {
		data, err = linkPost(prvdr, sch)
		if err != nil {
			return
		}
	}

	d.SetId(data.Id)

	return
}

func linkUpdate(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadLink(d)

	data, err := linkPut(prvdr, sch)
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

func linkRead(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadLink(d)

	data, err := linkGet(prvdr, sch)
	if err != nil {
		return
	}

	if data == nil {
		d.SetId("")
		return
	}

	d.Set("name", data.Name)
	d.Set("ipv6", data.Ipv6)
	d.SetId(data.Id)

	return
}

func linkDelete(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadLink(d)

	err = linkDel(prvdr, sch)
	if err != nil {
		return
	}

	d.SetId("")

	return
}
