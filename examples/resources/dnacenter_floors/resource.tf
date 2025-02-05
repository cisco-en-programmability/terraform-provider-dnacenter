
resource "dnacenter_floors" "example" {
  provider = dnacenter

  parameters {

    floor_number     = 1
    height           = 1.0
    id               = "string"
    length           = 1.0
    name             = "string"
    parent_id        = "string"
    rf_model         = "string"
    units_of_measure = "string"
    width            = 1.0
  }
}

output "dnacenter_floors_example" {
  value = dnacenter_floors.example
}
