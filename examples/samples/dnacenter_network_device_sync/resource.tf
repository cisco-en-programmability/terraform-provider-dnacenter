
resource "dnacenter_network_device_sync" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    force_sync = "false"
  }
}