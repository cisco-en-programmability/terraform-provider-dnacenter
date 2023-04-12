
resource "dnacenter_network_device_sync" "example" {
  provider = dnacenter
  parameters {
    payload    = ["string"]
    force_sync = false
  }
}

output "dnacenter_network_device_sync_example" {
  value = dnacenter_network_device_sync.example
}