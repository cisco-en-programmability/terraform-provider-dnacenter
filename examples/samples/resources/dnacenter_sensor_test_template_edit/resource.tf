

terraform {
  required_providers {
    dnacenter = {
      version = "1.0.9-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_sensor_test_template_edit" "example" {
  provider = dnacenter

  parameters {
    location_info_list {

      all_sensors    = "false"
      location_id    = "string"
      location_type  = "string"
      site_hierarchy = "string"
    }
    schedule {

      frequency {

        unit  = "string"
        value = 1
      }
      schedule_range {

        day = "string"
        time_range {

          frequency {

            unit  = "string"
            value = 1
          }
          from = "string"
          to   = "string"
        }
      }
      test_schedule_mode = "string"
    }
    template_name = "string"
  }
}