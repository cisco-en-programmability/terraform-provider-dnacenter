
resource "dnacenter_wireless_access_points_provision" "example" {
  provider = dnacenter
  parameters {

    ap_zone_name = "string"
    network_devices {

      device_id = "string"
      mesh_role = "string"
    }
    rf_profile_name = "string"
    site_id         = "string"
  }
}

output "dnacenter_wireless_access_points_provision_example" {
  value = dnacenter_wireless_access_points_provision.example
}