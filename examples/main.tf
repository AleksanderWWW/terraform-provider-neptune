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
  name      = "Terraform"
  workspace = "aleksander.wojnarowicz"
}

resource "neptune_project_member" "my_project_member" {
  project   = "e2e-20231130-0854-mouth-southern-trip"
  workspace = "e2e-tests"
  username  = "e2e.regularuser"
  role      = "member"
}
