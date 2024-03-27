package awx

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	awx "github.com/josh-silvas/terraform-provider-awx/tools/goawx"
)

//nolint:funlen
func resourceInventorySource() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource Inventory Source is used to manage inventory sources in AWX.",
		CreateContext: resourceInventorySourceCreate,
		ReadContext:   resourceInventorySourceRead,
		UpdateContext: resourceInventorySourceUpdate,
		DeleteContext: resourceInventorySourceDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the inventory source.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the inventory source.",
			},
			"enabled_var": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The variable that determines if the inventory source is enabled.",
			},
			"enabled_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The value of the variable that determines if the inventory source is enabled.",
			},
			"overwrite": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether to overwrite the inventory source.",
			},
			"overwrite_vars": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether to overwrite the inventory source variables.",
			},
			"update_on_launch": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether to update the inventory source on launch.",
			},
			"inventory_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The inventory to use for the inventory source.",
			},
			"credential_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The credential to use for the inventory source.",
			},
			"source": {
				Type:        schema.TypeString,
				Default:     "scm",
				Optional:    true,
				Description: "The source of the inventory source.",
			},
			"source_vars": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The variables for the inventory source.",
			},
			"host_filter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The host filter for the inventory source.",
			},
			"update_cache_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     30,
				Description: "The update cache timeout for the inventory source.",
			},
			"verbosity": {
				Type:        schema.TypeInt,
				Default:     1,
				Optional:    true,
				Description: "The verbosity for the inventory source. [0,1,2,3]",
			},
			// obsolete schema added so terraform doesn't break
			// these don't do anything in later versions of AWX! Update your code.
			"source_regions": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "[Obsolete] The source regions for the inventory source.",
			},
			"instance_filters": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "[Obsolete] The instance filters for the inventory source.",
			},
			"group_by": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "[Obsolete] The group by for the inventory source.",
			},
			"source_project_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "[Obsolete] The source project for the inventory source.",
			},
			"source_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "[Obsolete] The source path for the inventory source.",
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceInventorySourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*awx.AWX)
	awxService := client.InventorySourcesService

	createInventorySourceData := map[string]interface{}{
		"name":                 d.Get("name").(string),
		"description":          d.Get("description").(string),
		"enabled_var":          d.Get("enabled_var").(string),
		"enabled_value":        d.Get("enabled_value").(string),
		"overwrite":            d.Get("overwrite").(bool),
		"overwrite_vars":       d.Get("overwrite_vars").(bool),
		"update_on_launch":     d.Get("update_on_launch").(bool),
		"inventory":            d.Get("inventory_id").(int),
		"source":               d.Get("source").(string),
		"source_vars":          d.Get("source_vars").(string),
		"host_filter":          d.Get("host_filter").(string),
		"update_cache_timeout": d.Get("update_cache_timeout").(int),
		"verbosity":            d.Get("verbosity").(int),
		// obsolete schema added so terraform doesn't break
		// these don't do anything in later versions of AWX! Update your code.
		"source_regions":   d.Get("source_regions").(string),
		"instance_filters": d.Get("instance_filters").(string),
		"group_by":         d.Get("group_by").(string),
		"source_path":      d.Get("source_path").(string),
	}
	if _, ok := d.GetOk("credential_id"); ok {
		createInventorySourceData["credential"] = d.Get("credential_id").(int)
	}
	if _, ok := d.GetOk("source_project_id"); ok {
		createInventorySourceData["source_project"] = d.Get("source_project_id").(int)
	}

	result, err := awxService.CreateInventorySource(createInventorySourceData, map[string]string{})
	if err != nil {
		return buildDiagCreateFail(diagElementInventorySourceTitle, err)
	}

	d.SetId(strconv.Itoa(result.ID))
	return resourceInventorySourceRead(ctx, d, m)

}

func resourceInventorySourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*awx.AWX)
	awxService := client.InventorySourcesService
	id, diags := convertStateIDToNummeric(diagElementInventorySourceTitle, d)
	if diags.HasError() {
		return diags
	}

	updateInventorySourceData := map[string]interface{}{
		"name":                 d.Get("name").(string),
		"description":          d.Get("description").(string),
		"enabled_var":          d.Get("enabled_var").(string),
		"enabled_value":        d.Get("enabled_value").(string),
		"overwrite":            d.Get("overwrite").(bool),
		"overwrite_vars":       d.Get("overwrite_vars").(bool),
		"update_on_launch":     d.Get("update_on_launch").(bool),
		"inventory":            d.Get("inventory_id").(int),
		"source":               d.Get("source").(string),
		"source_vars":          d.Get("source_vars").(string),
		"host_filter":          d.Get("host_filter").(string),
		"update_cache_timeout": d.Get("update_cache_timeout").(int),
		"verbosity":            d.Get("verbosity").(int),
		// obsolete schema added so terraform doesn't break
		// these don't do anything in later versions of AWX! Update your code.
		"source_regions":   d.Get("source_regions").(string),
		"instance_filters": d.Get("instance_filters").(string),
		"group_by":         d.Get("group_by").(string),
		"source_path":      d.Get("source_path").(string),
	}
	if _, ok := d.GetOk("credential_id"); ok {
		updateInventorySourceData["credential"] = d.Get("credential_id").(int)
	}
	if _, ok := d.GetOk("source_project_id"); ok {
		updateInventorySourceData["source_project"] = d.Get("source_project_id").(int)
	}

	_, err := awxService.UpdateInventorySource(id, updateInventorySourceData, nil)
	if err != nil {
		return buildDiagUpdateFail(diagElementInventorySourceTitle, id, err)
	}

	return resourceInventorySourceRead(ctx, d, m)
}

func resourceInventorySourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*awx.AWX)
	awxService := client.InventorySourcesService
	id, diags := convertStateIDToNummeric(diagElementInventorySourceTitle, d)
	if diags.HasError() {
		return diags
	}
	if _, err := awxService.DeleteInventorySource(id); err != nil {
		return buildDiagDeleteFail(
			"inventroy source",
			fmt.Sprintf("inventroy source %v, got %s ",
				id, err.Error()))
	}
	d.SetId("")
	return nil
}

func resourceInventorySourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*awx.AWX)
	awxService := client.InventorySourcesService
	id, diags := convertStateIDToNummeric(diagElementInventorySourceTitle, d)
	if diags.HasError() {
		return diags
	}
	res, err := awxService.GetInventorySourceByID(id, make(map[string]string))
	if err != nil {
		return buildDiagNotFoundFail(diagElementInventorySourceTitle, id, err)
	}
	d = setInventorySourceResourceData(d, res)
	return nil
}

func setInventorySourceResourceData(d *schema.ResourceData, r *awx.InventorySource) *schema.ResourceData {
	if err := d.Set("name", r.Name); err != nil {
		return d
	}
	if err := d.Set("description", r.Description); err != nil {
		return d
	}
	if err := d.Set("enabled_var", r.EnabledVar); err != nil {
		return d
	}
	if err := d.Set("enabled_value", r.EnabledValue); err != nil {
		return d
	}
	if err := d.Set("overwrite", r.Overwrite); err != nil {
		return d
	}
	if err := d.Set("overwrite_vars", r.OverwriteVars); err != nil {
		return d
	}
	if err := d.Set("update_on_launch", r.UpdateOnLaunch); err != nil {
		return d
	}
	if err := d.Set("inventory_id", r.Inventory); err != nil {
		return d
	}
	if err := d.Set("credential_id", r.Credential); err != nil {
		return d
	}
	if err := d.Set("source", r.Source); err != nil {
		return d
	}
	if err := d.Set("source_vars", normalizeJsonYaml(r.SourceVars)); err != nil {
		return d
	}
	if err := d.Set("host_filter", r.HostFilter); err != nil {
		return d
	}
	if err := d.Set("update_cache_timeout", r.UpdateCacheTimeout); err != nil {
		return d
	}
	if err := d.Set("verbosity", r.Verbosity); err != nil {
		return d
	}
	if err := d.Set("source_project_id", r.SourceProject); err != nil {
		return d
	}
	if err := d.Set("source_path", r.SourcePath); err != nil {
		return d
	}
	// obsolete schema added so terraform doesn't break
	// these don't do anything in later versions of AWX! Update your code.
	if err := d.Set("source_regions", r.SourceRegions); err != nil {
		return d
	}
	if err := d.Set("instance_filters", r.InstanceFilters); err != nil {
		return d
	}
	if err := d.Set("group_by", r.GroupBy); err != nil {
		return d
	}
	if err := d.Set("source_project_id", r.SourceProject); err != nil {
		return d
	}
	if err := d.Set("source_path", r.SourcePath); err != nil {
		return d
	}

	return d
}
