
data "dnacenter_site_membership" "example" {
  provider      = dnacenter
  device_family = "string"
  limit         = "string"
  offset        = "string"
  serial_number = "string"
  site_id       = "string"
}

output "dnacenter_site_membership_example" {
  value = data.dnacenter_site_membership.example.item
}
