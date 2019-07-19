package config

import "testing"

func TestConfig(t *testing.T) {
	config := BaseConfig{}

	config.Title = "Test"
	config.Devices = map[string]deviceConfig{
		"sampleDevice": deviceConfig{Type: "ReturnPrinter"},
	}
}

func TestParseConfig(t *testing.T) {
	config, _ := ParseConfig("sample.cfg")
	want := "My Cool Sample Config"

	if config.Title != want {
		t.Errorf("got %q want %q", config.Title, want)
	}

	device := config.Devices["return"]

	if device.Type != "ReturnPrinter" {
		t.Errorf("Device type incorrect")
	}
}

func TestMissingConfig(t *testing.T) {
	_, err := ParseConfig("beepboop.cfg")

	if err == nil {
		t.Error("Expected error but did not get anything")
	}
}