package ciscoasa

import (
	"fmt"
	"testing"

	"github.com/CiscoDevNet/go-ciscoasa/ciscoasa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCiscoASAPhysicalInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASAPhysicalInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoPhysicalInterfaces,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASAPhysicalInterfaceExists([]string{"ciscoasa_interface_physical.ipv4_static_physical_interface",
						"ciscoasa_interface_physical.ipv6_static_physical_interface"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv4_static_physical_interface", "ip_address.0.static.0.ip", "192.168.10.5"),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv4_static_physical_interface", "ip_address.0.static.0.net_mask", "255.255.255.0"),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv4_static_physical_interface", "name", "ipv4_static_physical_interface"),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv4_static_physical_interface", "security_level", "0"),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv6_static_physical_interface", "ipv6_info.0.ipv6_addresses.0.address", "2001:db8:a0b:12f0::47"),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv6_static_physical_interface", "ipv6_info.0.ipv6_addresses.0.standby", "2001:db8:a0b:12f0::46"),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv6_static_physical_interface", "ipv6_info.0.ipv6_addresses.0.prefix_length", "64"),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv6_static_physical_interface", "name", "ipv6_static_physical_interface"),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv6_static_physical_interface", "security_level", "0"),
				),
			},
		},
	})
}

func TestAccCiscoASAPhysicalInterfaceUpdate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASAPhysicalInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoPhysicalInterfaces,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASAPhysicalInterfaceExists([]string{"ciscoasa_interface_physical.ipv4_static_physical_interface",
						"ciscoasa_interface_physical.ipv6_static_physical_interface"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv4_static_physical_interface", "ip_address.0.static.0.ip", "192.168.10.5"),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv4_static_physical_interface", "ip_address.0.static.0.net_mask", "255.255.255.0"),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv4_static_physical_interface", "name", "ipv4_static_physical_interface"),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv4_static_physical_interface", "security_level", "0"),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv6_static_physical_interface", "ipv6_info.0.ipv6_addresses.0.address", "2001:db8:a0b:12f0::47"),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv6_static_physical_interface", "ipv6_info.0.ipv6_addresses.0.standby", "2001:db8:a0b:12f0::46"),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv6_static_physical_interface", "ipv6_info.0.ipv6_addresses.0.prefix_length", "64"),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv6_static_physical_interface", "name", "ipv6_static_physical_interface"),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv6_static_physical_interface", "security_level", "0"),
				),
			},
			{
				Config: testAccCiscoPhysicalInterfacesUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASAPhysicalInterfaceExists([]string{"ciscoasa_interface_physical.ipv4_static_physical_interface",
						"ciscoasa_interface_physical.ipv6_static_physical_interface"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv4_static_physical_interface", "ip_address.0.static.0.ip", "192.168.10.6"),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv4_static_physical_interface", "security_level", "15"),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv6_static_physical_interface", "ipv6_info.0.link_local_address.0.address", "fe80::202:b3ff:eef1:7234"),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv6_static_physical_interface", "ipv6_info.0.link_local_address.0.standby", "fe80::202:b3ff"),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv6_static_physical_interface", "ipv6_info.0.ipv6_addresses.0.address", "2001:db8:a0b:12f0::49"),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv6_static_physical_interface", "ipv6_info.0.ipv6_addresses.0.standby", "2001:db8:a0b:12f0::48"),
					resource.TestCheckResourceAttr(
						"ciscoasa_interface_physical.ipv6_static_physical_interface", "security_level", "15"),
				),
			},
		},
	})
}

func testAccCheckCiscoASAPhysicalInterfaceExists(resnames []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, n := range resnames {
			rs, ok := s.RootModule().Resources[n]
			if !ok {
				return fmt.Errorf("Not found: %s", n)
			}

			if rs.Primary.ID == "" {
				return fmt.Errorf("Physical Interface ID not set for %s", n)
			}
			ca := testAccProvider.Meta().(*ciscoasa.Client)
			o, err := ca.Interfaces.GetPhysicalInterface(rs.Primary.ID)
			if err != nil {
				return err
			}

			if o.ObjectID != rs.Primary.ID {
				return fmt.Errorf("Physical Interface %s not found", n)
			}
		}
		return nil
	}
}

func testAccCheckCiscsoASAPhysicalInterfaceDestroy(s *terraform.State) error {
	ca := testAccProvider.Meta().(*ciscoasa.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ciscoasa_interface_physical" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Physical Interface ID not set")
		}

		// Physical Interface cannot be destroyed, hence checking for errors
		_, err := ca.Interfaces.GetPhysicalInterface(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("Error destroying Physical Interface %s: %v", rs.Primary.ID, err)
		}
	}

	return nil
}

var testAccCiscoPhysicalInterfaces = fmt.Sprintf(`
resource "ciscoasa_interface_physical" "ipv4_static_physical_interface" {
  hardware_id    = "%s/%s"
	name = "ipv4_static_physical_interface"
  ip_address {
    static {
      ip       = "192.168.10.5"
      net_mask = "255.255.255.0"
    }
  }
  security_level = 0
}

resource "ciscoasa_interface_physical" "ipv6_static_physical_interface" {
  hardware_id    = "%s/%s"
	name = "ipv6_static_physical_interface"
  ipv6_info {
    ipv6_addresses {
      address       = "2001:db8:a0b:12f0::47"
      standby       = "2001:db8:a0b:12f0::46"
      prefix_length = 64
    }
  }
  security_level = 0
}`,
	CISCOASA_INTERFACE_HW_ID_BASE,
	CISCOASA_INTERFACE_HW_IDS[0],
	CISCOASA_INTERFACE_HW_ID_BASE,
	CISCOASA_INTERFACE_HW_IDS[1])

var testAccCiscoPhysicalInterfacesUpdate = fmt.Sprintf(`
resource "ciscoasa_interface_physical" "ipv4_static_physical_interface" {
  hardware_id    = "%s/%s"
	name = "ipv4_static_physical_interface"
  ip_address {
    static {
      ip       = "192.168.10.6"
      net_mask = "255.255.255.0"
    }
  }
  security_level = 15
}

resource "ciscoasa_interface_physical" "ipv6_static_physical_interface" {
  hardware_id    = "%s/%s"
	name = "ipv6_static_physical_interface"
  ipv6_info {
		link_local_address {
			address = "fe80::202:b3ff:eef1:7234"
			standby = "fe80::202:b3ff"
		}
    ipv6_addresses {
      address       = "2001:db8:a0b:12f0::49"
      standby       = "2001:db8:a0b:12f0::48"
      prefix_length = 64
    }
  }
  security_level = 15
}`,
	CISCOASA_INTERFACE_HW_ID_BASE,
	CISCOASA_INTERFACE_HW_IDS[0],
	CISCOASA_INTERFACE_HW_ID_BASE,
	CISCOASA_INTERFACE_HW_IDS[1])
