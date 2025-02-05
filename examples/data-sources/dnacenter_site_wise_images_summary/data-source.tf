
data "dnacenter_site_wise_images_summary" "example" {
  provider = dnacenter
  site_id  = "string"
}

output "dnacenter_site_wise_images_summary_example" {
  value = data.dnacenter_site_wise_images_summary.example.item
}
