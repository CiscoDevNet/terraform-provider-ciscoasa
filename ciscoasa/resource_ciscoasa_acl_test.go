package ciscoasa

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
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
						"ciscoasa_acl.foo", "rule.1891598391.source", "192.168.10.0/23"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1891598391.destination", "192.168.12.0/23"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1891598391.destination_service", "icmp/0"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1891598391.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1891598391.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1860076939.source", "192.168.10.5"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1860076939.destination", "192.168.15.0/25"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1860076939.destination_service", "tcp/443"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1860076939.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1860076939.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3786451327.source", "192.168.10.0/24"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3786451327.source_service", "udp"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3786451327.destination", "192.168.15.6"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3786451327.destination_service", "udp/53"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3786451327.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3786451327.permit", "true"),
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
						"ciscoasa_acl.foo", "rule.1891598391.source", "192.168.10.0/23"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1891598391.destination", "192.168.12.0/23"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1891598391.destination_service", "icmp/0"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1891598391.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1891598391.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1860076939.source", "192.168.10.5"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1860076939.destination", "192.168.15.0/25"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1860076939.destination_service", "tcp/443"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1860076939.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.1860076939.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3786451327.source", "192.168.10.0/24"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3786451327.source_service", "udp"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3786451327.destination", "192.168.15.6"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3786451327.destination_service", "udp/53"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3786451327.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3786451327.permit", "true"),
				),
			},

			resource.TestStep{
				Config: testAccCiscoASAACL_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASAACLExists("ciscoasa_acl.foo"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.2718879164.source", "192.168.10.0/23"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.2718879164.destination", "192.168.12.0/24"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.2718879164.destination_service", "icmp/8"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.2718879164.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.2718879164.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3651644684.source", "192.168.12.0/24"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3651644684.source_service", "tcp"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3651644684.destination", "192.168.15.16"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3651644684.destination_service", "tcp/53"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3651644684.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3651644684.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3295053340.source", "192.168.10.0/24"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3295053340.destination", "192.168.15.0/25"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3295053340.destination_service", "tcp/443"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3295053340.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_acl.foo", "rule.3295053340.permit", "true"),
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
    "source" = "192.168.10.5"
    "destination" = "192.168.15.0/25"
    "destination_service" = "tcp/443"
  }
  rule {
    "source" = "192.168.10.0/24"
    "source_service" = "udp"
    "destination" = "192.168.15.6"
    "destination_service" = "udp/53"
  }
  rule {
    "source" = "192.168.10.0/23"
    "destination" = "192.168.12.0/23"
    "destination_service" = "icmp/0"
  }
}`,
	CISCOASA_OBJECT_PREFIX)

var testAccCiscoASAACL_update = fmt.Sprintf(`
resource "ciscoasa_acl" "foo" {
  name = "%s"
  rule {
    "source" = "192.168.10.0/24"
    "destination" = "192.168.15.0/25"
    "destination_service" = "tcp/443"
  }
  rule {
    "source" = "192.168.12.0/24"
    "source_service" = "tcp"
    "destination" = "192.168.15.16"
    "destination_service" = "tcp/53"
  }
  rule {
    "source" = "192.168.10.0/23"
    "destination" = "192.168.12.0/24"
    "destination_service" = "icmp/8"
  }
}`,
	CISCOASA_OBJECT_PREFIX)
