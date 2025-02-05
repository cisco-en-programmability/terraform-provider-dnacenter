
data "dnacenter_wireless_settings_dot11be_profiles" "example" {
  provider             = dnacenter
  is_mu_mimo_down_link = "false"
  is_mu_mimo_up_link   = "false"
  is_of_dma_down_link  = "false"
  is_of_dma_multi_ru   = "false"
  is_of_dma_up_link    = "false"
  limit                = 1
  offset               = 1
  profile_name         = "string"
}

output "dnacenter_wireless_settings_dot11be_profiles_example" {
  value = data.dnacenter_wireless_settings_dot11be_profiles.example.items
}

data "dnacenter_wireless_settings_dot11be_profiles" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_wireless_settings_dot11be_profiles_example" {
  value = data.dnacenter_wireless_settings_dot11be_profiles.example.item
}
