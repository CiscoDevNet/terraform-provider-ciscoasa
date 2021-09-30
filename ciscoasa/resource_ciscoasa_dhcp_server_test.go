package ciscoasa

import (
	"fmt"
	"testing"

	"github.com/CiscoDevNet/go-ciscoasa/ciscoasa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCiscoASADhcpServer_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASADhcpServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoDhcpServers_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASADhcpServerExists([]string{"ciscoasa_dhcp_server.dhcp_test"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "auto_config_enabled", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "auto_config_interface", "tengig"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "ddns_override_client_settings", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "ddns_update_both_records", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "ddns_update_dns_client", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "dns_ip_primary", "3.3.3.3"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "dns_ip_secondary", "5.5.5.5"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "domain_name", "testing1"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "enabled", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "interface", "test"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "lease_length", "305"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "options.0.code", "2"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "options.0.type", "hex"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "options.0.value1", "c52f"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "options.1.code", "4"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "options.1.type", "ascii"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "options.1.value1", "1261"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "options.2.code", "13"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "options.2.type", "ip"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "options.2.value1", "1.1.1.2"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "options.2.value2", "1.1.2.1"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "ping_timeout", "40"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "pool_end_ip", "8.8.8.20"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "pool_start_ip", "8.8.8.4"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "vpn_override", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "wins_ip_primary", "4.4.4.4"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "wins_ip_secondary", "6.6.6.6"),
				),
			},
		},
	})
}

func TestAccCiscoASADhcpServer_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASADhcpServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoDhcpServers_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASADhcpServerExists([]string{"ciscoasa_dhcp_server.dhcp_test"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "auto_config_enabled", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "auto_config_interface", "tengig"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "ddns_override_client_settings", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "ddns_update_both_records", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "ddns_update_dns_client", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "dns_ip_primary", "3.3.3.3"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "dns_ip_secondary", "5.5.5.5"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "domain_name", "testing1"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "enabled", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "interface", "test"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "lease_length", "305"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "options.0.code", "2"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "options.0.type", "hex"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "options.0.value1", "c52f"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "options.1.code", "4"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "options.1.type", "ascii"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "options.1.value1", "1261"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "options.2.code", "13"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "options.2.type", "ip"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "options.2.value1", "1.1.1.2"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "options.2.value2", "1.1.2.1"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "ping_timeout", "40"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "pool_end_ip", "8.8.8.20"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "pool_start_ip", "8.8.8.4"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "vpn_override", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "wins_ip_primary", "4.4.4.4"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "wins_ip_secondary", "6.6.6.6"),
				),
			},
			{
				Config: testAccCiscoDhcpServers_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASADhcpServerExists([]string{"ciscoasa_dhcp_server.dhcp_test"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "auto_config_enabled", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "auto_config_interface", "tengig"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "ddns_override_client_settings", "false"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "ddns_update_both_records", "false"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "ddns_update_dns_client", "false"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "dns_ip_primary", ""),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "dns_ip_secondary", ""),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "domain_name", ""),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "enabled", "false"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "interface", "test"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "lease_length", ""),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "ping_timeout", ""),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "pool_end_ip", ""),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "pool_start_ip", ""),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "vpn_override", "false"),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "wins_ip_primary", ""),
					resource.TestCheckResourceAttr(
						"ciscoasa_dhcp_server.dhcp_test", "wins_ip_secondary", ""),
				),
			},
		},
	})
}

func testAccCheckCiscoASADhcpServerExists(resnames []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, n := range resnames {
			rs, ok := s.RootModule().Resources[n]
			if !ok {
				return fmt.Errorf("Not found: %s", n)
			}

			if rs.Primary.ID == "" {
				return fmt.Errorf("DHCP Server ID not set for %s", n)
			}
			ca := testAccProvider.Meta().(*ciscoasa.Client)
			o, err := ca.Dhcp.GetDhcpServer(rs.Primary.ID)
			if err != nil {
				return err
			}

			if o.ObjectId != rs.Primary.ID {
				return fmt.Errorf("DHCP Server %s not found", n)
			}
		}
		return nil
	}
}

func testAccCheckCiscsoASADhcpServerDestroy(s *terraform.State) error {
	ca := testAccProvider.Meta().(*ciscoasa.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ciscoasa_dhcp_server" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("DHCP Server ID not set")
		}

		_, err := ca.Dhcp.GetDhcpServer(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("DHCP Server %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

var testAccCiscoDhcpServers_basic = fmt.Sprintf(`
resource "ciscoasa_interface_physical" "ipv4_interface" {
  name           = "test"
  hardware_id    = "%s/%s"
  interface_desc = "test descr"
  ip_address {
    static {
      ip       = "8.8.8.1"
      net_mask = "255.255.255.0"
    }
  }
  security_level = 0
}

resource "ciscoasa_interface_physical" "tengig_ipv4_dhcp" {
  name           = "tengig"
  hardware_id    = "%s/%s"
  interface_desc = "Interface DHCP"
  ip_address {
    dhcp {
      dhcp_option_using_mac = false
      dhcp_broadcast        = true
      dhcp_client {
        set_default_route = true
        metric            = 1
        primary_track_id  = -1
        tracking_enabled  = false
      }
    }
  }
  security_level = 0
}


resource "ciscoasa_dhcp_server" "dhcp_test" {
  interface             = ciscoasa_interface_physical.ipv4_interface.name
  enabled               = true
  pool_start_ip         = "8.8.8.4"
  pool_end_ip           = "8.8.8.20"
  dns_ip_primary        = "3.3.3.3"
  dns_ip_secondary      = "5.5.5.5"
  wins_ip_primary       = "4.4.4.4"
  wins_ip_secondary     = "6.6.6.6"
  lease_length          = "305"
  ping_timeout          = "40"
  domain_name           = "testing1"
  auto_config_enabled   = true
  auto_config_interface = ciscoasa_interface_physical.tengig_ipv4_dhcp.name
  vpn_override          = true
  options {
    type   = "hex"
    code   = 2
    value1 = "c52f"
  }
  options {
    type   = "ascii"
    code   = 4
    value1 = "1261"
  }
  options {
    type   = "ip"
    code   = 13
    value1 = "1.1.1.2"
    value2 = "1.1.2.1"
  }
  ddns_update_dns_client        = true
  ddns_update_both_records      = true
  ddns_override_client_settings = true
}`,
	CISCOASA_INTERFACE_HW_ID_BASE,
	CISCOASA_INTERFACE_HW_IDS[0],
	CISCOASA_INTERFACE_HW_ID_BASE,
	CISCOASA_INTERFACE_HW_IDS[1])

var testAccCiscoDhcpServers_update = fmt.Sprintf(`
	resource "ciscoasa_interface_physical" "ipv4_interface" {
		name           = "test"
		hardware_id    = "%s/%s"
		interface_desc = "test descr"
		ip_address {
			static {
				ip       = "8.8.8.1"
				net_mask = "255.255.255.0"
			}
		}
		security_level = 0
	}

	resource "ciscoasa_interface_physical" "tengig_ipv4_dhcp" {
		name           = "tengig"
		hardware_id    = "%s/%s"
		interface_desc = "Interface DHCP"
		ip_address {
			dhcp {
				dhcp_option_using_mac = false
				dhcp_broadcast        = true
				dhcp_client {
					set_default_route = true
					metric            = 1
					primary_track_id  = -1
					tracking_enabled  = false
				}
			}
		}
		security_level = 0
	}
	
	
	resource "ciscoasa_dhcp_server" "dhcp_test" {
		interface             = ciscoasa_interface_physical.ipv4_interface.name
		auto_config_enabled   = true
		auto_config_interface = ciscoasa_interface_physical.tengig_ipv4_dhcp.name
}`,
	CISCOASA_INTERFACE_HW_ID_BASE,
	CISCOASA_INTERFACE_HW_IDS[0],
	CISCOASA_INTERFACE_HW_ID_BASE,
	CISCOASA_INTERFACE_HW_IDS[1])
