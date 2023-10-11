terraform {
  required_providers {
    dnacenter = {
      version = "1.1.20-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

data "dnacenter_app_policy_default" "example" {
  provider = dnacenter
}

output "dnacenter_app_policy_default_example" {
  value = data.dnacenter_app_policy_default.example.items
}
