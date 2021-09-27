package ciscoasa

import (
	"fmt"
	"testing"

	"github.com/CiscoDevNet/go-ciscoasa/ciscoasa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCiscoASALicenseConfig_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASALicenseConfigDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoLicenseConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASALicenseConfigExists([]string{"ciscoasa_license_config.test"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_license_config.test", "license_server_url", "https://tools.cisco.com/its/service/oddce/services/DDCEService"),
					resource.TestCheckResourceAttr(
						"ciscoasa_license_config.test", "privacy_host_name", "false"),
					resource.TestCheckResourceAttr(
						"ciscoasa_license_config.test", "privacy_version", "false"),
					resource.TestCheckResourceAttr(
						"ciscoasa_license_config.test", "throughput", "2G"),
					resource.TestCheckResourceAttr(
						"ciscoasa_license_config.test", "transport_url", "true"),
				),
			},
		},
	})
}

func TestAccCiscoASALicenseConfig_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASALicenseConfigDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoLicenseConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASALicenseConfigExists([]string{"ciscoasa_license_config.test"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_license_config.test", "license_server_url", "https://tools.cisco.com/its/service/oddce/services/DDCEService"),
					resource.TestCheckResourceAttr(
						"ciscoasa_license_config.test", "privacy_host_name", "false"),
					resource.TestCheckResourceAttr(
						"ciscoasa_license_config.test", "privacy_version", "false"),
					resource.TestCheckResourceAttr(
						"ciscoasa_license_config.test", "throughput", "2G"),
					resource.TestCheckResourceAttr(
						"ciscoasa_license_config.test", "transport_url", "true"),
				),
			},
			{
				Config: testAccCiscoLicenseConfig_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASALicenseConfigExists([]string{"ciscoasa_license_config.test"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_license_config.test", "license_server_url", "https://tools.cisco.com/its/service/oddce/services/DDCEService"),
					resource.TestCheckResourceAttr(
						"ciscoasa_license_config.test", "privacy_host_name", "false"),
					resource.TestCheckResourceAttr(
						"ciscoasa_license_config.test", "privacy_version", "false"),
					resource.TestCheckResourceAttr(
						"ciscoasa_license_config.test", "throughput", "1G"),
					resource.TestCheckResourceAttr(
						"ciscoasa_license_config.test", "transport_url", "true"),
				),
			},
		},
	})
}

func testAccCheckCiscoASALicenseConfigExists(resnames []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, n := range resnames {
			rs, ok := s.RootModule().Resources[n]
			if !ok {
				return fmt.Errorf("Not found: %s", n)
			}

			if rs.Primary.ID == "" {
				return fmt.Errorf("LicenseConfig ID not set for %s", n)
			}
			ca := testAccProvider.Meta().(*ciscoasa.Client)
			_, err := ca.Licensing.GetLicenseConfig()
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func testAccCheckCiscsoASALicenseConfigDestroy(s *terraform.State) error {
	// There is no destroy available for this object, it will be there always. Hence just mocking destroy.
	return nil
}

var testAccCiscoLicenseConfig_basic = `
resource "ciscoasa_license_config" "test" {
  throughput         = "2G"
  license_server_url = "https://tools.cisco.com/its/service/oddce/services/DDCEService"
}`

var testAccCiscoLicenseConfig_update = `
resource "ciscoasa_license_config" "test" {
  throughput         = "1G"
  license_server_url = "https://tools.cisco.com/its/service/oddce/services/DDCEService"
}`
