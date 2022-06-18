
resource "dnacenter_assign_device_to_site" "example" {
  provider = dnacenter
  parameters {

    device {

      ip = "string"
    }
    site_id = "string"
  }
}

output "dnacenter_assign_device_to_site_example" {
  value = dnacenter_assign_device_to_site.example
}