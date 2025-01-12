package main

import (
	"github.com/Andriizachepilo/test/cicd" // Replace with the actual module path
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cicd.Provider, // Registering the CICD provider
	})
}
