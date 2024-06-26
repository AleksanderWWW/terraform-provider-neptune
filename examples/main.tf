terraform {
  required_providers {
    neptune = {
      source = "registry.terraform.io/hashicorp/neptune"
    }
  }
}

provider "neptune" {}

resource "neptune_project" "my_project" {
  name      = "Terraform-project"
  workspace = "e2e-tests"
}

resource "neptune_project_member" "my_project_member" {
  depends_on = [neptune_project.my_project]
  project    = "Terraform-project"
  workspace  = "e2e-tests"
  username   = "e2e.regularuser"
  role       = "member"
}
