terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

data "dnacenter_app_policy_queuing_profile" "example" {
  provider = dnacenter
  name     = "CVD_QUEUING_PROFILE"
}

output "dnacenter_app_policy_queuing_profile_example" {
  value = data.dnacenter_app_policy_queuing_profile.example
}
