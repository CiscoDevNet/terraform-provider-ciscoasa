package ciscoasa

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

func TestAccCiscoASANetworkServiceGroup(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoAsaNetworkServiceGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCiscoAsaNetworkServiceGroup,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoAsaNetworkServiceGroupExists([]string{
						"ciscoasa_network_service_group.srvgrp_mixed",
						"ciscoasa_network_service_group.srvgrp_nested",
					}),

					resource.TestCheckResourceAttr(
						"ciscoasa_network_service_group.srvgrp_mixed", "members.#", "4"),
					resource.TestCheckResourceAttr(
						"ciscoasa_network_service_group.srvgrp_mixed", "members.1318937322", "tcp/80"),
					resource.TestCheckResourceAttr(
						"ciscoasa_network_service_group.srvgrp_mixed", "members.2028646336", "tcp/6001-6500"),
					resource.TestCheckResourceAttr(
						"ciscoasa_network_service_group.srvgrp_mixed", "members.3592067446", "udp/53"),
					resource.TestCheckResourceAttr(
						"ciscoasa_network_service_group.srvgrp_mixed", "members.1524852438", "icmp/0"),

					resource.TestCheckResourceAttr(
						"ciscoasa_network_service_group.srvgrp_nested", "members.#", "2"),
					resource.TestCheckResourceAttr(
						"ciscoasa_network_service_group.srvgrp_nested", "members.2453335262", "icmp/8"),
				),
			},
		},
	})
}

func testAccCheckCiscoAsaNetworkServiceGroupExists(resnames []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, n := range resnames {
			rs, ok := s.RootModule().Resources[n]
			if !ok {
				return fmt.Errorf("Not found: %s", n)
			}

			if rs.Primary.ID == "" {
				return fmt.Errorf("Network Service Group ID not set for %s", n)
			}
			ca := testAccProvider.Meta().(*ciscoasa.Client)
			o, err := ca.Objects.GetNetworkServiceGroup(rs.Primary.ID)
			if err != nil {
				return err
			}

			if o.ObjectID != rs.Primary.ID {
				return fmt.Errorf("Network Service Group %s not found", rs.Primary.ID)
			}
		}
		return nil
	}
}

func testAccCheckCiscsoAsaNetworkServiceGroupDestroy(s *terraform.State) error {
	ca := testAccProvider.Meta().(*ciscoasa.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ciscoasa_network_service_group" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Network Service Group ID not set")
		}

		_, err := ca.Objects.GetNetworkServiceGroup(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Network Object Group %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

var testAccCiscoAsaNetworkServiceGroup = fmt.Sprintf(`
resource "ciscoasa_network_service_group" "srvgrp_mixed" {
  name = "%s_srvgrp_mixed"
  members = ["tcp/80",
             "udp/53",
             "tcp/6001-6500",
             "icmp/0"]
}

resource "ciscoasa_network_service_group" "srvgrp_nested" {
  name = "%s_srvgrp_nested"
  members = ["${ciscoasa_network_service_group.srvgrp_mixed.name}",
             "icmp/8"]
}`,
	CISCOASA_OBJECT_PREFIX,
	CISCOASA_OBJECT_PREFIX)
