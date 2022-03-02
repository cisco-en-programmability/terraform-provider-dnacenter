
resource "dnacenter_site" "example" {
  provider = dnacenter
  parameters {
    site {
      area {
        name        = "string"
        parent_name = "string"
      }
      building {
        name        = "string"
        address     = "string"
        parent_name = "string"
        latitude    = 1
        longitude   = 1
      }
      floor {
        height      = 1
        length      = 1
        name        = "string"
        parent_name = "string"
        rf_model    = "string"
        width       = 1
      }
    }
    type = "string"
  }
}

output "dnacenter_site_example" {
  value = dnacenter_site.example
}