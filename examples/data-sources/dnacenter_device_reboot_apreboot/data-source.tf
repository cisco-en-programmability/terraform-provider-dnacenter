
data "dnacenter_device_reboot_apreboot" "example" {
  provider       = dnacenter
  parent_task_id = "string"
}

output "dnacenter_device_reboot_apreboot_example" {
  value = data.dnacenter_device_reboot_apreboot.example.items
}
