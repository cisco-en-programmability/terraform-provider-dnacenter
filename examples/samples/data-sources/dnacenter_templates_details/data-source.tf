
data "dnacenter_templates_details" "example" {
  provider                     = dnacenter
  all_template_attributes      = "false"
  filter_conflicting_templates = "false"
  id                           = "string"
  include_version_details      = "false"
  limit                        = 1
  name                         = "string"
  offset                       = 1
  product_family               = "string"
  product_series               = "string"
  product_type                 = "string"
  project_id                   = "string"
  project_name                 = "string"
  software_type                = "string"
  software_version             = "string"
  sort_order                   = "string"
  tags                         = ["string"]
  un_committed                 = "false"
}

output "dnacenter_templates_details_example" {
  value = data.dnacenter_templates_details.example.items
}
