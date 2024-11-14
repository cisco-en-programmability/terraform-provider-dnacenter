
data "dnacenter_sda_provision_devices_count" "example" {
  provider = dnacenter
  site_id  = "string"
}

output "dnacenter_sda_provision_devices_count_example" {
  value = data.dnacenter_sda_provision_devices_count.example.item
}
