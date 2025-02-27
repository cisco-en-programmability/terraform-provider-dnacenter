---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dnacenter_security_advisories_summary Data Source - terraform-provider-dnacenter"
subcategory: ""
description: |-
  It performs read operation on Security Advisories.
  Retrieves summary of advisories on the network.
---

# dnacenter_security_advisories_summary (Data Source)

It performs read operation on Security Advisories.

- Retrieves summary of advisories on the network.

## Example Usage

```terraform
data "dnacenter_security_advisories_summary" "example" {
  provider = dnacenter
}

output "dnacenter_security_advisories_summary_example" {
  value = data.dnacenter_security_advisories_summary.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `cri_tic_al` (List of Object) (see [below for nested schema](#nestedobjatt--item--cri_tic_al))
- `hig_h` (List of Object) (see [below for nested schema](#nestedobjatt--item--hig_h))
- `inf_orm_ati_ona_l` (List of Object) (see [below for nested schema](#nestedobjatt--item--inf_orm_ati_ona_l))
- `low` (List of Object) (see [below for nested schema](#nestedobjatt--item--low))
- `med_ium` (List of Object) (see [below for nested schema](#nestedobjatt--item--med_ium))
- `na` (List of Object) (see [below for nested schema](#nestedobjatt--item--na))

<a id="nestedobjatt--item--cri_tic_al"></a>
### Nested Schema for `item.cri_tic_al`

Read-Only:

- `con_fig` (Number)
- `cus_tom_con_fig` (Number)
- `tot_al` (Number)
- `ver_sio_n` (Number)


<a id="nestedobjatt--item--hig_h"></a>
### Nested Schema for `item.hig_h`

Read-Only:

- `con_fig` (Number)
- `cus_tom_con_fig` (Number)
- `tot_al` (Number)
- `ver_sio_n` (Number)


<a id="nestedobjatt--item--inf_orm_ati_ona_l"></a>
### Nested Schema for `item.inf_orm_ati_ona_l`

Read-Only:

- `con_fig` (Number)
- `cus_tom_con_fig` (Number)
- `tot_al` (Number)
- `ver_sio_n` (Number)


<a id="nestedobjatt--item--low"></a>
### Nested Schema for `item.low`

Read-Only:

- `con_fig` (Number)
- `cus_tom_con_fig` (Number)
- `tot_al` (Number)
- `ver_sio_n` (Number)


<a id="nestedobjatt--item--med_ium"></a>
### Nested Schema for `item.med_ium`

Read-Only:

- `con_fig` (Number)
- `cus_tom_con_fig` (Number)
- `tot_al` (Number)
- `ver_sio_n` (Number)


<a id="nestedobjatt--item--na"></a>
### Nested Schema for `item.na`

Read-Only:

- `con_fig` (Number)
- `cus_tom_con_fig` (Number)
- `tot_al` (Number)
- `ver_sio_n` (Number)
