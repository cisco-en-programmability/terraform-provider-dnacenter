
data "dnacenter_device_replacement_count" "example" {
  provider           = dnacenter
  replacement_status = ["string"]
}

output "dnacenter_device_replacement_count_example" {
  value = data.dnacenter_device_replacement_count.example.item
}
