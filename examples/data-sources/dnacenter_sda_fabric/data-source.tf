
data "dnacenter_sda_fabric" "example" {
  provider    = dnacenter
  fabric_name = "string"
}

output "dnacenter_sda_fabric_example" {
  value = data.dnacenter_sda_fabric.example.item
}
