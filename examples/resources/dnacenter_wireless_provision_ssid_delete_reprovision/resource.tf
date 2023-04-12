
resource "dnacenter_wireless_provision_ssid_delete_reprovision" "example" {
  provider = dnacenter
  parameters {

    managed_aplocations = "string"
    ssid_name           = "string"
    persistbapioutput   = "false"
  }
}

output "dnacenter_wireless_provision_ssid_delete_reprovision_example" {
  value = dnacenter_wireless_provision_ssid_delete_reprovision.example
}