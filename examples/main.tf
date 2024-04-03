terraform {
  required_providers {
    neptune = {
      source = "registry.terraform.io/hashicorp/neptune"
    }
  }
}

provider "neptune" {
  api_token = ""
}

resource "neptune_project" "my_project" {
  name      = "Terraform-project"
  workspace = "e2e"
  key       = "TEST"
}

resource "neptune_project_member" "my_project_member" {
  project   = "Terraform-user"
  workspace = "e2e"
  username  = "team-member"
  role      = "viewer"
}
