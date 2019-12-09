package ciscoasa

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

func TestAccCiscoASAStaticRoute(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASAStaticRouteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCiscoStaticRoutes,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASAStaticRouteExists([]string{"ciscoasa_static_route.ipv4_static_route",
						"ciscoasa_static_route.ipv6_static_route"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_static_route.ipv4_static_route", "gateway", "192.168.10.20"),
					resource.TestCheckResourceAttr(
						"ciscoasa_static_route.ipv4_static_route", "network", "10.254.0.0/16"),
					resource.TestCheckResourceAttr(
						"ciscoasa_static_route.ipv4_static_route", "metric", "1"),
					resource.TestCheckResourceAttr(
						"ciscoasa_static_route.ipv4_static_route", "tracked", "false"),
					resource.TestCheckResourceAttr(
						"ciscoasa_static_route.ipv4_static_route", "tunneled", "false"),

					resource.TestCheckResourceAttr(
						"ciscoasa_static_route.ipv6_static_route", "gateway", "fd01:1338::1"),
					resource.TestCheckResourceAttr(
						"ciscoasa_static_route.ipv6_static_route", "network", "fd01:1337::/64"),
					resource.TestCheckResourceAttr(
						"ciscoasa_static_route.ipv6_static_route", "metric", "1"),
					resource.TestCheckResourceAttr(
						"ciscoasa_static_route.ipv6_static_route", "tracked", "false"),
					resource.TestCheckResourceAttr(
						"ciscoasa_static_route.ipv6_static_route", "tunneled", "false"),
				),
			},
		},
	})
}

func testAccCheckCiscoASAStaticRouteExists(resnames []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, n := range resnames {
			rs, ok := s.RootModule().Resources[n]
			if !ok {
				return fmt.Errorf("Not found: %s", n)
			}

			if rs.Primary.ID == "" {
				return fmt.Errorf("Static Route ID not set for %s", n)
			}
			ca := testAccProvider.Meta().(*ciscoasa.Client)
			o, err := ca.Routing.GetStaticRoute(rs.Primary.ID)
			if err != nil {
				return err
			}

			if o.ObjectID != rs.Primary.ID {
				return fmt.Errorf("Static Route %s not found", n)
			}
		}
		return nil
	}
}

func testAccCheckCiscsoASAStaticRouteDestroy(s *terraform.State) error {
	ca := testAccProvider.Meta().(*ciscoasa.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ciscoasa_static_route" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Static Route ID not set")
		}

		_, err := ca.Routing.GetStaticRoute(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Static Route %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

var testAccCiscoStaticRoutes = fmt.Sprintf(`
resource "ciscoasa_static_route" "ipv4_static_route" {
  interface = "%s"
  network = "10.254.0.0/16"
  gateway = "192.168.10.20"
}

resource "ciscoasa_static_route" "ipv6_static_route" {
  interface = "%s"
  network = "fd01:1337::/64"
  gateway = "fd01:1338::1"
}`,
	CISCOASA_INTERFACE_NAME,
	CISCOASA_INTERFACE_NAME)
