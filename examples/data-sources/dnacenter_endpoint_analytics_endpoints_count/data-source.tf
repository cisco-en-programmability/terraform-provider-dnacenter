
data "dnacenter_endpoint_analytics_endpoints_count" "example" {
  provider                    = dnacenter
  ai_spoofing_trust_level     = "string"
  anc_policy                  = "string"
  auth_method                 = "string"
  changed_profile_trust_level = "string"
  concurrent_mac_trust_level  = "string"
  device_type                 = "string"
  hardware_manufacturer       = "string"
  hardware_model              = "string"
  ip                          = "string"
  ip_blocklist_detected       = "false"
  mac_address                 = "string"
  mac_addresses               = ["string"]
  nat_trust_level             = "string"
  operating_system            = "string"
  posture_status              = "string"
  profiling_status            = "string"
  random_mac                  = "false"
  registered                  = "false"
  trust_score                 = "string"
  unauth_port_detected        = "false"
  weak_cred_detected          = "false"
}

output "dnacenter_endpoint_analytics_endpoints_count_example" {
  value = data.dnacenter_endpoint_analytics_endpoints_count.example.item
}
