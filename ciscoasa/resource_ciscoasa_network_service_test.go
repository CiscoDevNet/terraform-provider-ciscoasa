package ciscoasa

import (
	"fmt"
	"testing"

	"github.com/CiscoDevNet/go-ciscoasa/ciscoasa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCiscoASANetworkService_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASANetworkServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoNetworkServices_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASANetworkServiceExists([]string{"ciscoasa_network_service.tcp-with-source",
						"ciscoasa_network_service.icmp4"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_network_service.tcp-with-source", "name", "tcp-with-source"),
					resource.TestCheckResourceAttr(
						"ciscoasa_network_service.tcp-with-source", "value", "tcp/https,source=7000-8000"),

					resource.TestCheckResourceAttr(
						"ciscoasa_network_service.icmp4", "name", "icmp4"),
					resource.TestCheckResourceAttr(
						"ciscoasa_network_service.icmp4", "value", "icmp/traceroute/70"),
				),
			},
		},
	})
}

func TestAccCiscoASANetworkService_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASANetworkServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoNetworkServices_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASANetworkServiceExists([]string{"ciscoasa_network_service.tcp-with-source",
						"ciscoasa_network_service.icmp4"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_network_service.tcp-with-source", "name", "tcp-with-source"),
					resource.TestCheckResourceAttr(
						"ciscoasa_network_service.tcp-with-source", "value", "tcp/https,source=7000-8000"),

					resource.TestCheckResourceAttr(
						"ciscoasa_network_service.icmp4", "name", "icmp4"),
					resource.TestCheckResourceAttr(
						"ciscoasa_network_service.icmp4", "value", "icmp/traceroute/70"),
				),
			},
			{
				Config: testAccCiscoNetworkServices_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASANetworkServiceExists([]string{"ciscoasa_network_service.tcp-with-source",
						"ciscoasa_network_service.icmp4"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_network_service.tcp-with-source", "name", "tcp-with-source"),
					resource.TestCheckResourceAttr(
						"ciscoasa_network_service.tcp-with-source", "value", "tcp/http,source=44-56"),

					resource.TestCheckResourceAttr(
						"ciscoasa_network_service.icmp4", "name", "icmp4"),
					resource.TestCheckResourceAttr(
						"ciscoasa_network_service.icmp4", "value", "icmp/50/60"),
				),
			},
		},
	})
}

func testAccCheckCiscoASANetworkServiceExists(resnames []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, n := range resnames {
			rs, ok := s.RootModule().Resources[n]
			if !ok {
				return fmt.Errorf("Not found: %s", n)
			}

			if rs.Primary.ID == "" {
				return fmt.Errorf("Network Service ID not set for %s", n)
			}
			ca := testAccProvider.Meta().(*ciscoasa.Client)
			o, err := ca.Objects.GetNetworkService(rs.Primary.ID)
			if err != nil {
				return err
			}

			if o.ObjectID != rs.Primary.ID {
				return fmt.Errorf("Network Service %s not found", n)
			}
		}
		return nil
	}
}

func testAccCheckCiscsoASANetworkServiceDestroy(s *terraform.State) error {
	ca := testAccProvider.Meta().(*ciscoasa.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ciscoasa_network_service" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Network Service ID not set")
		}

		_, err := ca.Objects.GetNetworkService(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Network Service %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

var testAccCiscoNetworkServices_basic = fmt.Sprintf(`
resource "ciscoasa_network_service" "tcp-with-source" {
  name  = "tcp-with-source"
  value = "tcp/https,source=7000-8000"
}

resource "ciscoasa_network_service" "icmp4" {
  name  = "icmp4"
  value = "icmp/traceroute/70"
}`)

var testAccCiscoNetworkServices_update = fmt.Sprintf(`
resource "ciscoasa_network_service" "tcp-with-source" {
  name  = "tcp-with-source"
  value = "tcp/http,source=44-56"
}

resource "ciscoasa_network_service" "icmp4" {
  name  = "icmp4"
  value = "icmp/50/60"
}`)
