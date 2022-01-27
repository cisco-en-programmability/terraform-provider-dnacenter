
data "dnacenter_wireless_provision_ssid_delete_reprovision" "example" {
  provider            = dnacenter
  managed_aplocations = "string"
  ssid_name           = "string"
}