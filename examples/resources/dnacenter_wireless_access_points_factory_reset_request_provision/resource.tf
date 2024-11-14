
resource "dnacenter_wireless_access_points_factory_reset_request_provision" "example" {
  provider = dnacenter
  parameters {

    ap_mac_addresses     = ["string"]
    keep_static_ipconfig = "false"
  }
}

output "dnacenter_wireless_access_points_factory_reset_request_provision_example" {
  value = dnacenter_wireless_access_points_factory_reset_request_provision.example
}