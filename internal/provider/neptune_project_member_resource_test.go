// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNeptuneProjectMemberResource(t *testing.T) {
	projectName := "Terraform-" + fake.Hash().MD5()
	key := strings.ToUpper(fake.Lorem().Text(10))

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

				resource "neptune_project_member" "test_member" {
					depends_on = [neptune_project.test]
					project   = "Terraform-user"
					workspace = "e2e-tests"
					username  = "e2e.regularuser"
					role      = "member"
				  }
				`, projectName, key),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("neptune_project_member.test_member", "project", "Terraform-user"),
					resource.TestCheckResourceAttr("neptune_project_member.test_member", "workspace", "e2e-tests"),
					resource.TestCheckResourceAttr("neptune_project_member.test_member", "username", "e2e.regularuser"),
					resource.TestCheckResourceAttr("neptune_project_member.test_member", "role", "member"),
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
			// Update testing
			{
				Config: providerConfig + `
				resource "neptune_project_member" "test_member" {
					project   = "Terraform-user"
					workspace = "e2e-tests"
					username  = "e2e.regularuser"
					role      = "viewer"
				  }
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("neptune_project_member.test_member", "project", "Terraform-user"),
					resource.TestCheckResourceAttr("neptune_project_member.test_member", "workspace", "e2e-tests"),
					resource.TestCheckResourceAttr("neptune_project_member.test_member", "username", "e2e.regularuser"),
					resource.TestCheckResourceAttr("neptune_project_member.test_member", "role", "viewer"),
				),
			},
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
