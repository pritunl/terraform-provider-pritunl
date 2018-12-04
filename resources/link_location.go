package resources

import (
	"fmt"
	"github.com/dropbox/godropbox/errors"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/errortypes"
	"github.com/pritunl/terraform-provider-pritunl/request"
	"github.com/pritunl/terraform-provider-pritunl/schemas"
)

func LinkLocation() *schema.Resource {
	return &schema.Resource{
		Create: linkLocationCreate,
		Read:   linkLocationRead,
		Update: linkLocationUpdate,
		Delete: linkLocationDelete,
		Schema: map[string]*schema.Schema{
			"link_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

type linkLocationPostData struct {
	Name string `json:"name"`
}

type linkLocationPutData struct {
	Name string `json:"name"`
}

type linkLocationData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func linkLocationGet(prvdr *schemas.Provider, sch *schemas.LinkLocation) (
	data *linkLocationData, err error) {

	req := request.Request{
		Method: "GET",
		Path:   fmt.Sprintf("/tf/link/%s/location", sch.LinkId),
		Query: map[string]string{
			"id":   sch.Id,
			"name": sch.Name,
		},
	}

	data = &linkLocationData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		data = nil
	}

	return
}

func linkLocationPut(prvdr *schemas.Provider, sch *schemas.LinkLocation) (
	data *linkLocationData, err error) {

	req := request.Request{
		Method: "PUT",
		Path:   fmt.Sprintf("/tf/link/%s/location/%s", sch.LinkId, sch.Id),
		Json: &linkLocationPutData{
			Name: sch.Name,
		},
	}

	data = &linkLocationData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		data = nil
	}

	return
}

func linkLocationPost(prvdr *schemas.Provider, sch *schemas.LinkLocation) (
	data *linkLocationData, err error) {

	req := request.Request{
		Method: "POST",
		Path:   fmt.Sprintf("/tf/link/%s/location", sch.LinkId),
		Json: &linkLocationPostData{
			Name: sch.Name,
		},
	}

	data = &linkLocationData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		err = &errortypes.RequestError{
			errors.New("linkLocation: Not found on post"),
		}
		return
	}

	return
}

func linkLocationDel(prvdr *schemas.Provider, sch *schemas.LinkLocation) (
	err error) {

	req := request.Request{
		Method: "DELETE",
		Path:   fmt.Sprintf("/tf/link/%s/location/%s", sch.LinkId, sch.Id),
	}

	_, err = req.Do(prvdr, nil)
	if err != nil {
		return
	}

	return
}

func linkLocationCreate(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadLinkLocation(d)

	data, err := linkLocationGet(prvdr, sch)
	if err != nil {
		return
	}

	if data != nil {
		sch.Id = data.Id

		data, err = linkLocationPut(prvdr, sch)
		if err != nil {
			return
		}
	}

	if data == nil {
		data, err = linkLocationPost(prvdr, sch)
		if err != nil {
			return
		}
	}

	d.SetId(data.Id)

	return
}

func linkLocationUpdate(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadLinkLocation(d)

	data, err := linkLocationPut(prvdr, sch)
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

func linkLocationRead(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadLinkLocation(d)

	data, err := linkLocationGet(prvdr, sch)
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

func linkLocationDelete(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadLinkLocation(d)

	err = linkLocationDel(prvdr, sch)
	if err != nil {
		return
	}

	d.SetId("")

	return
}
