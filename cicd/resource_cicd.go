package cicd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCICDBuild() *schema.Resource {
	return &schema.Resource{
		Create: resourceCICDBuildCreate,
		Read:   resourceCICDBuildRead,
		Delete: resourceCICDBuildDelete,

		Schema: map[string]*schema.Schema{
			"build_tool": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_dir": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCICDBuildCreate(d *schema.ResourceData, m interface{}) error {
	buildTool := d.Get("build_tool").(string)
	sourceDir := d.Get("source_dir").(string)

	// Validate the build tool
	if buildTool != "maven" {
		return fmt.Errorf("unsupported build tool: %s", buildTool)
	}

	// Run the build command
	cmd := exec.Command("mvn", "clean", "package")
	cmd.Dir = sourceDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("build failed: %w", err)
	}

	// Mark the resource as created
	d.SetId(sourceDir)
	return nil
}

func resourceCICDBuildRead(d *schema.ResourceData, m interface{}) error {
	// Read is a no-op for this resource
	return nil
}

func resourceCICDBuildDelete(d *schema.ResourceData, m interface{}) error {
	// Deletion is a no-op for this resource
	d.SetId("")
	return nil
}
