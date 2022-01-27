
data "dnacenter_network_device_list" "example" {
  provider                   = dnacenter
  associated_wlc_ip          = ["string"]
  collection_interval        = ["string"]
  collection_status          = ["string"]
  device_support_level       = "string"
  error_code                 = ["string"]
  error_description          = ["string"]
  family                     = ["string"]
  hostname                   = ["string"]
  id                         = "string"
  license_name               = ["string"]
  license_status             = ["string"]
  license_type               = ["string"]
  location                   = ["string"]
  location_name              = ["string"]
  mac_address                = ["string"]
  management_ip_address      = ["string"]
  module_equpimenttype       = ["string"]
  module_name                = ["string"]
  module_operationstatecode  = ["string"]
  module_partnumber          = ["string"]
  module_servicestate        = ["string"]
  module_vendorequipmenttype = ["string"]
  not_synced_for_minutes     = ["string"]
  platform_id                = ["string"]
  reachability_status        = ["string"]
  role                       = ["string"]
  serial_number              = ["string"]
  series                     = ["string"]
  software_type              = ["string"]
  software_version           = ["string"]
  type                       = ["string"]
  up_time                    = ["string"]
}

output "dnacenter_network_device_list_example" {
  value = data.dnacenter_network_device_list.example.items
}
