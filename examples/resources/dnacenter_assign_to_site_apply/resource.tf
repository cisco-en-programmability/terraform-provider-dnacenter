
resource "dnacenter_assign_to_site_apply" "example" {
  provider = dnacenter
  parameters {

    device_ids = ["string"]
    site_id    = "string"
  }
}

output "dnacenter_assign_to_site_apply_example" {
  value = dnacenter_assign_to_site_apply.example
}