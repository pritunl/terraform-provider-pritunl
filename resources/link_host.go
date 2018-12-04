package resources

import (
	"fmt"
	"github.com/dropbox/godropbox/errors"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pritunl/terraform-provider-pritunl/errortypes"
	"github.com/pritunl/terraform-provider-pritunl/request"
	"github.com/pritunl/terraform-provider-pritunl/schemas"
)

func LinkHost() *schema.Resource {
	return &schema.Resource{
		Create: linkHostCreate,
		Read:   linkHostRead,
		Update: linkHostUpdate,
		Delete: linkHostDelete,
		Schema: map[string]*schema.Schema{
			"location_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"static": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"public_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"local_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"address6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uri": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

type linkHostPostData struct {
	Name          string `json:"name"`
	Timeout       int    `json:"timeout"`
	Priority      int    `json:"priority"`
	Static        bool   `json:"static"`
	PublicAddress string `json:"public_address"`
	LocalAddress  string `json:"local_address"`
	Address6      string `json:"address6"`
}

type linkHostPutData struct {
	Name          string `json:"name"`
	Timeout       int    `json:"timeout"`
	Priority      int    `json:"priority"`
	Static        bool   `json:"static"`
	PublicAddress string `json:"public_address"`
	LocalAddress  string `json:"local_address"`
	Address6      string `json:"address6"`
}

type linkHostData struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Timeout       int    `json:"timeout"`
	Priority      int    `json:"priority"`
	Static        bool   `json:"static"`
	PublicAddress string `json:"public_address"`
	LocalAddress  string `json:"local_address"`
	Address6      string `json:"address6"`
	Uri           string `json:"uri"`
}

func linkHostGet(prvdr *schemas.Provider, sch *schemas.LinkHost) (
	data *linkHostData, err error) {

	req := request.Request{
		Method: "GET",
		Path:   fmt.Sprintf("/tf/link/%s/host", sch.LocationId),
		Query: map[string]string{
			"id":   sch.Id,
			"name": sch.Name,
		},
	}

	data = &linkHostData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		data = nil
	}

	return
}

func linkHostPut(prvdr *schemas.Provider, sch *schemas.LinkHost) (
	data *linkHostData, err error) {

	req := request.Request{
		Method: "PUT",
		Path:   fmt.Sprintf("/tf/link/%s/host/%s", sch.LocationId, sch.Id),
		Json: &linkHostPutData{
			Name:          sch.Name,
			Timeout:       sch.Timeout,
			Priority:      sch.Priority,
			Static:        sch.Static,
			PublicAddress: sch.PublicAddress,
			LocalAddress:  sch.LocalAddress,
			Address6:      sch.Address6,
		},
	}

	data = &linkHostData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		data = nil
	}

	return
}

func linkHostPost(prvdr *schemas.Provider, sch *schemas.LinkHost) (
	data *linkHostData, err error) {

	req := request.Request{
		Method: "POST",
		Path:   fmt.Sprintf("/tf/link/%s/host", sch.LocationId),
		Json: &linkHostPostData{
			Name:          sch.Name,
			Timeout:       sch.Timeout,
			Priority:      sch.Priority,
			Static:        sch.Static,
			PublicAddress: sch.PublicAddress,
			LocalAddress:  sch.LocalAddress,
			Address6:      sch.Address6,
		},
	}

	data = &linkHostData{}

	resp, err := req.Do(prvdr, data)
	if err != nil {
		return
	}

	if resp.StatusCode == 404 {
		err = &errortypes.RequestError{
			errors.New("linkHost: Not found on post"),
		}
		return
	}

	return
}

func linkHostDel(prvdr *schemas.Provider, sch *schemas.LinkHost) (
	err error) {

	req := request.Request{
		Method: "DELETE",
		Path:   fmt.Sprintf("/tf/link/%s/host/%s", sch.LocationId, sch.Id),
	}

	_, err = req.Do(prvdr, nil)
	if err != nil {
		return
	}

	return
}

func linkHostCreate(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadLinkHost(d)

	data, err := linkHostGet(prvdr, sch)
	if err != nil {
		return
	}

	if data != nil {
		sch.Id = data.Id

		data, err = linkHostPut(prvdr, sch)
		if err != nil {
			return
		}
	}

	if data == nil {
		data, err = linkHostPost(prvdr, sch)
		if err != nil {
			return
		}
	}

	d.SetId(data.Id)

	return
}

func linkHostUpdate(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadLinkHost(d)

	data, err := linkHostPut(prvdr, sch)
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

func linkHostRead(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadLinkHost(d)

	data, err := linkHostGet(prvdr, sch)
	if err != nil {
		return
	}

	if data == nil {
		d.SetId("")
		return
	}



	d.Set("name", data.Name)
	d.Set("timeout", data.Timeout)
	d.Set("priority", data.Priority)
	d.Set("static", data.Static)
	d.Set("public_address", data.PublicAddress)
	d.Set("local_address", data.LocalAddress)
	d.Set("address6", data.Address6)
	d.Set("uri", data.Uri + prvdr.PritunlHost)
	d.SetId(data.Id)

	return
}

func linkHostDelete(d *schema.ResourceData, m interface{}) (err error) {
	prvdr := m.(*schemas.Provider)
	sch := schemas.LoadLinkHost(d)

	err = linkHostDel(prvdr, sch)
	if err != nil {
		return
	}

	d.SetId("")

	return
}
