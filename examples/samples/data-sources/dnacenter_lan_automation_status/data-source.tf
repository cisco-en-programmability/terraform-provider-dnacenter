terraform {
  required_providers {
    dnacenter = {
      version = "1.1.27-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

data "dnacenter_lan_automation_status" "example" {
  provider = dnacenter
  //limit = 1
  //offset = 1
}

output "dnacenter_lan_automation_status_example" {
  value = data.dnacenter_lan_automation_status.example.items
}

/*data "dnacenter_lan_automation_status" "example" {
    provider = dnacenter
    id = "string"
}

output "dnacenter_lan_automation_status_example" {
    value = data.dnacenter_lan_automation_status.example.item
}
*/
