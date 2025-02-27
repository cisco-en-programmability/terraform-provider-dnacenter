
terraform {
  required_providers {
    dnacenter = {
      version = "1.3.1-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}
data "dnacenter_lan_automation_log" "example" {
  provider = dnacenter
  //limit = 1
  //offset = 1
}

output "dnacenter_lan_automation_log_example" {
  value = data.dnacenter_lan_automation_log.example.items
}

/*data "dnacenter_lan_automation_log" "example" {
    provider = dnacenter
    //id = "string"
}

output "dnacenter_lan_automation_log_example" {
    value = data.dnacenter_lan_automation_log.example.item
}
*/
