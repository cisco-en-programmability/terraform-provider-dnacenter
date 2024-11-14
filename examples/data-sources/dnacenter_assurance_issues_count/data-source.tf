
data "dnacenter_assurance_issues_count" "example" {
  provider                  = dnacenter
  ai_driven                 = "false"
  category                  = "string"
  device_type               = "string"
  end_time                  = 1609459200
  entity_id                 = "string"
  entity_type               = "string"
  fabric_driven             = "false"
  fabric_site_driven        = "false"
  fabric_site_id            = "string"
  fabric_transit_driven     = "false"
  fabric_transit_site_id    = "string"
  fabric_vn_driven          = "false"
  fabric_vn_name            = "string"
  is_global                 = "false"
  issue_id                  = "string"
  mac_address               = "string"
  name                      = "string"
  network_device_id         = "string"
  network_device_ip_address = "string"
  priority                  = "string"
  severity                  = "string"
  site_hierarchy            = "string"
  site_hierarchy_id         = "string"
  site_id                   = "string"
  site_name                 = "string"
  start_time                = 1609459200
  status                    = "string"
  updated_by                = "string"
  xca_lle_rid               = "string"
}

output "dnacenter_assurance_issues_count_example" {
  value = data.dnacenter_assurance_issues_count.example.item
}
