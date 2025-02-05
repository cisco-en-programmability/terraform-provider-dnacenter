
resource "dnacenter_network_device_maintenance_schedules_id" "example" {
  provider = dnacenter

  parameters {

    description = "string"
    id          = "string"
    maintenance_schedule {

      end_time = 1609459200
      recurrence {

        interval            = 1
        recurrence_end_time = 1.0
      }
      start_time = 1609459200
    }
    network_device_ids = ["string"]
  }
}

output "dnacenter_network_device_maintenance_schedules_id_example" {
  value = dnacenter_network_device_maintenance_schedules_id.example
}
