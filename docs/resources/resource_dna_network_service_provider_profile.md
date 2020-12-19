---
page_title: "dna_network_service_provider_profile Resource - terraform-provider-dnacenter"
subcategory: "Network Settings"
description: |-
  The dna_network_service_provider_profile resource allows you to configure a DNACenter Service Provider profile.
---

# Resource dna_network_service_provider_profile

The dna_network_service_provider_profile resource allows you to configure a DNACenter Service Provider profile.

## Example Usage

```hcl
resource "dna_network_service_provider_profile" "response" {
  provider     = dnacenter
  profile_name = "Test1"
  model        = "6-class-model"
  wan_provider = "test1-provider"
}
```

## Argument Reference

- `profile_name` - (Required) The server provider profile's name. If it's changed it forces the creation of a new resource.
- `model` - (Required) The server provider profile's model. Available values are "4-class-model", "5-class-model", "6-class-model" and "8-class-model".
- `wan_provider` - (Required) The server provider profile's wan_provider.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `last_updated` - The service provider profile's updated time with format RFC850.
