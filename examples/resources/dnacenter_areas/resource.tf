
resource "dnacenter_areas" "example" {
  provider = dnacenter
  parameters {

    id        = "string"
    name      = "string"
    parent_id = "string"
  }
}

output "dnacenter_areas_example" {
  value = dnacenter_areas.example
}
