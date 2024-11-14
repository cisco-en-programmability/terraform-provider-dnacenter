
data "dnacenter_sda_port_channels_count" "example" {
  provider              = dnacenter
  connected_device_type = "string"
  fabric_id             = "string"
  network_device_id     = "string"
  port_channel_name     = "string"
}

output "dnacenter_sda_port_channels_count_example" {
  value = data.dnacenter_sda_port_channels_count.example.item
}
