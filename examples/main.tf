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
    name = "Terraform"
    workspace = "aleksander.wojnarowicz"
}