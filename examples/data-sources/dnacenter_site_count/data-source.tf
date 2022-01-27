
data "dnacenter_site_count" "example" {
  provider = dnacenter
  site_id  = "string"
}

output "dnacenter_site_count_example" {
  value = data.dnacenter_site_count.example.item
}
