package ciscoasa

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

func TestAccCiscoASAACL_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASAACLDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCiscoASAACL_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASAACLExists("ciscoasa_acl.foo"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.#", "3"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.0.source", "192.168.10.5/32"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.0.destination", "192.168.15.0/25"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.0.destination_service", "tcp/443"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.0.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.0.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1.source", "192.168.10.0/24"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1.source_service", "udp"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1.destination", "192.168.15.6/32"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1.destination_service", "udp/53"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.2.source", "192.168.10.0/23"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.2.destination", "192.168.12.0/23"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.2.destination_service", "icmp/0"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.2.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.2.permit", "true"),
				),
			},
		},
	})
}

func TestAccCiscoASAACL_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASAACLDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCiscoASAACL_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASAACLExists("ciscoasa_acl.foo"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.#", "3"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.0.source", "192.168.10.5/32"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.0.destination", "192.168.15.0/25"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.0.destination_service", "tcp/443"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.0.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.0.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1.source", "192.168.10.0/24"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1.source_service", "udp"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1.destination", "192.168.15.6/32"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1.destination_service", "udp/53"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.2.source", "192.168.10.0/23"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.2.destination", "192.168.12.0/23"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.2.destination_service", "icmp/0"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.2.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.2.permit", "true"),
				),
			},

			resource.TestStep{
				Config: testAccCiscoASAACL_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASAACLExists("ciscoasa_acl.foo"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.#", "5"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.0.source", "192.168.10.5/32"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.0.destination", "192.168.15.0/25"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.0.destination_service", "tcp/443"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.0.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.0.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1.source", "192.168.12.0/24"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1.source_service", "tcp"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1.destination", "192.168.15.16/32"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1.destination_service", "tcp/53"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.2.source", "0.0.0.0/0"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.2.destination", "192.168.12.0/24"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.2.destination_service", "icmp/8"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.2.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.2.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3.source", "192.168.10.0/24"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3.source_service", "udp"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3.destination", "192.168.15.6/32"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3.destination_service", "udp/53"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.4.source", "192.168.10.0/23"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.4.destination", "192.168.12.0/23"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.4.destination_service", "icmp/0"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.4.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.4.permit", "true"),
				),
			},
		},
	})
}

func testAccCheckCiscoASAACLExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Extended ACL ID is set")
		}

		ca := testAccProvider.Meta().(*ciscoasa.Client)
		o, err := ca.Objects.ListExtendedACLACEs(rs.Primary.ID)

		if err != nil {
			return err
		}

		if len(o.Items) == 0 {
			return fmt.Errorf("Extended ACL %s not found", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckCiscsoASAACLDestroy(s *terraform.State) error {
	ca := testAccProvider.Meta().(*ciscoasa.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ciscoasa_acl" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Extended ACL ID is set")
		}

		_, err := ca.Objects.ListExtendedACLACEs(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Extended ACL %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

var testAccCiscoASAACL_basic = fmt.Sprintf(`
resource "ciscoasa_acl" "foo" {
  name = "%s"
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
	CISCOASA_OBJECT_PREFIX)

var testAccCiscoASAACL_update = fmt.Sprintf(`
resource "ciscoasa_acl" "foo" {
  name = "%s"
  rule {
    source = "192.168.10.5/32"
    destination = "192.168.15.0/25"
    destination_service = "tcp/443"
  }
  rule {
    source = "192.168.12.0/24"
    source_service = "tcp"
    destination = "192.168.15.16/32"
    destination_service = "tcp/53"
    remarks = [
      "terraform-test"
    ]
  }
  rule {
    source = "0.0.0.0/0"
    destination = "192.168.12.0/24"
    destination_service = "icmp/8"
		log_interval = 10
		log_status = "Errors"
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
	CISCOASA_OBJECT_PREFIX)
