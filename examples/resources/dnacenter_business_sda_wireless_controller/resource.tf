
provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_business_sda_wireless_controller" "example" {
  provider = dnacenter
  parameters {
    device_name         = "string"
    site_name_hierarchy = "string"
  }
}

output "dnacenter_business_sda_wireless_controller_example" {
  value = dnacenter_business_sda_wireless_controller.example
}

data "dnacenter_dnacaap_management_execution_status" "example" {
  depends_on   = [dnacenter_business_sda_wireless_controller.example]
  provider     = dnacenter
  execution_id = dnacenter_business_sda_wireless_controller.example.item.0.execution_id
}

output "dnacenter_dnacaap_management_execution_status_example" {
  value = data.dnacenter_dnacaap_management_execution_status.example.item
}
