package natureremo

type ApplianceModel struct {
	ID           string `json:"id"`
	Manufacturer string `json:"manufacturer"`
	RemoteName   string `json:"remote_name"`
	Name         string `json:"name"`
	Image        string `json:"image"`
}

type Appliance struct {
	ID     string          `json:"id"`
	Device *DeviceCore     `json:"device"`
	Model  *ApplianceModel `json:"model"`
	// TODO
}