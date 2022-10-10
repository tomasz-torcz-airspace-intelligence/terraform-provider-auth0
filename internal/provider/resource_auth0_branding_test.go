package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/auth0/terraform-provider-auth0/internal/recorder"
)

const testAccBrandingConfigCreate = `
resource "auth0_branding" "my_brand" {
	logo_url = "https://mycompany.org/v1/logo.png"
	favicon_url = "https://mycompany.org/favicon.ico"
}
`

const testAccBrandingConfigUpdateAllFields = `
resource "auth0_branding" "my_brand" {
	logo_url = "https://mycompany.org/v2/logo.png"
	favicon_url = "https://mycompany.org/favicon.ico"

	colors {
		primary = "#0059d6"
		page_background = "#000000"
	}

	font {
		url = "https://mycompany.org/font/myfont.ttf"
	}

	universal_login {
		body = "<!DOCTYPE html><html><head>{%- auth0:head -%}</head><body>{%- auth0:widget -%}</body></html>"
	}
}
`

const testAccBrandingConfigUpdateAgain = `
resource "auth0_branding" "my_brand" {
	logo_url = "https://mycompany.org/v3/logo.png"
	favicon_url = "https://mycompany.org/favicon.ico"

	colors {
		primary = "#0059d6"
	}

	font {
		url = "https://mycompany.org/font/myfont.ttf"
	}

	universal_login {
		# Setting this to an empty string should
		# not trigger any API call, so the value
		# stays the same as the previous scenario.
		body = ""
	}
}
`

const testAccBrandingConfigReset = `
resource "auth0_branding" "my_brand" {
	logo_url = "https://mycompany.org/v1/logo.png"
	favicon_url = "https://mycompany.org/favicon.ico"
}
`

func TestAccBranding(t *testing.T) {
	httpRecorder := recorder.New(t)

	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviders(httpRecorder),
		Steps: []resource.TestStep{
			{
				Config: testAccBrandingConfigCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "logo_url", "https://mycompany.org/v1/logo.png"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "favicon_url", "https://mycompany.org/favicon.ico"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "colors.#", "0"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "font.#", "0"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "universal_login.#", "0"),
				),
			},
			{
				Config: testAccBrandingConfigUpdateAllFields,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "logo_url", "https://mycompany.org/v2/logo.png"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "favicon_url", "https://mycompany.org/favicon.ico"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "colors.#", "1"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "colors.0.primary", "#0059d6"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "colors.0.page_background", "#000000"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "font.#", "1"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "font.0.url", "https://mycompany.org/font/myfont.ttf"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "universal_login.#", "1"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "universal_login.0.body", "<!DOCTYPE html><html><head>{%- auth0:head -%}</head><body>{%- auth0:widget -%}</body></html>"),
				),
			},
			{
				Config: testAccBrandingConfigUpdateAgain,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "logo_url", "https://mycompany.org/v3/logo.png"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "favicon_url", "https://mycompany.org/favicon.ico"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "colors.#", "1"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "colors.0.primary", "#0059d6"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "colors.0.page_background", "#000000"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "font.#", "1"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "font.0.url", "https://mycompany.org/font/myfont.ttf"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "universal_login.#", "1"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "universal_login.0.body", "<!DOCTYPE html><html><head>{%- auth0:head -%}</head><body>{%- auth0:widget -%}</body></html>"),
				),
			},
			{
				Config: testAccBrandingConfigReset,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "logo_url", "https://mycompany.org/v1/logo.png"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "favicon_url", "https://mycompany.org/favicon.ico"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "colors.#", "0"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "font.#", "0"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "universal_login.#", "0"),
				),
			},
		},
	})
}