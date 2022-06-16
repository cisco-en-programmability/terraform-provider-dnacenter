
data "dnacenter_deploy_template" "example" {
  provider      = dnacenter
  deployment_id = "string"
}

output "dnacenter_deploy_template_example" {
  value = data.dnacenter_deploy_template.example.item
}
