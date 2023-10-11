terraform {
  required_providers {
    dnacenter = {
      version = "1.1.22-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

data "dnacenter_event_config_connector_types" "example" {
  provider = dnacenter
}

output "dnacenter_event_config_connector_types_example" {
  value = data.dnacenter_event_config_connector_types.example.items
}
