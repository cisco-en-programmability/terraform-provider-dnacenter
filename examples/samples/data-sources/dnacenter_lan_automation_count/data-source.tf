terraform {
  required_providers {
    dnacenter = {
      version = "1.0.12-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

data "dnacenter_lan_automation_count" "example" {
  provider = dnacenter
}

output "dnacenter_lan_automation_count_example" {
  value = data.dnacenter_lan_automation_count.example.item
}
