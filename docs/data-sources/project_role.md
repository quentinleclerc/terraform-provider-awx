---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awx_project_role Data Source - terraform-provider-awx"
subcategory: ""
description: |-
  Use this data source to get the details of a project role in AWX.
---

# awx_project_role (Data Source)

Use this data source to get the details of a project role in AWX.

## Example Usage

```terraform
data "awx_project" "example" {
  name = "my_project"
}

data "awx_project_role" "example" {
  name       = "Admin"
  project_id = awx_project.example.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `project_id` (Number) The unique identifier of the project.

### Optional

- `id` (Number) The unique identifier of the project role.
- `name` (String) The name of the project role.
