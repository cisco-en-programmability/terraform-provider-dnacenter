terraform {
  required_providers {
    dnacenter = {
      version = "1.3.1-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}



data "dnacenter_application_policy_application_set" "example" {
  provider = dnacenter
  #  attributes = "APPLICATION_GROUP"
  # limit      = 1
  # name       = "string"
  # offset     = 1
}

output "dnacenter_application_policy_application_set_example" {
  value = data.dnacenter_application_policy_application_set.example.items
}
