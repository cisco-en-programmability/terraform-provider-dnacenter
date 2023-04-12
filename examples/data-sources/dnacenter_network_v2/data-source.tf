
data "dnacenter_network_v2" "example" {
  provider = dnacenter
  site_id  = "string"
}

output "dnacenter_network_v2_example" {
  value = data.dnacenter_network_v2.example.items
}
