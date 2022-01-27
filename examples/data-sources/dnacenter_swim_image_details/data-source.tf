
data "dnacenter_swim_image_details" "example" {
  provider                = dnacenter
  application_type        = "string"
  created_time            = 1
  family                  = "string"
  image_integrity_status  = "string"
  image_name              = "string"
  image_series            = "string"
  image_size_greater_than = 1
  image_size_lesser_than  = 1
  image_uuid              = "string"
  is_cco_latest           = "false"
  is_cco_recommended      = "false"
  is_tagged_golden        = "false"
  limit                   = 1
  name                    = "string"
  offset                  = 1
  sort_by                 = "string"
  sort_order              = "string"
  version                 = "string"
}

output "dnacenter_swim_image_details_example" {
  value = data.dnacenter_swim_image_details.example.items
}
