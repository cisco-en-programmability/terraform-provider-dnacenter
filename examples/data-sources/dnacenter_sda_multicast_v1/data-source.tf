
data "dnacenter_sda_multicast_v1" "example" {
  provider  = dnacenter
  fabric_id = "string"
  limit     = 1
  offset    = 1
}

output "dnacenter_sda_multicast_v1_example" {
  value = data.dnacenter_sda_multicast_v1.example.items
}
