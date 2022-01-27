
data "dnacenter_sensor_test_template_edit" "example" {
  provider = dnacenter
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