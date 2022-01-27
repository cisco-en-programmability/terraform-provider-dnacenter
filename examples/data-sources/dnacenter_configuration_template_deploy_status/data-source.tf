
data "dnacenter_configuration_template_deploy_status" "example" {
  provider      = dnacenter
  deployment_id = "string"
}

output "dnacenter_configuration_template_deploy_status_example" {
  value = data.dnacenter_configuration_template_deploy_status.example.item
}
