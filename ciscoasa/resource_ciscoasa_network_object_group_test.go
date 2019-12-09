package ciscoasa

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

func TestAccCiscoASANetworkObjectGroup(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoAsaNetworkObjectGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCiscoAsaNetworkObjectGroup,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoAsaNetworkObjectGroupExists([]string{
						"ciscoasa_network_object_group.objgrp_mixed",
						"ciscoasa_network_object_group.objgrp_nested",
					}),
					resource.TestCheckResourceAttr(
						"ciscoasa_network_object_group.objgrp_mixed", "members.#", "3"),
					resource.TestCheckResourceAttr(
						"ciscoasa_network_object_group.objgrp_mixed", "members.1086291305", "192.168.10.15"),
					resource.TestCheckResourceAttr(
						"ciscoasa_network_object_group.objgrp_mixed", "members.2041400964", "10.5.10.0/24"),
					resource.TestCheckResourceAttr(
						"ciscoasa_network_object_group.objgrp_nested", "members.#", "3"),
					resource.TestCheckResourceAttr(
						"ciscoasa_network_object_group.objgrp_nested", "members.3744480902", "192.168.20.14"),
					resource.TestCheckResourceAttr(
						"ciscoasa_network_object_group.objgrp_nested", "members.233334120", "10.25.10.0/24"),
				),
			},
		},
	})
}

func testAccCheckCiscoAsaNetworkObjectGroupExists(resnames []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, n := range resnames {
			rs, ok := s.RootModule().Resources[n]
			if !ok {
				return fmt.Errorf("Not found: %s", n)
			}

			if rs.Primary.ID == "" {
				return fmt.Errorf("Network Object Group ID not set for %s", n)
			}
			ca := testAccProvider.Meta().(*ciscoasa.Client)
			o, err := ca.Objects.GetNetworkObjectGroup(rs.Primary.ID)
			if err != nil {
				return err
			}

			if o.ObjectID != rs.Primary.ID {
				return fmt.Errorf("Network Object Group %s not found", rs.Primary.ID)
			}
		}
		return nil
	}
}

func testAccCheckCiscsoAsaNetworkObjectGroupDestroy(s *terraform.State) error {
	ca := testAccProvider.Meta().(*ciscoasa.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ciscoasa_network_object_group" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Network Object Group ID not set")
		}

		_, err := ca.Objects.GetNetworkObjectGroup(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Network Object Group %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

var testAccCiscoAsaNetworkObjectGroup = fmt.Sprintf(`
resource "ciscoasa_network_object" "obj_ipv4host" {
  name = "%s_ipv4host"
  value = "192.168.10.5"
}

resource "ciscoasa_network_object_group" "objgrp_mixed" {
  name = "%s_objgrp_mixed"
  members = ["${ciscoasa_network_object.obj_ipv4host.name}",
  			 "192.168.10.15",
  			 "10.5.10.0/24"]
}

resource "ciscoasa_network_object_group" "objgrp_nested" {
  name = "%s_objgrp_nested"
  members = ["${ciscoasa_network_object_group.objgrp_mixed.name}",
  			 "192.168.20.14",
  			 "10.25.10.0/24"]
}`,
	CISCOASA_OBJECT_PREFIX,
	CISCOASA_OBJECT_PREFIX,
	CISCOASA_OBJECT_PREFIX)
