package main

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
    "github.com/Andriizachepilo/test/cicd" // Import cicd module
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: cicd.Provider, // Referencing cicd.Provider
    })
}
