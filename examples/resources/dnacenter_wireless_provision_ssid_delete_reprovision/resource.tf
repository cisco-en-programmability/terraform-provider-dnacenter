
resource "dnacenter_wireless_provision_ssid_delete_reprovision" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    managed_aplocations = "string"
    ssid_name           = "string"
  }
}