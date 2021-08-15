package freeroam

type UDPConfig struct {
	ListenAddress string
}

type FMSConfig struct {
	ListenAddress  string
	AllowedOrigin  string
	UpdateInterval int
}

type TLSConfig struct {
	EnforceTLS     bool
	PublicKeyFile  string
	PrivateKeyFile string
}

type Config struct {
	UDP UDPConfig
	FMS FMSConfig
	TLS TLSConfig
}

func DefaultConfig() Config {
	return Config{
		UDP: UDPConfig{
			ListenAddress: ":9999",
		},
		FMS: FMSConfig{
			ListenAddress: "127.0.0.1:6996",
			AllowedOrigin: "127.0.0.1",
		},
		TLS: TLSConfig{
			EnforceTLS: false,
			PublicKeyFile: "",
			PrivateKeyFile: "",
		},
	}
}
