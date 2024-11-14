
data "dnacenter_site_count_v2" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_site_count_v2_example" {
  value = data.dnacenter_site_count_v2.example.item
}
