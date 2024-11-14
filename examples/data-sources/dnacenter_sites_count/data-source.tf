
data "dnacenter_sites_count" "example" {
  provider = dnacenter
  name     = "string"
}

output "dnacenter_sites_count_example" {
  value = data.dnacenter_sites_count.example.items
}
