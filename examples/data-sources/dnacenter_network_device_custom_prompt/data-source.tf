
data "dnacenter_network_device_custom_prompt" "example" {
  provider = dnacenter
}

output "dnacenter_network_device_custom_prompt_example" {
  value = data.dnacenter_network_device_custom_prompt.example.item
}
