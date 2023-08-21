
resource "dnacenter_building" "example" {
  provider = dnacenter
  parameters {

    site {

      building {

        address     = "string"
        country     = "string"
        latitude    = 1.0
        longitude   = 1.0
        name        = "string"
        parent_name = "string"
      }
    }
    site_id = "string"
    type    = "string"
  }
}

output "dnacenter_building_example" {
  value = dnacenter_building.example
}