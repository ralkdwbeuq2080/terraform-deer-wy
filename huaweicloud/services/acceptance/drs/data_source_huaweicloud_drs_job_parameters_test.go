package drs

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func TestAccDataSourceDrsJobParameters_basic(t *testing.T) {
	dataSourceName := "data.huaweicloud_drs_job_parameters.test"
	dc := acceptance.InitDataSourceCheck(dataSourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDrsJobParameters_basic,
				Check: resource.ComposeTestCheckFunc(
					dc.CheckResourceExists(),
					resource.TestCheckResourceAttrSet(dataSourceName, "total_count"),
					resource.TestCheckResourceAttrSet(dataSourceName, "parameter_config_list.#"),
					resource.TestCheckResourceAttrSet(dataSourceName, "parameter_config_list.0.name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "parameter_config_list.0.value"),
				),
			},
		},
	})
}

const testAccDataSourceDrsJobParameters_basic string = `
data "huaweicloud_drs_job_parameters" "test" {
  job_id = "52e440b9-e89a-45b8-b520-b8e0482jb201"
}
`
