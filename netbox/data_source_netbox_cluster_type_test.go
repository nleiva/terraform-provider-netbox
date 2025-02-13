package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxClusterTypeDataSource_basic(t *testing.T) {

	testSlug := "clstrtyp_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}

data "netbox_cluster_type" "test" {
  depends_on = [netbox_cluster_type.test]
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_cluster_type.test", "cluster_type_id", "netbox_cluster_type.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_cluster_type.test", "id", "netbox_cluster_type.test", "id"),
				),
			},
		},
	})
}
