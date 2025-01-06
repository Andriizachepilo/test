package cicd

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns a terraform.Provider for your plugin.
func Provider() *schema.Provider {
    return &schema.Provider{
        Schema: map[string]*schema.Schema{
            // Define configuration options for your provider
        },
        ResourcesMap: map[string]*schema.Resource{
            // Define resources here, e.g., "cicd_resource": resourceCICD(),
        },
        DataSourcesMap: map[string]*schema.Resource{
            // Define data sources here, e.g., "cicd_data_source": dataSourceCICD(),
        },
    }
}
