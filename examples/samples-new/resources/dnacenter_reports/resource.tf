terraform {
  required_providers {
  dnacenter = {
    version = "0.0.3"
    source  = "hashicorp.com/edu/dnacenter"
    # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
  }
  }
}
resource "dnacenter_reports" "example" {
  provider = dnacenter
  parameters {

    #deliveries = ["string"]
    name       = "string"
    #report_id  = "string"
    #schedule   = ["string"]
    #tags       = ["string"]
    /*view {

      field_groups {

        field_group_display_name = "string"
        field_group_name         = "string"
        fields {

          display_name = "string"
          name         = "string"
        }
      }
      filters {

        display_name = "string"
        name         = "string"
        type         = "string"
        value        = ["string"]
      }
      format {

        format_type = "string"
        name        = "string"
      }
      name    = "string"
      view_id = "string"
    }
    view_group_id      = "string"
    view_group_version = "string"*/
  }
}

output "dnacenter_reports_example" {
  value = dnacenter_reports.example
}