package ciscoasa

import (
	"fmt"
	"testing"

	"github.com/CiscoDevNet/go-ciscoasa/ciscoasa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCiscoASANat_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASANatDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoNat_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASANatExists([]string{"ciscoasa_nat.auto_test"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "description", "static auto test"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "extended", "false"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "flat", "false"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "mode", "static"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "original_interface_name", "inside"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "original_source_value", "inside-host"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "section", "auto"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "translated_interface_name", "outside"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "translated_source_value", "inside-host-translated"),
				),
			},
		},
	})
}

func TestAccCiscoASANat_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASANatDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoNat_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASANatExists([]string{"ciscoasa_nat.auto_test"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "description", "static auto test"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "extended", "false"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "flat", "false"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "mode", "static"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "original_interface_name", "inside"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "original_source_value", "inside-host"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "section", "auto"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "translated_interface_name", "outside"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "translated_source_value", "inside-host-translated"),
				),
			},
			{
				Config: testAccCiscoNat_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASANatExists([]string{"ciscoasa_nat.auto_test"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "description", "static auto test updated"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "extended", "false"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "flat", "false"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "mode", "static"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "original_interface_name", "inside"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "original_source_value", "inside-host"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "section", "auto"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "translated_interface_name", "outside"),
					resource.TestCheckResourceAttr(
						"ciscoasa_nat.auto_test", "translated_source_value", "inside-host-translated"),
				),
			},
		},
	})
}

func testAccCheckCiscoASANatExists(resnames []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, n := range resnames {
			rs, ok := s.RootModule().Resources[n]
			if !ok {
				return fmt.Errorf("Not found: %s", n)
			}

			if rs.Primary.ID == "" {
				return fmt.Errorf("NAT ID not set for %s", n)
			}
			ca := testAccProvider.Meta().(*ciscoasa.Client)
			o, err := ca.Nat.GetNat("auto", rs.Primary.ID)
			if err != nil {
				return err
			}

			if o.OriginalSource.ObjectId != rs.Primary.ID {
				return fmt.Errorf("NAT %s not found", n)
			}
		}
		return nil
	}
}

func testAccCheckCiscsoASANatDestroy(s *terraform.State) error {
	ca := testAccProvider.Meta().(*ciscoasa.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ciscoasa_nat" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("NAT ID not set")
		}

		_, err := ca.Nat.GetNat("auto", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("NAT %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

var testAccCiscoNat_basic = fmt.Sprintf(`
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

resource "ciscoasa_interface_physical" "outside" {
  name        = "outside"
  hardware_id = "%s/%s"
  ip_address {
    static {
      ip       = "209.209.209.1"
      net_mask = "255.255.255.252"
    }
  }
  shutdown       = false
  security_level = 50
}

resource "ciscoasa_network_object" "inside_host" {
  name  = "inside-host"
  value = "192.168.10.10"
}

resource "ciscoasa_network_object" "inside_host_translated" {
  name  = "inside-host-translated"
  value = "208.208.208.1"
}

resource "ciscoasa_nat" "auto_test" {
  section                   = "auto"
  description               = "static auto test"
  mode                      = "static"
  original_interface_name   = ciscoasa_interface_physical.inside.name
  translated_interface_name = ciscoasa_interface_physical.outside.name
  original_source_kind      = "objectRef#NetworkObj"
  original_source_value     = ciscoasa_network_object.inside_host.name
  translated_source_kind    = "objectRef#NetworkObj"
  translated_source_value   = ciscoasa_network_object.inside_host_translated.name
}`,
	CISCOASA_INTERFACE_HW_ID_BASE,
	CISCOASA_INTERFACE_HW_IDS[0],
	CISCOASA_INTERFACE_HW_ID_BASE,
	CISCOASA_INTERFACE_HW_IDS[1])

var testAccCiscoNat_update = fmt.Sprintf(`
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

resource "ciscoasa_interface_physical" "outside" {
  name        = "outside"
  hardware_id = "%s/%s"
  ip_address {
    static {
      ip       = "209.209.209.1"
      net_mask = "255.255.255.252"
    }
  }
  shutdown       = false
  security_level = 50
}

resource "ciscoasa_network_object" "inside_host" {
  name  = "inside-host"
  value = "192.168.10.10"
}

resource "ciscoasa_network_object" "inside_host_translated" {
  name  = "inside-host-translated"
  value = "208.208.208.1"
}

resource "ciscoasa_nat" "auto_test" {
  section                   = "auto"
  description               = "static auto test updated"
  mode                      = "static"
  original_interface_name   = ciscoasa_interface_physical.inside.name
  translated_interface_name = ciscoasa_interface_physical.outside.name
  original_source_kind      = "objectRef#NetworkObj"
  original_source_value     = ciscoasa_network_object.inside_host.name
  translated_source_kind    = "objectRef#NetworkObj"
  translated_source_value   = ciscoasa_network_object.inside_host_translated.name
}`,
	CISCOASA_INTERFACE_HW_ID_BASE,
	CISCOASA_INTERFACE_HW_IDS[0],
	CISCOASA_INTERFACE_HW_ID_BASE,
	CISCOASA_INTERFACE_HW_IDS[1])
