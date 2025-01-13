package main

import (
    "github.com/hashicorp/terraform/plugin"
    "github.com/hashicorp/terraform/terraform"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: func() terraform.ResourceProvider {
            return &schema.Provider{
                ResourcesMap: map[string]*schema.Resource{
                    "cicd_pipeline": resourceCICDPipeline(),
                },
            }
        },
    })
}
