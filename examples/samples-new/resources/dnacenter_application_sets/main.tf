
terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

resource "dnacenter_application_sets" "example" {
  provider = dnacenter
  parameters {

    name = "TestApp"
  }
}

output "dnacenter_application_sets_example" {
  value = dnacenter_application_sets.example
}
