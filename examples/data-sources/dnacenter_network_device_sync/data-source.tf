
data "dnacenter_network_device_sync" "example" {
  provider   = dnacenter
  force_sync = "false"
}