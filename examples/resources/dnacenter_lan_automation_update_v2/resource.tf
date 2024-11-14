
resource "dnacenter_lan_automation_update_v2" "example" {
  provider = dnacenter
  id       = "string"
  parameters {

    device_management_ipaddress = "string"
    new_host_name               = "string"
    new_loopback0_ipaddress     = "string"
  }
}

output "dnacenter_lan_automation_update_v2_example" {
  value = dnacenter_lan_automation_update_v2.example
}