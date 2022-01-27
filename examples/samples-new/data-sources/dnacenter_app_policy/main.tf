terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

data "dnacenter_app_policy" "example" {
  provider     = dnacenter
  policy_scope = "string"
}

output "dnacenter_app_policy_example" {
  value = data.dnacenter_app_policy.example.items
}