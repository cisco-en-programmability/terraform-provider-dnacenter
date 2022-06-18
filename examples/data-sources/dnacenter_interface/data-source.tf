
data "dnacenter_interface" "example" {
  provider       = dnacenter
  interface_uuid = "string"
}

output "dnacenter_interface_example" {
  value = data.dnacenter_interface.example.item
}
