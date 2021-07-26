package flag

import (
	"github.com/spf13/pflag"
)

func (f *Flag) IsCORS() bool {
	return f.cors
}

func (f *Flag) CORS() *Flag {
	f.cors = true

	pflag.Bool("cors-allow-credentials", true, "")
	pflag.StringSlice("cors-allowed-headers", []string{
		"Origin",
		"X-Requested-With",
		"Content-Type",
		"Authorization",
		"Accept",
		"User-Agent",
		"X-Refresh-Token",
		"X-MORPH",
	}, "")
	pflag.StringSlice("cors-exposed-headers", []string{
		"Origin",
		"X-Requested-With",
		"Content-Type",
		"Accept",
		"User-Agent",
	}, "")
	pflag.StringSlice("cors-allowed-methods", []string{
		"GET",
		"POST",
		"PUT",
		"PATCH",
		"DELETE",
		"OPTIONS",
		"HEAD",
	}, "")
	pflag.StringSlice("cors-allowed-origins", []string{
		"https://0.0.0.0:3008", "https://localhost:3000", "http://localhost:4200", "http://0.0.0.0:4200",
	}, "")

	return f
}
