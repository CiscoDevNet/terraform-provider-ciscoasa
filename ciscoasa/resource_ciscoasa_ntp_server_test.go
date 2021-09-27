package ciscoasa

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

func TestAccCiscoASANtpServer_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASANtpServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoNtpServers_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASANtpServerExists([]string{"ciscoasa_ntp_server.ntp_test"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_ntp_server.ntp_test", "ip_address", "2.2.2.2"),
					resource.TestCheckResourceAttrSet(
						"ciscoasa_ntp_server.ntp_test", "interface"),
					resource.TestCheckResourceAttr(
						"ciscoasa_ntp_server.ntp_test", "preferred", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_ntp_server.ntp_test", "key_number", "3"),
					resource.TestCheckResourceAttr(
						"ciscoasa_ntp_server.ntp_test", "key_value", "test3"),
					resource.TestCheckResourceAttr(
						"ciscoasa_ntp_server.ntp_test", "key_trusted", "false"),
				),
			},
		},
	})
}

func TestAccCiscoASANtpServer_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASANtpServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoNtpServers_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASANtpServerExists([]string{"ciscoasa_ntp_server.ntp_test"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_ntp_server.ntp_test", "ip_address", "2.2.2.2"),
					resource.TestCheckResourceAttrSet(
						"ciscoasa_ntp_server.ntp_test", "interface"),
					resource.TestCheckResourceAttr(
						"ciscoasa_ntp_server.ntp_test", "preferred", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_ntp_server.ntp_test", "key_number", "3"),
					resource.TestCheckResourceAttr(
						"ciscoasa_ntp_server.ntp_test", "key_value", "test3"),
					resource.TestCheckResourceAttr(
						"ciscoasa_ntp_server.ntp_test", "key_trusted", "false"),
				),
			},
			{
				Config: testAccCiscoNtpServers_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASANtpServerExists([]string{"ciscoasa_ntp_server.ntp_test"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_ntp_server.ntp_test", "ip_address", "3.3.3.3"),
					resource.TestCheckResourceAttrSet(
						"ciscoasa_ntp_server.ntp_test", "interface"),
					resource.TestCheckResourceAttr(
						"ciscoasa_ntp_server.ntp_test", "preferred", "false"),
					resource.TestCheckResourceAttr(
						"ciscoasa_ntp_server.ntp_test", "key_number", "4"),
					resource.TestCheckResourceAttr(
						"ciscoasa_ntp_server.ntp_test", "key_value", "test4"),
					resource.TestCheckResourceAttr(
						"ciscoasa_ntp_server.ntp_test", "key_trusted", "false"),
				),
			},
		},
	})
}

func testAccCheckCiscoASANtpServerExists(resnames []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, n := range resnames {
			rs, ok := s.RootModule().Resources[n]
			if !ok {
				return fmt.Errorf("Not found: %s", n)
			}

			if rs.Primary.ID == "" {
				return fmt.Errorf("NTP Server ID not set for %s", n)
			}
			ca := testAccProvider.Meta().(*ciscoasa.Client)
			o, err := ca.DeviceSetup.GetNtpServer(rs.Primary.ID)
			if err != nil {
				return err
			}

			if o.IpAddress != rs.Primary.ID {
				return fmt.Errorf("NTP Server %s not found", n)
			}
		}
		return nil
	}
}

func testAccCheckCiscsoASANtpServerDestroy(s *terraform.State) error {
	ca := testAccProvider.Meta().(*ciscoasa.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ciscoasa_ntp_server" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("NTP Server ID not set")
		}

		_, err := ca.DeviceSetup.GetNtpServer(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("NTP Server %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

var testAccCiscoNtpServers_basic = fmt.Sprintf(`
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

resource "ciscoasa_ntp_server" "ntp_test" {
  ip_address = "2.2.2.2"
  interface  = ciscoasa_interface_physical.ipv4_static_physical_interface.name
  key_number = "3"
  key_value  = "test3"
  preferred  = true
}`,
	CISCOASA_INTERFACE_HW_ID_BASE,
	CISCOASA_INTERFACE_HW_IDS[0],
	CISCOASA_INTERFACE_NAME)

var testAccCiscoNtpServers_update = fmt.Sprintf(`
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

resource "ciscoasa_ntp_server" "ntp_test" {
  ip_address = "3.3.3.3"
  interface  = ciscoasa_interface_physical.ipv4_static_physical_interface.name
  key_number = "4"
  key_value  = "test4"
  preferred  = false
}`,
	CISCOASA_INTERFACE_HW_ID_BASE,
	CISCOASA_INTERFACE_HW_IDS[0],
	CISCOASA_INTERFACE_NAME)
