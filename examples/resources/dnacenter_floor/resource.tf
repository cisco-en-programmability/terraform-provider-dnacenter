
resource "dnacenter_floor" "example" {
  provider = dnacenter
  parameters {

    site {

      floor {

        floor_number = 1.0
        height       = 1.0
        length       = 1.0
        name         = "string"
        parent_name  = "string"
        rf_model     = "string"
        width        = 1.0
      }
    }
    site_id = "string"
    type    = "string"
  }
}

output "dnacenter_floor_example" {
  value = dnacenter_floor.example
}