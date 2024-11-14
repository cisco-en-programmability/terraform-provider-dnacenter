
data "dnacenter_buildings" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_buildings_example" {
  value = data.dnacenter_buildings.example.item
}
