
data "dnacenter_site_membership" "example" {
  provider      = dnacenter
  device_family = "string"
  limit         = 1
  offset        = 1
  serial_number = "string"
  site_id       = "string"
}

output "dnacenter_site_membership_example" {
  value = data.dnacenter_site_membership.example.item
}
