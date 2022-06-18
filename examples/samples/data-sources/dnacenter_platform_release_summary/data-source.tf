
data "dnacenter_platform_release_summary" "example" {
  provider = dnacenter
}

output "dnacenter_platform_release_summary_example" {
  value = data.dnacenter_platform_release_summary.example.item
}
