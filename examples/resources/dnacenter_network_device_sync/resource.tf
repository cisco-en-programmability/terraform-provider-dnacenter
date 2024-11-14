
resource "dnacenter_network_device_sync" "example" {
  provider   = dnacenter
  force_sync = "false"
  parameters = ["string"]
}

output "dnacenter_network_device_sync_example" {
  value = dnacenter_network_device_sync.example
}