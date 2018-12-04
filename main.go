package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"github.com/pritunl/terraform-provider-pritunl/provider"
	"github.com/pritunl/terraform-provider-pritunl/utils"
)

func main() {
	utils.OutputOpen()
	defer utils.OutputClose()

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return provider.Provider()
		},
	})
}
