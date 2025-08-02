// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"net/http"
	"testing"
)

var staticRoutes = []route{
	{http.MethodGet, "/"},
	{http.MethodGet, "/cmd.html"},
	{http.MethodGet, "/code.html"},
	{http.MethodGet, "/contrib.html"},
	{http.MethodGet, "/contribute.html"},
	{http.MethodGet, "/debugging_with_gdb.html"},
	{http.MethodGet, "/docs.html"},
	{http.MethodGet, "/effective_go.html"},
	{http.MethodGet, "/files.log"},
	{http.MethodGet, "/gccgo_contribute.html"},
	{http.MethodGet, "/gccgo_install.html"},
	{http.MethodGet, "/go-logo-black.png"},
	{http.MethodGet, "/go-logo-blue.png"},
	{http.MethodGet, "/go-logo-white.png"},
	{http.MethodGet, "/go1.1.html"},
	{http.MethodGet, "/go1.2.html"},
	{http.MethodGet, "/go1.html"},
	{http.MethodGet, "/go1compat.html"},
	{http.MethodGet, "/go_faq.html"},
	{http.MethodGet, "/go_mem.html"},
	{http.MethodGet, "/go_spec.html"},
	{http.MethodGet, "/help.html"},
	{http.MethodGet, "/ie.css"},
	{http.MethodGet, "/install-source.html"},
	{http.MethodGet, "/install.html"},
	{http.MethodGet, "/logo-153x55.png"},
	{http.MethodGet, "/Makefile"},
	{http.MethodGet, "/root.html"},
	{http.MethodGet, "/share.png"},
	{http.MethodGet, "/sieve.gif"},
	{http.MethodGet, "/tos.html"},
	{http.MethodGet, "/articles"},
	{http.MethodGet, "/articles/go_command.html"},
	{http.MethodGet, "/articles/index.html"},
	{http.MethodGet, "/articles/wiki"},
	{http.MethodGet, "/articles/wiki/edit.html"},
	{http.MethodGet, "/articles/wiki/final-noclosure.go"},
	{http.MethodGet, "/articles/wiki/final-noerror.go"},
	{http.MethodGet, "/articles/wiki/final-parsetemplate.go"},
	{http.MethodGet, "/articles/wiki/final-template.go"},
	{http.MethodGet, "/articles/wiki/final.go"},
	{http.MethodGet, "/articles/wiki/get.go"},
	{http.MethodGet, "/articles/wiki/http-sample.go"},
	{http.MethodGet, "/articles/wiki/index.html"},
	{http.MethodGet, "/articles/wiki/Makefile"},
	{http.MethodGet, "/articles/wiki/notemplate.go"},
	{http.MethodGet, "/articles/wiki/part1-noerror.go"},
	{http.MethodGet, "/articles/wiki/part1.go"},
	{http.MethodGet, "/articles/wiki/part2.go"},
	{http.MethodGet, "/articles/wiki/part3-errorhandling.go"},
	{http.MethodGet, "/articles/wiki/part3.go"},
	{http.MethodGet, "/articles/wiki/test.bash"},
	{http.MethodGet, "/articles/wiki/test_edit.good"},
	{http.MethodGet, "/articles/wiki/test_Test.txt.good"},
	{http.MethodGet, "/articles/wiki/test_view.good"},
	{http.MethodGet, "/articles/wiki/view.html"},
	{http.MethodGet, "/codewalk"},
	{http.MethodGet, "/codewalk/codewalk.css"},
	{http.MethodGet, "/codewalk/codewalk.js"},
	{http.MethodGet, "/codewalk/codewalk.xml"},
	{http.MethodGet, "/codewalk/functions.xml"},
	{http.MethodGet, "/codewalk/markov.go"},
	{http.MethodGet, "/codewalk/markov.xml"},
	{http.MethodGet, "/codewalk/pig.go"},
	{http.MethodGet, "/codewalk/popout.png"},
	{http.MethodGet, "/codewalk/run"},
	{http.MethodGet, "/codewalk/sharemem.xml"},
	{http.MethodGet, "/codewalk/urlpoll.go"},
	{http.MethodGet, "/devel"},
	{http.MethodGet, "/devel/release.html"},
	{http.MethodGet, "/devel/weekly.html"},
	{http.MethodGet, "/gopher"},
	{http.MethodGet, "/gopher/appenginegopher.jpg"},
	{http.MethodGet, "/gopher/appenginegophercolor.jpg"},
	{http.MethodGet, "/gopher/appenginelogo.gif"},
	{http.MethodGet, "/gopher/bumper.png"},
	{http.MethodGet, "/gopher/bumper192x108.png"},
	{http.MethodGet, "/gopher/bumper320x180.png"},
	{http.MethodGet, "/gopher/bumper480x270.png"},
	{http.MethodGet, "/gopher/bumper640x360.png"},
	{http.MethodGet, "/gopher/doc.png"},
	{http.MethodGet, "/gopher/frontpage.png"},
	{http.MethodGet, "/gopher/gopherbw.png"},
	{http.MethodGet, "/gopher/gophercolor.png"},
	{http.MethodGet, "/gopher/gophercolor16x16.png"},
	{http.MethodGet, "/gopher/help.png"},
	{http.MethodGet, "/gopher/pkg.png"},
	{http.MethodGet, "/gopher/project.png"},
	{http.MethodGet, "/gopher/ref.png"},
	{http.MethodGet, "/gopher/run.png"},
	{http.MethodGet, "/gopher/talks.png"},
	{http.MethodGet, "/gopher/pencil"},
	{http.MethodGet, "/gopher/pencil/gopherhat.jpg"},
	{http.MethodGet, "/gopher/pencil/gopherhelmet.jpg"},
	{http.MethodGet, "/gopher/pencil/gophermega.jpg"},
	{http.MethodGet, "/gopher/pencil/gopherrunning.jpg"},
	{http.MethodGet, "/gopher/pencil/gopherswim.jpg"},
	{http.MethodGet, "/gopher/pencil/gopherswrench.jpg"},
	{http.MethodGet, "/play"},
	{http.MethodGet, "/play/fib.go"},
	{http.MethodGet, "/play/hello.go"},
	{http.MethodGet, "/play/life.go"},
	{http.MethodGet, "/play/peano.go"},
	{http.MethodGet, "/play/pi.go"},
	{http.MethodGet, "/play/sieve.go"},
	{http.MethodGet, "/play/solitaire.go"},
	{http.MethodGet, "/play/tree.go"},
	{http.MethodGet, "/progs"},
	{http.MethodGet, "/progs/cgo1.go"},
	{http.MethodGet, "/progs/cgo2.go"},
	{http.MethodGet, "/progs/cgo3.go"},
	{http.MethodGet, "/progs/cgo4.go"},
	{http.MethodGet, "/progs/defer.go"},
	{http.MethodGet, "/progs/defer.out"},
	{http.MethodGet, "/progs/defer2.go"},
	{http.MethodGet, "/progs/defer2.out"},
	{http.MethodGet, "/progs/eff_bytesize.go"},
	{http.MethodGet, "/progs/eff_bytesize.out"},
	{http.MethodGet, "/progs/eff_qr.go"},
	{http.MethodGet, "/progs/eff_sequence.go"},
	{http.MethodGet, "/progs/eff_sequence.out"},
	{http.MethodGet, "/progs/eff_unused1.go"},
	{http.MethodGet, "/progs/eff_unused2.go"},
	{http.MethodGet, "/progs/error.go"},
	{http.MethodGet, "/progs/error2.go"},
	{http.MethodGet, "/progs/error3.go"},
	{http.MethodGet, "/progs/error4.go"},
	{http.MethodGet, "/progs/go1.go"},
	{http.MethodGet, "/progs/gobs1.go"},
	{http.MethodGet, "/progs/gobs2.go"},
	{http.MethodGet, "/progs/image_draw.go"},
	{http.MethodGet, "/progs/image_package1.go"},
	{http.MethodGet, "/progs/image_package1.out"},
	{http.MethodGet, "/progs/image_package2.go"},
	{http.MethodGet, "/progs/image_package2.out"},
	{http.MethodGet, "/progs/image_package3.go"},
	{http.MethodGet, "/progs/image_package3.out"},
	{http.MethodGet, "/progs/image_package4.go"},
	{http.MethodGet, "/progs/image_package4.out"},
	{http.MethodGet, "/progs/image_package5.go"},
	{http.MethodGet, "/progs/image_package5.out"},
	{http.MethodGet, "/progs/image_package6.go"},
	{http.MethodGet, "/progs/image_package6.out"},
	{http.MethodGet, "/progs/interface.go"},
	{http.MethodGet, "/progs/interface2.go"},
	{http.MethodGet, "/progs/interface2.out"},
	{http.MethodGet, "/progs/json1.go"},
	{http.MethodGet, "/progs/json2.go"},
	{http.MethodGet, "/progs/json2.out"},
	{http.MethodGet, "/progs/json3.go"},
	{http.MethodGet, "/progs/json4.go"},
	{http.MethodGet, "/progs/json5.go"},
	{http.MethodGet, "/progs/run"},
	{http.MethodGet, "/progs/slices.go"},
	{http.MethodGet, "/progs/timeout1.go"},
	{http.MethodGet, "/progs/timeout2.go"},
	{http.MethodGet, "/progs/update.bash"},
}

var (
	staticHttpServeMux http.Handler

	staticAce         http.Handler
	staticBear        http.Handler
	staticBeego       http.Handler
	staticBone        http.Handler
	staticChi         http.Handler
	staticDenco       http.Handler
	staticEcho        http.Handler
	staticGin         http.Handler
	staticGocraftWeb  http.Handler
	staticGoji        http.Handler
	staticGojiv2      http.Handler
	staticGoJsonRest  http.Handler
	staticGoRestful   http.Handler
	staticGorillaMux  http.Handler
	staticGowwwRouter http.Handler
	staticHttpRouter  http.Handler
	staticHttpTreeMux http.Handler
	staticKocha       http.Handler
	staticLARS        http.Handler
	staticMacaron     http.Handler
	staticMartini     http.Handler
	staticPat         http.Handler
	staticR2router    http.Handler
	staticRivet       http.Handler
	staticTigerTonic  http.Handler
	staticTraffic     http.Handler
	staticVulcan      http.Handler
)

func init() {
	println("#Static Routes:", len(staticRoutes))

	calcMem("HttpServeMux", func() {
		serveMux := http.NewServeMux()
		for _, route := range staticRoutes {
			serveMux.HandleFunc(route.path, httpHandlerFunc)
		}
		staticHttpServeMux = serveMux
	})

	calcMem("Ace", func() {
		staticAce = loadAce(staticRoutes)
	})
	calcMem("Bear", func() {
		staticBear = loadBear(staticRoutes)
	})
	calcMem("Beego", func() {
		staticBeego = loadBeego(staticRoutes)
	})
	calcMem("Bone", func() {
		staticBone = loadBone(staticRoutes)
	})
	calcMem("Chi", func() {
		staticChi = loadChi(staticRoutes)
	})
	calcMem("Denco", func() {
		staticDenco = loadDenco(staticRoutes)
	})
	calcMem("Echo", func() {
		staticEcho = loadEcho(staticRoutes)
	})
	calcMem("Gin", func() {
		staticGin = loadGin(staticRoutes)
	})
	calcMem("GocraftWeb", func() {
		staticGocraftWeb = loadGocraftWeb(staticRoutes)
	})
	calcMem("Goji", func() {
		staticGoji = loadGoji(staticRoutes)
	})
	calcMem("Gojiv2", func() {
		staticGojiv2 = loadGojiv2(staticRoutes)
	})
	calcMem("GoJsonRest", func() {
		staticGoJsonRest = loadGoJsonRest(staticRoutes)
	})
	calcMem("GoRestful", func() {
		staticGoRestful = loadGoRestful(staticRoutes)
	})
	calcMem("GorillaMux", func() {
		staticGorillaMux = loadGorillaMux(staticRoutes)
	})
	calcMem("GowwwRouter", func() {
		staticGowwwRouter = loadGowwwRouter(staticRoutes)
	})
	calcMem("HttpRouter", func() {
		staticHttpRouter = loadHttpRouter(staticRoutes)
	})
	calcMem("HttpTreeMux", func() {
		staticHttpTreeMux = loadHttpTreeMux(staticRoutes)
	})
	calcMem("Kocha", func() {
		staticKocha = loadKocha(staticRoutes)
	})
	calcMem("LARS", func() {
		staticLARS = loadLARS(staticRoutes)
	})
	calcMem("Macaron", func() {
		staticMacaron = loadMacaron(staticRoutes)
	})
	calcMem("Martini", func() {
		staticMartini = loadMartini(staticRoutes)
	})
	calcMem("Pat", func() {
		staticPat = loadPat(staticRoutes)
	})
	calcMem("R2router", func() {
		staticR2router = loadR2router(staticRoutes)
	})
	// calcMem("Revel", func() {
	// 	staticRevel = loadRevel(staticRoutes)
	// })
	calcMem("Rivet", func() {
		staticRivet = loadRivet(staticRoutes)
	})
	calcMem("TigerTonic", func() {
		staticTigerTonic = loadTigerTonic(staticRoutes)
	})
	calcMem("Traffic", func() {
		staticTraffic = loadTraffic(staticRoutes)
	})
	calcMem("Vulcan", func() {
		staticVulcan = loadVulcan(staticRoutes)
	})
	// calcMem("Zeus", func() {
	// 	staticZeus = loadZeus(staticRoutes)
	// })

	println()
}

// All routes

func BenchmarkAce_StaticAll(b *testing.B) {
	benchRoutes(b, staticAce, staticRoutes)
}
func BenchmarkHttpServeMux_StaticAll(b *testing.B) {
	benchRoutes(b, staticHttpServeMux, staticRoutes)
}
func BenchmarkBeego_StaticAll(b *testing.B) {
	benchRoutes(b, staticBeego, staticRoutes)
}
func BenchmarkBear_StaticAll(b *testing.B) {
	benchRoutes(b, staticBear, staticRoutes)
}
func BenchmarkBone_StaticAll(b *testing.B) {
	benchRoutes(b, staticBone, staticRoutes)
}
func BenchmarkChi_StaticAll(b *testing.B) {
	benchRoutes(b, staticChi, staticRoutes)
}
func BenchmarkDenco_StaticAll(b *testing.B) {
	benchRoutes(b, staticDenco, staticRoutes)
}
func BenchmarkEcho_StaticAll(b *testing.B) {
	benchRoutes(b, staticEcho, staticRoutes)
}
func BenchmarkGin_StaticAll(b *testing.B) {
	benchRoutes(b, staticGin, staticRoutes)
}
func BenchmarkGocraftWeb_StaticAll(b *testing.B) {
	benchRoutes(b, staticGocraftWeb, staticRoutes)
}
func BenchmarkGoji_StaticAll(b *testing.B) {
	benchRoutes(b, staticGoji, staticRoutes)
}
func BenchmarkGojiv2_StaticAll(b *testing.B) {
	benchRoutes(b, staticGojiv2, staticRoutes)
}
func BenchmarkGoJsonRest_StaticAll(b *testing.B) {
	benchRoutes(b, staticGoJsonRest, staticRoutes)
}
func BenchmarkGoRestful_StaticAll(b *testing.B) {
	benchRoutes(b, staticGoRestful, staticRoutes)
}
func BenchmarkGorillaMux_StaticAll(b *testing.B) {
	benchRoutes(b, staticGorillaMux, staticRoutes)
}
func BenchmarkGowwwRouter_StaticAll(b *testing.B) {
	benchRoutes(b, staticGowwwRouter, staticRoutes)
}
func BenchmarkHttpRouter_StaticAll(b *testing.B) {
	benchRoutes(b, staticHttpRouter, staticRoutes)
}
func BenchmarkHttpTreeMux_StaticAll(b *testing.B) {
	benchRoutes(b, staticHttpRouter, staticRoutes)
}
func BenchmarkKocha_StaticAll(b *testing.B) {
	benchRoutes(b, staticKocha, staticRoutes)
}
func BenchmarkLARS_StaticAll(b *testing.B) {
	benchRoutes(b, staticLARS, staticRoutes)
}
func BenchmarkMacaron_StaticAll(b *testing.B) {
	benchRoutes(b, staticMacaron, staticRoutes)
}
func BenchmarkMartini_StaticAll(b *testing.B) {
	benchRoutes(b, staticMartini, staticRoutes)
}
func BenchmarkPat_StaticAll(b *testing.B) {
	benchRoutes(b, staticPat, staticRoutes)
}
func BenchmarkR2router_StaticAll(b *testing.B) {
	benchRoutes(b, staticR2router, staticRoutes)
}

//	func BenchmarkRevel_StaticAll(b *testing.B) {
//		benchRoutes(b, staticRevel, staticRoutes)
//	}
func BenchmarkRivet_StaticAll(b *testing.B) {
	benchRoutes(b, staticRivet, staticRoutes)
}

func BenchmarkTigerTonic_StaticAll(b *testing.B) {
	benchRoutes(b, staticTigerTonic, staticRoutes)
}
func BenchmarkTraffic_StaticAll(b *testing.B) {
	benchRoutes(b, staticTraffic, staticRoutes)
}
func BenchmarkVulcan_StaticAll(b *testing.B) {
	benchRoutes(b, staticVulcan, staticRoutes)
}

// func BenchmarkZeus_StaticAll(b *testing.B) {
// 	benchRoutes(b, staticZeus, staticRoutes)
// }
