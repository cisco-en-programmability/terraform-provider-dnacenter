
data "dnacenter_areas" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_areas_example" {
  value = data.dnacenter_areas.example.item
}
