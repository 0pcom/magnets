# Maintainer: 0pcom
pkgname=haltingstate
pkgver=0.0.1
pkgrel=1
pkgdesc="haltingstate.net website - telegram chat export random quote generator"
arch=('any')
url="https://haltingstate.net"
license=('GPL')
groups=()
depends=('go' 'ansifilter')
source=("haltingstate.service")

package() {
	install -Dm644 "${srcdir}/haltingstate.service" "${pkgdir}/usr/lib/systemd/system/haltingstate.service"
	install -Dm644 "${srcdir}/haltingstate.service" "${pkgdir}/etc/skel/.config/systemd/user/haltingstate.service"
}
sha256sums=('4de4e9e6ebee08e1f3d8e0c3434d8d02fbdad812e6a318d489d3d25cc0293bb7')
