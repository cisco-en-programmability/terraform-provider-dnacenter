
data "dnacenter_network" "example" {
  provider = dnacenter
  site_id  = "string"
}

output "dnacenter_network_example" {
  value = data.dnacenter_network.example.items
}
