terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

resource "dnacenter_configuration_template_project" "example" {
  provider = dnacenter
  parameters {

    create_time = 1
    description = "Cloud DayN Templates 3"
    #id = "string"
    last_update_time = 1
    name             = "Cloud DayN Templates News"
    /*
      project_id = "string"
      tags {  
        id = "string"
        name = "string"
      }
      templates {
        
        author = "string"
        composite = "false"
        containing_templates {
          
          composite = "false"
          description = "string"
          device_types {
            
            product_family = "string"
            product_series = "string"
            product_type = "string"
          }
          id = "string"
          language = "string"
          name = "string"
          project_name = "string"
          rollback_template_params {
            
            binding = "string"
            custom_order = 1
            data_type = "string"
            default_value = "string"
            description = "string"
            display_name = "string"
            group = "string"
            id = "string"
            instruction_text = "string"
            key = "string"
            not_param = "false"
            order = 1
            param_array = "false"
            parameter_name = "string"
            provider = "string"
            range {
              
              id = "string"
              max_value = 1
              min_value = 1
            }
            required = "false"
            selection {
              
              default_selected_values = ["string"]
              id = "string"
              selection_type = "string"
              selection_values = ["string"]
            }
          }
          tags {
            
            id = "string"
            name = "string"
          }
          template_content = "string"
          template_params {
            
            binding = "string"
            custom_order = 1
            data_type = "string"
            default_value = "string"
            description = "string"
            display_name = "string"
            group = "string"
            id = "string"
            instruction_text = "string"
            key = "string"
            not_param = "false"
            order = 1
            param_array = "false"
            parameter_name = "string"
            provider = "string"
            range {
              
              id = "string"
              max_value = 1
              min_value = 1
            }
            required = "false"
            selection {
              
              default_selected_values = ["string"]
              id = "string"
              selection_type = "string"
              selection_values = ["string"]
            }
          }
          version = "string"
        }
        create_time = 1
        custom_params_order = "false"
        description = "string"
        device_types {
          
          product_family = "string"
          product_series = "string"
          product_type = "string"
        }
        document_database = "false"
        failure_policy = "string"
        id = "string"
        language = "string"
        last_update_time = 1
        latest_version_time = 1
        name = "string"
        parent_template_id = "string"
        project_associated = "false"
        project_id = "string"
        project_name = "string"
        rollback_template_content = "string"
        rollback_template_params {
          
          binding = "string"
          custom_order = 1
          data_type = "string"
          default_value = "string"
          description = "string"
          display_name = "string"
          group = "string"
          id = "string"
          instruction_text = "string"
          key = "string"
          not_param = "false"
          order = 1
          param_array = "false"
          parameter_name = "string"
          provider = "string"
          range {
            
            id = "string"
            max_value = 1
            min_value = 1
          }
          required = "false"
          selection {
            
            default_selected_values = ["string"]
            id = "string"
            selection_type = "string"
            selection_values = ["string"]
          }
        }
        software_type = "string"
        software_variant = "string"
        software_version = "string"
        tags {
          
          id = "string"
          name = "string"
        }
        template_content = "string"
        template_params {
          
          binding = "string"
          custom_order = 1
          data_type = "string"
          default_value = "string"
          description = "string"
          display_name = "string"
          group = "string"
          id = "string"
          instruction_text = "string"
          key = "string"
          not_param = "false"
          order = 1
          param_array = "false"
          parameter_name = "string"
          provider = "string"
          range {
            
            id = "string"
            max_value = 1
            min_value = 1
          }
          required = "false"
          selection {
            
            default_selected_values = ["string"]
            id = "string"
            selection_type = "string"
            selection_values = ["string"]
          }
        }
        validation_errors {
          
          rollback_template_errors = ["string"]
          template_errors = ["string"]
          template_id = "string"
          template_version = "string"
        }
        version = "string"
      }*/
    templates {
      name                = "Cloud Test Template  6"
      composite           = false
      language            = "VELOCITY"
      id                  = "331a801a-56f5-43c0-93d0-9d713f27cfd3"
      custom_params_order = false
      last_update_time    = 1636580362198
      latest_version_time = 0
      project_associated  = true
      document_database   = false
    }
  }
}

output "dnacenter_configuration_template_project_example" {
  value = dnacenter_configuration_template_project.example
}
