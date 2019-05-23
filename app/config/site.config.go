package config

// Addr Addr ...
type Addr struct {
	Login  string
	Logout string
	Home   string
}

// Site for Site
type Site struct {
	Addr
}

// SiteConfig info
func SiteConfig() *Site {
	/*
		return &Site{
			Addr{
				Login:  "https://account.iuu.pub/login",
				Logout: "https://account.iuu.pub/login",
			},
		}
	*/
	return &Site{
		Addr{
			Login:  "http://127.0.0.1/account/login",
			Logout: "http://127.0.0.1/account/login",
			Home:   "http://127.0.0.1/account/home",
		},
	}
}
