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

data "dnacenter_event_subscription" "example" {
  provider  = dnacenter
  event_ids = "string"
  limit     = 1
  offset    = 1
  order     = "string"
  sort_by   = "string"
}

output "dnacenter_event_subscription_example" {
  value = data.dnacenter_event_subscription.example.items
}
