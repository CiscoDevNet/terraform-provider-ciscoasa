package ciscoasa

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

func TestAccCiscoASAAccessInRules_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASAAccessInRulesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoASAAccessInRules_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASAAccessInRulesExists("ciscoasa_access_in_rules.foo"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.#", "3"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.2.source", "192.168.10.0/23"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.2.destination", "192.168.12.0/23"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.2.destination_service", "icmp/0"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.2.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.2.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.0.source", "192.168.10.5/32"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.0.destination", "192.168.15.0/25"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.0.destination_service", "tcp/443"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.0.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.0.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.1.source", "192.168.10.0/24"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.1.source_service", "udp"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.1.destination", "192.168.15.6/32"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.1.destination_service", "udp/53"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.1.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.1.permit", "true"),
				),
			},
		},
	})
}

func TestAccCiscoASAAccessInRules_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASAAccessInRulesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoASAAccessInRules_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASAAccessInRulesExists("ciscoasa_access_in_rules.foo"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.2.source", "192.168.10.0/23"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.2.destination", "192.168.12.0/23"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.2.destination_service", "icmp/0"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.2.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.2.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.0.source", "192.168.10.5/32"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.0.destination", "192.168.15.0/25"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.0.destination_service", "tcp/443"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.0.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.0.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.1.source", "192.168.10.0/24"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.1.source_service", "udp"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.1.destination", "192.168.15.6/32"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.1.destination_service", "udp/53"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.1.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.1.permit", "true"),
				),
			},
			{
				Config: testAccCiscoASAAccessInRules_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASAAccessInRulesExists("ciscoasa_access_in_rules.foo"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.2.source", "192.168.10.0/23"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.2.destination", "192.168.12.0/24"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.2.destination_service", "icmp/8"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.2.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.2.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.1.source", "192.168.12.0/24"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.1.source_service", "tcp"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.1.destination", "192.168.15.16/32"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.1.destination_service", "tcp/53"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.1.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.1.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.0.source", "192.168.10.0/24"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.0.destination", "192.168.15.0/25"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.0.destination_service", "tcp/443"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.0.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_in_rules.foo", "rule.0.permit", "true"),
				),
			},
		},
	})
}

func testAccCheckCiscoASAAccessInRulesExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No interface ID is set")
		}

		ca := testAccProvider.Meta().(*ciscoasa.Client)
		l, err := ca.Access.ListAccessInRules(rs.Primary.ID)

		if err != nil {
			return err
		}

		if len(l.Items) == 0 {
			return fmt.Errorf("No rules for interface %s found", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckCiscsoASAAccessInRulesDestroy(s *terraform.State) error {
	ca := testAccProvider.Meta().(*ciscoasa.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ciscoasa_access_in_rules" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No interface ID is set")
		}

		_, err := ca.Access.ListAccessInRules(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Rules for interface %s still exist", rs.Primary.ID)
		}

	}

	return nil
}

var testAccCiscoASAAccessInRules_basic = fmt.Sprintf(`
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

resource "ciscoasa_access_in_rules" "foo" {
  interface = ciscoasa_interface_physical.ipv4_static_physical_interface.name
  rule {
	source = "192.168.10.5/32"
    destination = "192.168.15.0/25"
    destination_service = "tcp/443"
  }
  rule {
    source = "192.168.10.0/24"
	source_service = "udp"
    destination = "192.168.15.6/32"
    destination_service = "udp/53"
  }
  rule {
    source = "192.168.10.0/23"
    destination = "192.168.12.0/23"
    destination_service = "icmp/0"
  }
}`,
	CISCOASA_INTERFACE_HW_ID_BASE,
	CISCOASA_INTERFACE_HW_IDS[0],
	CISCOASA_INTERFACE_NAME)

var testAccCiscoASAAccessInRules_update = fmt.Sprintf(`
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

resource "ciscoasa_access_in_rules" "foo" {
  interface = ciscoasa_interface_physical.ipv4_static_physical_interface.name
  rule {
    source = "192.168.10.0/24"
    destination = "192.168.15.0/25"
    destination_service = "tcp/443"
  }
  rule {
    source = "192.168.12.0/24"
	source_service = "tcp"
    destination = "192.168.15.16/32"
    destination_service = "tcp/53"
  }
  rule {
    source = "192.168.10.0/23"
    destination = "192.168.12.0/24"
    destination_service = "icmp/8"
  }
}`,
	CISCOASA_INTERFACE_HW_ID_BASE,
	CISCOASA_INTERFACE_HW_IDS[0],
	CISCOASA_INTERFACE_NAME)
