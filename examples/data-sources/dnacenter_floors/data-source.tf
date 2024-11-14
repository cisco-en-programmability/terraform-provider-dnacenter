
data "dnacenter_floors" "example" {
  provider         = dnacenter
  id               = "string"
  units_of_measure = "string"
}

output "dnacenter_floors_example" {
  value = data.dnacenter_floors.example.item
}
