
resource "dnacenter_area" "example" {
  provider = dnacenter
  parameters {

    site {

      area {

        name        = "string"
        parent_name = "string"
      }

    }
    site_id = "string"
    type    = "string"
  }
}

output "dnacenter_area_example" {
  value = dnacenter_area.example
}