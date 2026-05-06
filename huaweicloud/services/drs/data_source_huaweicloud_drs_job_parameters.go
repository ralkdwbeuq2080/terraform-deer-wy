package drs

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/chnsz/golangsdk"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/utils"
)

// @API DRS GET /v5/{project_id}/jobs/{job_id}/configurations
func DataSourceDrsJobParameters() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDrsJobParametersRead,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Specifies the region in which to query the resource.",
			},
			"job_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the ID of the DRS job.",
			},
			"offset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the offset of query results.",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Specifies the maximum number of query results.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies the parameter name for filtering.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total number of parameters.",
			},
			"parameter_config_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the parameter.",
						},
						"value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The value of the parameter.",
						},
						"default_value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The default value of the parameter.",
						},
						"value_range": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The value range of the parameter.",
						},
						"is_need_restart": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the parameter needs a restart.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the parameter.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time of the parameter.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Update time of the parameter.",
						},
					},
				},
			},
		},
	}
}

func buildJobParametersQueryParams(d *schema.ResourceData) string {
	params := url.Values{}
	if v, ok := d.GetOk("offset"); ok {
		params.Add("offset", fmt.Sprintf("%d", v.(int)))
	}
	if v, ok := d.GetOk("limit"); ok {
		params.Add("limit", fmt.Sprintf("%d", v.(int)))
	}
	if v, ok := d.GetOk("name"); ok {
		params.Add("name", v.(string))
	}

	if len(params) > 0 {
		return "?" + params.Encode()
	}
	return ""
}

func dataSourceDrsJobParametersRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conf := meta.(*config.Config)
	region := conf.GetRegion(d)
	client, err := conf.DrsV5Client(region)
	if err != nil {
		return diag.Errorf("error creating DRS v5 client, error: %s", err)
	}

	listParamsHttpUrl := "v5/{project_id}/jobs/{job_id}/configurations"
	listParamsPath := client.Endpoint + listParamsHttpUrl
	listParamsPath = strings.ReplaceAll(listParamsPath, "{project_id}", client.ProjectID)
	listParamsPath = strings.ReplaceAll(listParamsPath, "{job_id}", d.Get("job_id").(string))
	listParamsPath += buildJobParametersQueryParams(d)

	listParamsOpt := golangsdk.RequestOpts{
		KeepResponseBody: true,
	}

	listParamsResp, err := client.Request("GET", listParamsPath, &listParamsOpt)
	if err != nil {
		return diag.FromErr(err)
	}
	listParamsRespBody, err := utils.FlattenResponse(listParamsResp)
	if err != nil {
		return diag.FromErr(err)
	}

	uuid, err := uuid.GenerateUUID()
	if err != nil {
		return diag.Errorf("unable to generate ID: %s", err)
	}
	d.SetId(uuid)

	mErr := multierror.Append(nil,
		d.Set("region", region),
		d.Set("total_count", utils.PathSearch("total_count", listParamsRespBody, 0)),
		d.Set("parameter_config_list", flattenJobParameters(listParamsRespBody)),
	)

	return diag.FromErr(mErr.ErrorOrNil())
}

func flattenJobParameters(respBody interface{}) []interface{} {
	configList := utils.PathSearch("parameter_config_list", respBody, make([]interface{}, 0)).([]interface{})
	if len(configList) == 0 {
		return nil
	}

	rst := make([]interface{}, 0, len(configList))
	for _, v := range configList {
		rst = append(rst, map[string]interface{}{
			"name":            utils.PathSearch("name", v, nil),
			"value":           utils.PathSearch("value", v, nil),
			"default_value":   utils.PathSearch("default_value", v, nil),
			"value_range":     utils.PathSearch("value_range", v, nil),
			"is_need_restart": utils.PathSearch("is_need_restart", v, false),
			"description":     utils.PathSearch("description", v, nil),
			"created_at":      utils.PathSearch("created_at", v, nil),
			"updated_at":      utils.PathSearch("updated_at", v, nil),
		})
	}
	return rst
}
