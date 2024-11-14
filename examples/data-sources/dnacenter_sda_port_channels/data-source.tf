
data "dnacenter_sda_port_channels" "example" {
  provider              = dnacenter
  connected_device_type = "string"
  fabric_id             = "string"
  limit                 = 1
  network_device_id     = "string"
  offset                = 1
  port_channel_name     = "string"
}

output "dnacenter_sda_port_channels_example" {
  value = data.dnacenter_sda_port_channels.example.items
}
