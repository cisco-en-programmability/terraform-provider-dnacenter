
resource "dnacenter_deploy_template" "example" {
  provider = dnacenter
  parameters {

    force_push_template             = "false"
    is_composite                    = "false"
    main_template_id                = "string"
    member_template_deployment_info = ["string"]
    target_info {

      host_name             = "string"
      id                    = "string"
      params                = "string"
      resource_params       = ["string"]
      type                  = "string"
      versioned_template_id = "string"
    }
    template_id = "string"
  }
}

output "dnacenter_deploy_template_example" {
  value = dnacenter_deploy_template.example
}