package meta

import "fmt"

var pkgBanner = fmt.Sprintf(`
 ___  __    ___  _________
|\  \|\  \ |\  \|\___   ___\
\ \  \/  /|\ \  \|___ \  \_|
 \ \   ___  \ \  \   \ \  \
  \ \  \\ \  \ \  \   \ \  \
   \ \__\\ \__\ \__\   \ \__\
    \|__| \|__|\|__|    \|__| %s
`, GetVersion())

// GetBanner return the banner of KIT
func GetBanner() string {
	return pkgBanner
}
