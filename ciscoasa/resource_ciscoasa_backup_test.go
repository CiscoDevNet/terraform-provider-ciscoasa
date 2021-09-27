package ciscoasa

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCiscoASABackup(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoBackups,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASABackupExists([]string{"ciscoasa_backup.test"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_backup.test", "id", "full"),
					resource.TestCheckResourceAttr(
						"ciscoasa_backup.test", "passphrase", "123456"),
					resource.TestCheckResourceAttr(
						"ciscoasa_backup.test", "location", "disk0:/backup.cfg"),
				),
			},
		},
	})
}

// Here is just a mockup of real exists function, because there is no API for getting Backup object
func testAccCheckCiscoASABackupExists(resnames []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, n := range resnames {
			rs, ok := s.RootModule().Resources[n]
			if !ok {
				return fmt.Errorf("Not found: %s", n)
			}

			if rs.Primary.ID == "" {
				return fmt.Errorf("Backup ID not set for %s", n)
			}
		}
		return nil
	}
}

var testAccCiscoBackups = fmt.Sprintf(`
resource "ciscoasa_backup" "test" {
  passphrase = "123456"
  location   = "disk0:/backup.cfg"
}
`)
