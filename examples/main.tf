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
  project   = "e2e-20231207-0911-television-analysis"
  workspace = "e2e-tests"
  username  = "e2e.regularuser"
  role      = "member"
}
