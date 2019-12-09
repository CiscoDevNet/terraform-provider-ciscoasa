package ciscoasa

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

func TestAccCiscoASANetworkObjects(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoAsaNetworkObjectDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCiscoAsaNetworkObjects,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoAsaNetworkObjectExists([]string{"ciscoasa_network_object.obj_ipv4host",
						"ciscoasa_network_object.obj_ipv4range",
						"ciscoasa_network_object.obj_ipv4subnet"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_network_object.obj_ipv4host", "value", "192.168.10.5"),
					resource.TestCheckResourceAttr(
						"ciscoasa_network_object.obj_ipv4range", "value", "192.168.10.5-192.168.10.15"),
					resource.TestCheckResourceAttr(
						"ciscoasa_network_object.obj_ipv4subnet", "value", "192.168.10.128/25"),
				),
			},
		},
	})
}

func testAccCheckCiscoAsaNetworkObjectExists(resnames []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, n := range resnames {
			rs, ok := s.RootModule().Resources[n]
			if !ok {
				return fmt.Errorf("Not found: %s", n)
			}

			if rs.Primary.ID == "" {
				return fmt.Errorf("Network Object ID not set for %s", n)
			}
			ca := testAccProvider.Meta().(*ciscoasa.Client)
			o, err := ca.Objects.GetNetworkObject(rs.Primary.ID)
			if err != nil {
				return err
			}

			if o.ObjectID != rs.Primary.ID {
				return fmt.Errorf("Network Object %s not found", rs.Primary.ID)
			}
		}
		return nil
	}
}

func testAccCheckCiscsoAsaNetworkObjectDestroy(s *terraform.State) error {
	ca := testAccProvider.Meta().(*ciscoasa.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ciscoasa_network_object" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Network Object ID not set")
		}

		_, err := ca.Objects.GetNetworkObject(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Network Object %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

var testAccCiscoAsaNetworkObjects = fmt.Sprintf(`
resource "ciscoasa_network_object" "obj_ipv4host" {
  name = "%s_ipv4host"
  value = "192.168.10.5"
}
resource "ciscoasa_network_object" "obj_ipv4range" {
  name = "%s_ipv4range"
  value = "192.168.10.5-192.168.10.15"
}
resource "ciscoasa_network_object" "obj_ipv4subnet" {
  name = "%s_ipv4subnet"
  value = "192.168.10.128/25"
}`,
	CISCOASA_OBJECT_PREFIX,
	CISCOASA_OBJECT_PREFIX,
	CISCOASA_OBJECT_PREFIX)
