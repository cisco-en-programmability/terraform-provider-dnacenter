---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_floors_upload_image_v2 Resource - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs create operation on Site Design.
  Uploads floor image.
---

# dnacenter_floors_upload_image_v2 (Resource)

It performs create operation on Site Design.

- Uploads floor image.
~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in DNACenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "dnacenter_floors_upload_image_v2" "example" {
  provider = meraki
  id       = "string"
  parameters {

  }
}

output "dnacenter_floors_upload_image_v2_example" {
  value = dnacenter_floors_upload_image_v2.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `parameters` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (String)
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Required:

- `id` (String) id path parameter. Floor Id