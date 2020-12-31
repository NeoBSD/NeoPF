package freebsd

// System details a FreeBSD version
type System struct {
	Hostname          string             `json:"hostname"`
	Machine           string             `json:"machine"`
	MachineArch       string             `json:"machine_arch"`
	Hardware          Hardware           `json:"hardware"`
	Version           Version            `json:"version"`
	NetworkInterfaces []NetworkInterface `json:"network_interfaces"`
}

// GetSystemInfo returns general system infos
func GetSystemInfo() (System, error) {
	hostname, err := GetHostname()
	if err != nil {
		return System{}, err
	}

	machine, err := getCommandOutput("uname", "-m")
	if err != nil {
		return System{}, err
	}

	machineArch, err := getCommandOutput("uname", "-p")
	if err != nil {
		return System{}, err
	}

	hardware, err := GetHardwareInfo()
	if err != nil {
		return System{}, err
	}

	version, err := GetVersionInfo()
	if err != nil {
		return System{}, err
	}

	interfaces, err := GetNetworkInterfaces()
	if err != nil {
		return System{}, err
	}

	system := System{
		Hostname:          hostname,
		Machine:           machine,
		MachineArch:       machineArch,
		Hardware:          hardware,
		Version:           version,
		NetworkInterfaces: interfaces,
	}

	return system, nil
}
