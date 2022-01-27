
data "dnacenter_swim_trigger_activation" "example" {
  provider                     = dnacenter
  client_type                  = "string"
  client_url                   = "string"
  schedule_validate            = "false"
  activate_lower_image_version = "false"
  device_upgrade_mode          = "string"
  device_uuid                  = "string"
  distribute_if_needed         = "false"
  image_uuid_list              = ["string"]
  smu_image_uuid_list          = ["string"]
}