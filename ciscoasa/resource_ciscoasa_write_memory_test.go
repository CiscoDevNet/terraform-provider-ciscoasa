package ciscoasa

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccCiscoASAWriteMemory(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASAWriteMemoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoWriteMemory,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASAWriteMemoryExists([]string{"ciscoasa_write_memory.write_memory"}),
				),
			},
		},
	})
}

func testAccCheckCiscoASAWriteMemoryExists(resourceNames []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, n := range resourceNames {
			rs, ok := s.RootModule().Resources[n]
			if !ok {
				return fmt.Errorf("Not found: %s", n)
			}

			if rs.Primary.ID == "" {
				return fmt.Errorf("Write Memory ID not set for %s", n)
			}
		}
		return nil
	}
}

func testAccCheckCiscsoASAWriteMemoryDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ciscoasa_write_memory" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Write Memory resource ID was not set")
		}
	}

	return nil
}

var testAccCiscoWriteMemory = `
resource "ciscoasa_write_memory" "write_memory" { 
    triggers = {}
}`
