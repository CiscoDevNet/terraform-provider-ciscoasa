package ciscoasa

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/CiscoDevNet/go-ciscoasa/ciscoasa"
)

func TestAccCiscoASADhcpRelayLocal_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASADhcpRelayLocalDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoDhcpRelayLocals_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASADhcpRelayLocalExists([]string{"ciscoasa_dhcp_relay_local.dhcp_relay_test"}),
					resource.TestCheckResourceAttrSet(
						"ciscoasa_dhcp_relay_local.dhcp_relay_test", "interface"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_relay_local.dhcp_relay_test", "servers.0", "166.177.180.190"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_relay_local.dhcp_relay_test", "servers.1", "20.20.30.26"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_relay_local.dhcp_relay_test", "servers.2", "135.144.153.163"),
				),
			},
		},
	})
}

func TestAccCiscoASADhcpRelayLocal_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASADhcpRelayLocalDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoDhcpRelayLocals_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASADhcpRelayLocalExists([]string{"ciscoasa_dhcp_relay_local.dhcp_relay_test"}),
					resource.TestCheckResourceAttrSet(
						"ciscoasa_dhcp_relay_local.dhcp_relay_test", "interface"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_relay_local.dhcp_relay_test", "servers.0", "166.177.180.190"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_relay_local.dhcp_relay_test", "servers.1", "20.20.30.26"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_relay_local.dhcp_relay_test", "servers.2", "135.144.153.163"),
				),
			},
			{
				Config: testAccCiscoDhcpRelayLocals_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASADhcpRelayLocalExists([]string{"ciscoasa_dhcp_relay_local.dhcp_relay_test"}),
					resource.TestCheckResourceAttrSet(
						"ciscoasa_dhcp_relay_local.dhcp_relay_test", "interface"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_relay_local.dhcp_relay_test", "servers.0", "20.20.30.27"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_relay_local.dhcp_relay_test", "servers.1", "166.177.180.191"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_relay_local.dhcp_relay_test", "servers.2", "135.144.153.164"),
				),
			},
		},
	})
}

func testAccCheckCiscoASADhcpRelayLocalExists(resnames []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, n := range resnames {
			rs, ok := s.RootModule().Resources[n]
			if !ok {
				return fmt.Errorf("Not found: %s", n)
			}

			if rs.Primary.ID == "" {
				return fmt.Errorf("DHCP Relay Server interface ID not set for %s", n)
			}
			ca := testAccProvider.Meta().(*ciscoasa.Client)
			o, err := ca.Dhcp.GetDhcpRelayLocal(rs.Primary.ID)
			if err != nil {
				return err
			}

			if o.Interface != rs.Primary.ID {
				return fmt.Errorf("DHCP Relay Server interface %s not found", n)
			}
		}
		return nil
	}
}

func testAccCheckCiscsoASADhcpRelayLocalDestroy(s *terraform.State) error {
	ca := testAccProvider.Meta().(*ciscoasa.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ciscoasa_dhcp_relay_local" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("DHCP Relay Server interface ID not set")
		}

		_, err := ca.Dhcp.GetDhcpRelayLocal(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("DHCP Relay Server interface %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

var testAccCiscoDhcpRelayLocals_basic = fmt.Sprintf(`
resource "ciscoasa_interface_physical" "ipv4_static_physical_interface" {
  hardware_id    = "%s/%s"
	name = "%s"
  ip_address {
    static {
      ip       = "192.168.10.6"
      net_mask = "255.255.255.0"
    }
  }
  security_level = 15
}

resource "ciscoasa_dhcp_relay_local" "dhcp_relay_test" {
  interface = ciscoasa_interface_physical.ipv4_static_physical_interface.name
  servers = [
    "166.177.180.190",
    "20.20.30.26",
    "135.144.153.163"
  ]
}`,
	CISCOASA_INTERFACE_HW_ID_BASE,
	CISCOASA_INTERFACE_HW_IDS[0],
	CISCOASA_INTERFACE_NAME)

var testAccCiscoDhcpRelayLocals_update = fmt.Sprintf(`
resource "ciscoasa_interface_physical" "ipv4_static_physical_interface" {
  hardware_id    = "%s/%s"
	name = "%s"
  ip_address {
    static {
      ip       = "192.168.10.6"
      net_mask = "255.255.255.0"
    }
  }
  security_level = 15
}

resource "ciscoasa_dhcp_relay_local" "dhcp_relay_test" {
  interface = ciscoasa_interface_physical.ipv4_static_physical_interface.name
  servers = [
    "20.20.30.27",
		"166.177.180.191",
    "135.144.153.164"
  ]
}`,
	CISCOASA_INTERFACE_HW_ID_BASE,
	CISCOASA_INTERFACE_HW_IDS[0],
	CISCOASA_INTERFACE_NAME)
