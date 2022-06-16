
resource "dnacenter_network_device_custom_prompt" "example" {
  provider = dnacenter
  parameters {

    password_prompt = "******"
    username_prompt = "string"
  }
}

output "dnacenter_network_device_custom_prompt_example" {
  value = dnacenter_network_device_custom_prompt.example
}