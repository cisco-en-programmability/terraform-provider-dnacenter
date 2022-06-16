
data "dnacenter_discovery_count" "example" {
  provider = dnacenter
}

output "dnacenter_discovery_count_example" {
  value = data.dnacenter_discovery_count.example.item
}
