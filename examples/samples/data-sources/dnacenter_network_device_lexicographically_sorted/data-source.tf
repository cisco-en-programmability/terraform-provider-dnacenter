
data "dnacenter_network_device_lexicographically_sorted" "example" {
  provider                    = dnacenter
  associated_wlc_ip           = "string"
  collection_interval         = "string"
  collection_status           = "string"
  error_code                  = "string"
  family                      = "string"
  hostname                    = "string"
  limit                       = 1
  mac_address                 = "string"
  management_ip_address       = "string"
  offset                      = 1
  platform_id                 = "string"
  reachability_failure_reason = "string"
  reachability_status         = "string"
  role                        = "string"
  role_source                 = "string"
  serial_number               = "string"
  series                      = "string"
  software_type               = "string"
  software_version            = "string"
  type                        = "string"
  up_time                     = "string"
  vrf_name                    = "string"
}

output "dnacenter_network_device_lexicographically_sorted_example" {
  value = data.dnacenter_network_device_lexicographically_sorted.example.items
}
