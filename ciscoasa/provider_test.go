package ciscoasa

import (
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

var CISCOASA_INTERFACE_HW_IDS = []string{"0", "1"}

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"ciscoasa": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("CISCOASA_API_URL"); v == "" {
		t.Fatal("CISCOASA_API_URL must be set for acceptance tests")
	}
	if v := os.Getenv("CISCOASA_USERNAME"); v == "" {
		t.Fatal("CISCOASA_USERNAME must be set for acceptance tests")
	}
	if v := os.Getenv("CISCOASA_PASSWORD"); v == "" {
		t.Fatal("CISCOASA_PASSWORD must be set for acceptance tests")
	}
	if v := os.Getenv("CISCOASA_OBJECT_PREFIX"); v == "" {
		t.Fatal("CISCOASA_OBJECT_PREFIX must be set for acceptance tests")
	}
	if v := os.Getenv("CISCOASA_INTERFACE_NAME"); v == "" {
		t.Fatal("CISCOASA_INTERFACE_NAME must be set for acceptance tests")
	}
	if v := os.Getenv("CISCOASA_INTERFACE_HW_ID_BASE"); v == "" {
		t.Fatal("CISCOASA_INTERFACE_HW_ID_BASE must be set for acceptance tests")
	}
	if v := os.Getenv("CISCOASA_INTERFACE_HW_IDS"); v != "" {
		CISCOASA_INTERFACE_HW_IDS = strings.Split(v, ",")
	}
}

// Defines a prefix used to mark resources created by the acceptance tests.
var CISCOASA_OBJECT_PREFIX = os.Getenv("CISCOASA_OBJECT_PREFIX")

// The interface name (nameif) of an existing interface.
var CISCOASA_INTERFACE_NAME = os.Getenv("CISCOASA_INTERFACE_NAME")

// The physical interface ID base of an existing interface, e.g. 'TenGigabitEthernet0'
var CISCOASA_INTERFACE_HW_ID_BASE = os.Getenv("CISCOASA_INTERFACE_HW_ID_BASE")
