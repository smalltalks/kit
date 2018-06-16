package meta

import (
	"fmt"

	"github.com/fatih/color"
)

// Banner represent the iconic logo/text of a software
type Banner string

// Description prepending the banner on top of a description
func (b Banner) Description(desc string) string {
	return fmt.Sprintf("%s\n%s", b, desc)
}

var pkgBanner = Banner(color.BlueString(`
 ___  __    ___  _________
|\  \|\  \ |\  \|\___   ___\
\ \  \/  /|\ \  \|___ \  \_|
 \ \   ___  \ \  \   \ \  \
  \ \  \\ \  \ \  \   \ \  \
   \ \__\\ \__\ \__\   \ \__\
    \|__| \|__|\|__|    \|__|  ©smalltalks
`))

// GetBanner return the banner of KIT
func GetBanner() Banner {
	return pkgBanner
}
