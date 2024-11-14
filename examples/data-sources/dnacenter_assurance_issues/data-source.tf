
data "dnacenter_assurance_issues" "example" {
  provider                  = dnacenter
  accept_language           = "string"
  ai_driven                 = "false"
  attribute                 = "string"
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
  limit                     = 1
  mac_address               = "string"
  name                      = "string"
  network_device_id         = "string"
  network_device_ip_address = "string"
  offset                    = 1
  order                     = "string"
  priority                  = "string"
  severity                  = "string"
  site_hierarchy            = "string"
  site_hierarchy_id         = "string"
  site_id                   = "string"
  site_name                 = "string"
  sort_by                   = "string"
  start_time                = 1609459200
  status                    = "string"
  updated_by                = "string"
  view                      = "string"
  xca_lle_rid               = "string"
}

output "dnacenter_assurance_issues_example" {
  value = data.dnacenter_assurance_issues.example.items
}

data "dnacenter_assurance_issues" "example" {
  provider        = dnacenter
  accept_language = "string"
  attribute       = "string"
  id              = "string"
  view            = "string"
  xca_lle_rid     = "string"
}

output "dnacenter_assurance_issues_example" {
  value = data.dnacenter_assurance_issues.example.item
}
