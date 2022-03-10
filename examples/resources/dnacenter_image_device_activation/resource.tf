
resource "dnacenter_image_device_activation" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    client_type       = "string"
    client_url        = "string"
    schedule_validate = "false"
    payload {
      activate_lower_image_version = "false"
      device_upgrade_mode          = "string"
      device_uuid                  = "string"
      distribute_if_needed         = "false"
      image_uuid_list              = ["string"]
      smu_image_uuid_list          = ["string"]
    }
  }
}

output "dnacenter_image_device_activation_example" {
  value = dnacenter_image_device_activation.example
}

data "dnacenter_task" "example" {
  depends_on = [dnacenter_image_device_activation.example]
  provider = dnacenter
  task_id  = dnacenter_image_device_activation.example.item.0.task_id
}

output "dnacenter_task_example" {
  value = data.dnacenter_task.example.item
} 
