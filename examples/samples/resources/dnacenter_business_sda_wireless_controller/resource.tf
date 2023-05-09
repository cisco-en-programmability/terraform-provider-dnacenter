
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.5-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_business_sda_wireless_controller" "example" {
  provider = dnacenter
  parameters {
    device_name         = "C9K-Branch-SFO.dcloud.cisco.com"
    site_name_hierarchy = "Global/San Francisco"
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
