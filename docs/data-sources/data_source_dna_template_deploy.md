---
page_title: "dna_template_deploy Data Source - terraform-provider-dnacenter"
subcategory: "Configuration Templates"
description: |-
  The dna_template_deploy data source allows you to deploy a DNACenter template.
---

# Data Source dna_template_deploy

The dna_template_deploy data source allows you to deploy a DNACenter template.

## Example Usage

```hcl
data "dna_template_deploy" "response" {
  provider = dnacenter
  template_deployment_info {
    template_id = "5f5e3eca-2b43-4228-b7b6-3b957b56c110"
    target_info {
      hostname = "10.121.1.1"
      id       = "1a23a341-7ea2-41e9-8814-989a7d10c4be"
      params = {
        NetworkId = 1
      }
      type = "MANAGED_DEVICE_IP"
    }
  }
  member_templates_deployment_info {
    template_id = "5f5e3eca-2b43-4228-b7b6-3b957b56c110"
    target_info {
      hostname = "10.121.1.1"
      id       = "1a23a341-7ea2-41e9-8814-989a7d10c4be"
      params = {
        NetworkId = 1
      }
      type = "MANAGED_DEVICE_IP"
    }
  }
}
```

## Argument Reference

- `template_deployment_info` - (Required) The template_deployment_info argument. See [template_deployment_info](#template_deployment_info) below for details.
- `member_templates_deployment_info` - (Optional) The member_templates_deployment_info argument. See [member_templates_deployment_info](#member_templates_deployment_info) below for details.

### template_deployment_info

- `force_push_template` - (Optional) The template param's force push template.
- `is_composite` - (Optional) The template param's is composite.
- `main_template_id` - (Optional) The template param's main template id.
- `template_id` - (Required) The template param's template id.
- `target_info` - (Required) The template param's target info. See [target_info](#target_info)

### member_templates_deployment_info

- `force_push_template` - (Optional) The template param's force push template.
- `is_composite` - (Optional) The template param's is composite.
- `main_template_id` - (Optional) The template param's main template id.
- `template_id` - (Optional) The template param's template id.
- `target_info` - (Optional) The template param's target info. See [target_info](#target_info)

#### target_info

- `hostname` - (Optional) The target's hostname.
- `id` - (Optional) The target's id.
- `params` - (Optional) The target's params.
- `type` - (Optional) The target's type.

## Attributes Reference

The following attributes are exported.

- `item` - The item response. See [Item](#item) below for details.

### Item

- `deployment_id` - The template's deployment id.
- `deployment_name` - The template's deployment name.
- `duration` - The template's duration.
- `end_time` - The template's end time.
- `project_name` - The template's project name.
- `start_time` - The template's start time.
- `status` - The template's status.
- `template_name` - The template's template name.
- `template_version` - The template's template version.
- `devices` - The template's devices. See [devices](#devices) below for details.

#### devices

- `device_id` - The template device's id.
- `duration` - The template device's duration.
- `end_time` - The template device's end time.
- `ip_address` - The template device's ip address.
- `name` - The template device's name.
- `start_time` - The template device's start time.
- `status` - The template device's status.
