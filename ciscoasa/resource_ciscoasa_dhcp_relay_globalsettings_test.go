package ciscoasa

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

func TestAccCiscoASADhcpRelayGlobalsettings_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoDhcpRelayGlobalsettings_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASADhcpRelayGlobalsettingsExists([]string{"ciscoasa_dhcp_relay_globalsettings.test"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_relay_globalsettings.test", "ipv4_timeout", "90"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_relay_globalsettings.test", "ipv6_timeout", "90"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_relay_globalsettings.test", "trusted_on_all_interfaces", "true"),
				),
			},
		},
	})
}

func TestAccCiscoASADhcpRelayGlobalsettings_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoDhcpRelayGlobalsettings_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASADhcpRelayGlobalsettingsExists([]string{"ciscoasa_dhcp_relay_globalsettings.test"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_relay_globalsettings.test", "ipv4_timeout", "90"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_relay_globalsettings.test", "ipv6_timeout", "90"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_relay_globalsettings.test", "trusted_on_all_interfaces", "true"),
				),
			},
			{
				Config: testAccCiscoDhcpRelayGlobalsettings_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASADhcpRelayGlobalsettingsExists([]string{"ciscoasa_dhcp_relay_globalsettings.test"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_relay_globalsettings.test", "ipv4_timeout", "180"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_relay_globalsettings.test", "ipv6_timeout", "180"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_relay_globalsettings.test", "trusted_on_all_interfaces", "false"),
				),
			},
		},
	})
}

func testAccCheckCiscoASADhcpRelayGlobalsettingsExists(resnames []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, n := range resnames {
			rs, ok := s.RootModule().Resources[n]
			if !ok {
				return fmt.Errorf("Not found: %s", n)
			}

			if rs.Primary.ID == "" {
				return fmt.Errorf("DHCP Relay Globalsettings ID not set for %s", n)
			}
			ca := testAccProvider.Meta().(*ciscoasa.Client)
			_, err := ca.Dhcp.GetDhcpRelayGlobalsettings()
			if err != nil {
				return err
			}
		}
		return nil
	}
}

var testAccCiscoDhcpRelayGlobalsettings_basic = fmt.Sprintf(`
resource "ciscoasa_dhcp_relay_globalsettings" "test" {
  ipv4_timeout              = 90
  ipv6_timeout              = 90
  trusted_on_all_interfaces = true
}
`)

var testAccCiscoDhcpRelayGlobalsettings_update = fmt.Sprintf(`
resource "ciscoasa_dhcp_relay_globalsettings" "test" {
  ipv4_timeout              = 180
  ipv6_timeout              = 180
  trusted_on_all_interfaces = false
}
`)
