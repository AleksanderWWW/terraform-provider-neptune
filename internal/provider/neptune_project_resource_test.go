// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNeptuneProjectResource(t *testing.T) {
	projectName := "Terraform-" + fake.Hash().MD5()
	key := strings.ToUpper(fake.Lorem().Text(3))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + fmt.Sprintf(`
				resource "neptune_project" "test" {
					name      = "%s"
					workspace = "e2e-tests"
					key 	  = "%s"
				}
				`, projectName, key),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("neptune_project.test", "name", projectName),
					resource.TestCheckResourceAttr("neptune_project.test", "workspace", "e2e-tests"),
				),
			},
			// ImportState testing (not applicable)
			// {
			// 	ResourceName:      "neptune_project.test",
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// 	// This is not normally necessary, but is here because this
			// 	// example code does not have an actual upstream service.
			// 	// Once the Read method is able to refresh information from
			// 	// the upstream service, this can be removed.
			// },
			// Update and Read testing (not applicable)
			// {
			// 	Config: testAccExampleResourceConfig("two"),
			// 	Check: resource.ComposeAggregateTestCheckFunc(
			// 		resource.TestCheckResourceAttr("scaffolding_example.test", "configurable_attribute", "two"),
			// 	),
			// },
			// Delete testing automatically occurs in TestCase
		},
	})
}
