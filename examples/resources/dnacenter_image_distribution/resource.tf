
resource "dnacenter_image_distribution" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    payload {
      device_uuid = "string"
      image_uuid  = "string"
    }
  }
}

output "dnacenter_image_distribution_example" {
  value = dnacenter_image_distribution.example
}

data "dnacenter_task" "example" {
  depends_on = [dnacenter_image_distribution.example]
  provider = dnacenter
  task_id  = dnacenter_image_distribution.example.item.0.task_id
}

output "dnacenter_task_example" {
  value = data.dnacenter_task.example.item
} 
