terraform {
  required_providers {
    dnacenter = {
      version = "1.3.1-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}
data "dnacenter_discovery_count" "example" {
  provider = dnacenter
}

output "dnacenter_discovery_count_example" {
  value = data.dnacenter_discovery_count.example.item
}
