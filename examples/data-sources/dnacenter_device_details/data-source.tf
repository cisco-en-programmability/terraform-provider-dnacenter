
data "dnacenter_device_details" "example" {
  provider   = dnacenter
  identifier = "string"
  search_by  = "string"
  timestamp  = 1.0
}

output "dnacenter_device_details_example" {
  value = data.dnacenter_device_details.example.item
}
