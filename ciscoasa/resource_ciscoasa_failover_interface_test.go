package ciscoasa

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

func TestAccCiscoASAFailoverInterface_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASAFailoverInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoFailoverInterface_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASAFailoverInterfaceExists([]string{"ciscoasa_failover_interface.inside"}),
					resource.TestCheckResourceAttrSet(
						"ciscoasa_failover_interface.inside", "hardware_id"),
					resource.TestCheckResourceAttr(
						"ciscoasa_failover_interface.inside", "monitored", "false"),
					resource.TestCheckResourceAttr(
						"ciscoasa_failover_interface.inside", "standby_ip", "192.168.10.10"),
				),
			},
		},
	})
}

func TestAccCiscoASAFailoverInterface_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASAFailoverInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoFailoverInterface_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASAFailoverInterfaceExists([]string{"ciscoasa_failover_interface.inside"}),
					resource.TestCheckResourceAttrSet(
						"ciscoasa_failover_interface.inside", "hardware_id"),
					resource.TestCheckResourceAttr(
						"ciscoasa_failover_interface.inside", "monitored", "false"),
					resource.TestCheckResourceAttr(
						"ciscoasa_failover_interface.inside", "standby_ip", "192.168.10.10"),
				),
			},
			{
				Config: testAccCiscoFailoverInterface_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASAFailoverInterfaceExists([]string{"ciscoasa_failover_interface.inside"}),
					resource.TestCheckResourceAttrSet(
						"ciscoasa_failover_interface.inside", "hardware_id"),
					resource.TestCheckResourceAttr(
						"ciscoasa_failover_interface.inside", "monitored", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_failover_interface.inside", "standby_ip", "192.168.10.20"),
				),
			},
		},
	})
}

func testAccCheckCiscoASAFailoverInterfaceExists(resnames []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, n := range resnames {
			rs, ok := s.RootModule().Resources[n]
			if !ok {
				return fmt.Errorf("Not found: %s", n)
			}

			if rs.Primary.ID == "" {
				return fmt.Errorf("FailoverInterface ID not set for %s", n)
			}
			ca := testAccProvider.Meta().(*ciscoasa.Client)
			o, err := ca.Failover.GetFailoverInterface(rs.Primary.ID)
			if err != nil {
				return err
			}

			if o.ObjectId != rs.Primary.ID {
				return fmt.Errorf("FailoverInterface %s not found", n)
			}
		}
		return nil
	}
}

func testAccCheckCiscsoASAFailoverInterfaceDestroy(s *terraform.State) error {
	// There is no destroy available for this object, it will be there as long as Interface is there. Hence just mocking destroy.
	return nil
}

var testAccCiscoFailoverInterface_basic = fmt.Sprintf(`
resource "ciscoasa_interface_physical" "inside" {
  name        = "inside"
  hardware_id = "%s/%s"
  ip_address {
    static {
      ip       = "192.168.10.1"
      net_mask = "255.255.255.0"
    }
  }
  shutdown       = false
  security_level = 100
}

resource "ciscoasa_network_object" "inside_host" {
  name  = "inside-host"
  value = "192.168.10.10"
}

resource "ciscoasa_failover_interface" "inside" {
  hardware_id = ciscoasa_interface_physical.inside.hardware_id
  standby_ip  = ciscoasa_network_object.inside_host.value
  monitored   = false
}`,
	CISCOASA_INTERFACE_HW_ID_BASE,
	CISCOASA_INTERFACE_HW_IDS[0])

var testAccCiscoFailoverInterface_update = fmt.Sprintf(`
resource "ciscoasa_interface_physical" "inside" {
  name        = "inside"
  hardware_id = "%s/%s"
  ip_address {
    static {
      ip       = "192.168.10.1"
      net_mask = "255.255.255.0"
    }
  }
  shutdown       = false
  security_level = 100
}

resource "ciscoasa_network_object" "inside_host" {
  name  = "inside-host"
  value = "192.168.10.20"
}

resource "ciscoasa_failover_interface" "inside" {
  hardware_id = ciscoasa_interface_physical.inside.hardware_id
  standby_ip  = ciscoasa_network_object.inside_host.value
  monitored   = true
}`,
	CISCOASA_INTERFACE_HW_ID_BASE,
	CISCOASA_INTERFACE_HW_IDS[0])
