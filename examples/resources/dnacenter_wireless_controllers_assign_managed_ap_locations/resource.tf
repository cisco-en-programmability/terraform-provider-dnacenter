
resource "dnacenter_wireless_controllers_assign_managed_ap_locations" "example" {
  provider  = dnacenter
  device_id = "string"
  parameters {

    primary_managed_aplocations_site_ids   = ["string"]
    secondary_managed_aplocations_site_ids = ["string"]
  }
}

output "dnacenter_wireless_controllers_assign_managed_ap_locations_example" {
  value = dnacenter_wireless_controllers_assign_managed_ap_locations.example
}