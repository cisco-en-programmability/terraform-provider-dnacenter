
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.9-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_application_sets" "example" {
  provider = dnacenter
  parameters {
    payload {
      name = "TestAppTerraform222Update"
    }
  }
}

output "dnacenter_application_sets_example" {
  value = dnacenter_application_sets.example
}
