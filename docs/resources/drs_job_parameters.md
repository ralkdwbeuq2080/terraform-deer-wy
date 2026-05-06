---
subcategory: "Data Replication Service (DRS)"
layout: "huaweicloud"
page_title: "HuaweiCloud: huaweicloud_drs_job_parameters"
description: |-
  Use this data source to get the list of job parameter configurations for a specified DRS job within HuaweiCloud.
---

# huaweicloud_drs_job_parameters

Use this data source to get the list of job parameter configurations for a specified DRS job within HuaweiCloud.

## Example Usage

```hcl
data "huaweicloud_drs_job_parameters" "test" {
  job_id = "e11eaf8f-71ef-4fad-8890-aed7572ajb11"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String) Specifies the region in which to query the resource.
  If omitted, the provider-level region will be used.

* `job_id` - (Required, String) Specifies the ID of the DRS job.

* `offset` - (Optional, Int) Specifies the offset from which the query starts. The value must be greater than or equal to 0. Defaults to 0.

* `limit` - (Optional, Int) Specifies the number of records displayed on each page. The value ranges from 1 to 1000. Defaults to 10.

* `name` - (Optional, String) Specifies the parameter name to filter the results.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The data source ID.

* `total_count` - The total number of parameters matched by the query.

* `parameter_config_list` - The list of DRS job parameter configurations.

  The [parameter_config_list](#parameter_config_list_struct) structure is documented below.

<a name="parameter_config_list_struct"></a>
The `parameter_config_list` block supports:

* `name` - The name of the parameter.

* `value` - The current value of the parameter.

* `default_value` - The default value of the parameter.

* `value_range` - Indicates the value range.

* `is_need_restart` - Whether the job needs to be restarted for the parameter modification to take effect.

* `description` - The description of the parameter.

* `created_at` - The creation time of the parameter, in RFC3339 format.

* `updated_at` - The last update time of the parameter, in RFC3339 format.
