
data "dnacenter_topology_site" "example" {
  provider = dnacenter
}

output "dnacenter_topology_site_example" {
  value = data.dnacenter_topology_site.example.item
}
