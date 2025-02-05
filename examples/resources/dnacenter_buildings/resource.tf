
resource "dnacenter_buildings" "example" {
  provider = dnacenter

  parameters {

    address   = "string"
    country   = "string"
    id        = "string"
    latitude  = 1.0
    longitude = 1.0
    name      = "string"
    parent_id = "string"
  }
}

output "dnacenter_buildings_example" {
  value = dnacenter_buildings.example
}
