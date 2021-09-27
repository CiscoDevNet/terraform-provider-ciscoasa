package ciscoasa

import (
	"fmt"
	"testing"

	"github.com/CiscoDevNet/go-ciscoasa/ciscoasa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCiscoASATimeRange_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASATimeRangeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoTimeRanges_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASATimeRangeExists([]string{"ciscoasa_timerange.tr_no_periodic_test",
						"ciscoasa_timerange.tr_periodic_test"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_no_periodic_test", "name", "tr_no_periodic_test"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_no_periodic_test", "value.0.end", "never"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_no_periodic_test", "value.0.start", "now"),

					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "name", "tr_periodic_test"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.end", "03:47 May 14 2025"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.start", "now"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.0.end_hour", "1"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.0.end_minute", "1"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.0.frequency", "Tuesday Thursday Saturday "),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.0.start_hour", "0"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.0.start_minute", "0"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.1.end_hour", "5"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.1.end_minute", "1"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.1.frequency", "daily"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.1.start_hour", "4"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.1.start_minute", "0"),
				),
			},
		},
	})
}

func TestAccCiscoASATimeRange_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASATimeRangeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoTimeRanges_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASATimeRangeExists([]string{"ciscoasa_timerange.tr_no_periodic_test",
						"ciscoasa_timerange.tr_periodic_test"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_no_periodic_test", "name", "tr_no_periodic_test"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_no_periodic_test", "value.0.end", "never"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_no_periodic_test", "value.0.start", "now"),

					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "name", "tr_periodic_test"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.end", "03:47 May 14 2025"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.start", "now"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.0.end_hour", "1"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.0.end_minute", "1"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.0.frequency", "Tuesday Thursday Saturday "),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.0.start_hour", "0"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.0.start_minute", "0"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.1.end_hour", "5"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.1.end_minute", "1"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.1.frequency", "daily"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.1.start_hour", "4"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.1.start_minute", "0"),
				),
			},
			{
				Config: testAccCiscoTimeRanges_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASATimeRangeExists([]string{"ciscoasa_timerange.tr_no_periodic_test",
						"ciscoasa_timerange.tr_periodic_test"}),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_no_periodic_test", "name", "tr_no_periodic_test"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_no_periodic_test", "value.0.end", "03:47 May 14 2026"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_no_periodic_test", "value.0.start", "03:47 May 14 2025"),

					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "name", "tr_periodic_test"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.end", "never"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.start", "now"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.0.end_hour", "1"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.0.end_minute", "1"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.0.frequency", "weekdays"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.0.start_hour", "0"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.0.start_minute", "0"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.1.end_hour", "5"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.1.end_minute", "1"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.1.frequency", "weekend"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.1.start_hour", "4"),
					resource.TestCheckResourceAttr(
						"ciscoasa_timerange.tr_periodic_test", "value.0.periodic.1.start_minute", "0"),
				),
			},
		},
	})
}

func testAccCheckCiscoASATimeRangeExists(resnames []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, n := range resnames {
			rs, ok := s.RootModule().Resources[n]
			if !ok {
				return fmt.Errorf("Not found: %s", n)
			}

			if rs.Primary.ID == "" {
				return fmt.Errorf("Time Range ID not set for %s", n)
			}
			ca := testAccProvider.Meta().(*ciscoasa.Client)
			o, err := ca.Objects.GetTimeRange(rs.Primary.ID)
			if err != nil {
				return err
			}

			if o.ObjectID != rs.Primary.ID {
				return fmt.Errorf("Time Range %s not found", n)
			}
		}
		return nil
	}
}

func testAccCheckCiscsoASATimeRangeDestroy(s *terraform.State) error {
	ca := testAccProvider.Meta().(*ciscoasa.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ciscoasa_timerange" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Time Range ID not set")
		}

		_, err := ca.Objects.GetTimeRange(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Time Range %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

var testAccCiscoTimeRanges_basic = fmt.Sprintf(`
resource "ciscoasa_timerange" "tr_no_periodic_test" {
  name = "tr_no_periodic_test"
  value {
    start = "now"
    end   = "never"
  }
}

resource "ciscoasa_timerange" "tr_periodic_test" {
  name = "tr_periodic_test"
  value {
    start = "now"
    end   = "03:47 May 14 2025"
    periodic {
      frequency    = "Tuesday Thursday Saturday "
      start_hour   = 0
      start_minute = 0
      end_hour     = 1
      end_minute   = 1
    }
    periodic {
      frequency    = "daily"
      start_hour   = 4
      start_minute = 0
      end_hour     = 5
      end_minute   = 1
    }
  }
}
`)

var testAccCiscoTimeRanges_update = fmt.Sprintf(`
resource "ciscoasa_timerange" "tr_no_periodic_test" {
  name = "tr_no_periodic_test"
  value {
    start = "03:47 May 14 2025"
    end   = "03:47 May 14 2026"
  }
}

resource "ciscoasa_timerange" "tr_periodic_test" {
  name = "tr_periodic_test"
  value {
    start = "now"
    end   = "never"
    periodic {
      frequency    = "weekdays"
      start_hour   = 0
      start_minute = 0
      end_hour     = 1
      end_minute   = 1
    }
    periodic {
      frequency    = "weekend"
      start_hour   = 4
      start_minute = 0
      end_hour     = 5
      end_minute   = 1
    }
  }
}
`)
