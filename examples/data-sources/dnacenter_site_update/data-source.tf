
data "dnacenter_site_update" "example" {
  provider = dnacenter
  site_id  = "string"
  site {

    area {

      name        = "string"
      parent_name = "string"
    }
    building {

      address     = "string"
      latitude    = 1
      longitude   = 1
      name        = "string"
      parent_name = "string"
    }
    floor {

      height   = 1
      length   = 1
      name     = "string"
      rf_model = "string"
      width    = 1
    }
  }
  type = "string"
}