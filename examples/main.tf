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

resource "neptune_project" "example_project" {
    name = "Terraform"
    workspace = "aleksander.wojnarowicz"
    key = "TER"
    vis = "priv"
}