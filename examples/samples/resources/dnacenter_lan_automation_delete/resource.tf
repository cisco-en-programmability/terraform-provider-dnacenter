terraform {
  required_providers {
    dnacenter = {
      version = "1.0.17-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}
resource "dnacenter_lan_automation_delete" "example" {
  provider = dnacenter
  parameters {

    id = "string"
  }
}

output "dnacenter_lan_automation_delete_example" {
  value = dnacenter_lan_automation_delete.example
}