// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strings"
	"testing"
)

var benchRe *regexp.Regexp

func isTested(name string) bool {
	if benchRe == nil {
		// Get -test.bench flag value (not accessible via flag package)
		bench := ""
		for _, arg := range os.Args {
			if strings.HasPrefix(arg, "-test.bench=") {
				// ignore the benchmark name after an underscore
				bench = strings.SplitN(arg[12:], "_", 2)[0]
				break
			}
		}

		// Compile RegExp to match Benchmark names
		var err error
		benchRe, err = regexp.Compile(bench)
		if err != nil {
			panic(err.Error())
		}
	}
	return benchRe.MatchString(name)
}

func calcMem(name string, load func()) {
	if !isTested(name) {
		return
	}

	m := new(runtime.MemStats)

	// before
	// force GC multiple times, since Go is using a generational GC
	// TODO: find a better approach
	runtime.GC()
	runtime.GC()
	runtime.GC()
	runtime.GC()
	runtime.ReadMemStats(m)
	before := m.HeapAlloc

	load()

	// after
	runtime.GC()
	runtime.GC()
	runtime.GC()
	runtime.GC()
	runtime.ReadMemStats(m)
	after := m.HeapAlloc
	println("   "+name+":", after-before, "Bytes")
}

func benchRequest(b *testing.B, router http.Handler, r *http.Request) {
	w := new(mockResponseWriter)
	u := r.URL
	rq := u.RawQuery
	r.RequestURI = u.RequestURI()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		u.RawQuery = rq
		router.ServeHTTP(w, r)
	}
}

func benchRoutes(b *testing.B, router http.Handler, routes []route) {
	w := new(mockResponseWriter)
	r, _ := http.NewRequest(http.MethodGet, "/", nil)
	u := r.URL
	rq := u.RawQuery

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, route := range routes {
			r.Method = route.method
			r.RequestURI = route.path
			u.Path = route.path
			u.RawQuery = rq
			router.ServeHTTP(w, r)
		}
	}
}

// Micro Benchmarks

// Route with Param (no write)
func BenchmarkAce_Param(b *testing.B) {
	router := loadAceSingle(http.MethodGet, "/user/:name", aceHandle)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}

func BenchmarkBear_Param(b *testing.B) {
	router := loadBearSingle(http.MethodGet, "/user/{name}", bearHandler)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkBeego_Param(b *testing.B) {
	router := loadBeegoSingle(http.MethodGet, "/user/:name", beegoHandler)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkBone_Param(b *testing.B) {
	router := loadBoneSingle(http.MethodGet, "/user/:name", http.HandlerFunc(httpHandlerFunc))

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkChi_Param(b *testing.B) {
	router := loadChiSingle(http.MethodGet, "/user/{name}", httpHandlerFunc)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}

func BenchmarkSuperhttp_Param(b *testing.B) {
	router := loadSuperhttpSingle(http.MethodGet, "/user/{name}", httpHandlerFunc)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}

func BenchmarkDenco_Param(b *testing.B) {
	router := loadDencoSingle(http.MethodGet, "/user/:name", dencoHandler)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkEcho_Param(b *testing.B) {
	router := loadEchoSingle(http.MethodGet, "/user/:name", echoHandler)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_Param(b *testing.B) {
	router := loadGinSingle(http.MethodGet, "/user/:name", ginHandle)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGocraftWeb_Param(b *testing.B) {
	router := loadGocraftWebSingle(http.MethodGet, "/user/:name", gocraftWebHandler)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGoji_Param(b *testing.B) {
	router := loadGojiSingle(http.MethodGet, "/user/:name", httpHandlerFunc)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGojiv2_Param(b *testing.B) {
	router := loadGojiv2Single(http.MethodGet, "/user/:name", gojiv2Handler)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGoJsonRest_Param(b *testing.B) {
	router := loadGoJsonRestSingle(http.MethodGet, "/user/:name", goJsonRestHandler)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGoRestful_Param(b *testing.B) {
	router := loadGoRestfulSingle(http.MethodGet, "/user/{name}", goRestfulHandler)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGorillaMux_Param(b *testing.B) {
	router := loadGorillaMuxSingle(http.MethodGet, "/user/{name}", httpHandlerFunc)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGowwwRouter_Param(b *testing.B) {
	router := loadGowwwRouterSingle(http.MethodGet, "/user/:name", http.HandlerFunc(httpHandlerFunc))

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_Param(b *testing.B) {
	router := loadHttpRouterSingle(http.MethodGet, "/user/:name", httpRouterHandle)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpTreeMux_Param(b *testing.B) {
	router := loadHttpTreeMuxSingle(http.MethodGet, "/user/:name", httpTreeMuxHandler)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkKocha_Param(b *testing.B) {
	handler := new(kochaHandler)
	router := loadKochaSingle(
		http.MethodGet, "/user/:name",
		handler, http.HandlerFunc(handler.Get),
	)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkLARS_Param(b *testing.B) {
	router := loadLARSSingle(http.MethodGet, "/user/:name", larsHandler)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkMacaron_Param(b *testing.B) {
	router := loadMacaronSingle(http.MethodGet, "/user/:name", macaronHandler)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkMartini_Param(b *testing.B) {
	router := loadMartiniSingle(http.MethodGet, "/user/:name", martiniHandler)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkPat_Param(b *testing.B) {
	router := loadPatSingle(http.MethodGet, "/user/:name", http.HandlerFunc(httpHandlerFunc))

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}

func BenchmarkR2router_Param(b *testing.B) {
	router := loadR2routerSingle(http.MethodGet, "/user/:name", r2routerHandler)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}

// func BenchmarkRevel_Param(b *testing.B) {
// 	router := loadRevelSingle(http.MethodGet, "/user/:name", "RevelController.Handle")

//		r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
//		benchRequest(b, router, r)
//	}
func BenchmarkRivet_Param(b *testing.B) {
	router := loadRivetSingle(http.MethodGet, "/user/:name", rivetHandler)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}

func BenchmarkTigerTonic_Param(b *testing.B) {
	router := loadTigerTonicSingle(http.MethodGet, "/user/{name}", httpHandlerFunc)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkTraffic_Param(b *testing.B) {
	router := loadTrafficSingle(http.MethodGet, "/user/:name", trafficHandler)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkVulcan_Param(b *testing.B) {
	router := loadVulcanSingle(http.MethodGet, "/user/:name", vulcanHandler)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}

// func BenchmarkZeus_Param(b *testing.B) {
// 	router := loadZeusSingle(http.MethodGet, "/user/:name", http.HandlerFunc(httpHandlerFunc))

// 	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
// 	benchRequest(b, router, r)
// }

// Route with 5 Params (no write)
const fiveColon = "/:a/:b/:c/:d/:e"
const fiveBrace = "/{a}/{b}/{c}/{d}/{e}"
const fiveRoute = "/test/test/test/test/test"

func BenchmarkAce_Param5(b *testing.B) {
	router := loadAceSingle(http.MethodGet, fiveColon, aceHandle)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}

func BenchmarkBear_Param5(b *testing.B) {
	router := loadBearSingle(http.MethodGet, fiveBrace, bearHandler)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkBeego_Param5(b *testing.B) {
	router := loadBeegoSingle(http.MethodGet, fiveColon, beegoHandler)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkBone_Param5(b *testing.B) {
	router := loadBoneSingle(http.MethodGet, fiveColon, http.HandlerFunc(httpHandlerFunc))

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkChi_Param5(b *testing.B) {
	router := loadChiSingle(http.MethodGet, fiveBrace, httpHandlerFunc)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}

func BenchmarkSuperhttp_Param5(b *testing.B) {
	router := loadSuperhttpSingle(http.MethodGet, fiveBrace, httpHandlerFunc)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}

func BenchmarkDenco_Param5(b *testing.B) {
	router := loadDencoSingle(http.MethodGet, fiveColon, dencoHandler)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkEcho_Param5(b *testing.B) {
	router := loadEchoSingle(http.MethodGet, fiveColon, echoHandler)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_Param5(b *testing.B) {
	router := loadGinSingle(http.MethodGet, fiveColon, ginHandle)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGocraftWeb_Param5(b *testing.B) {
	router := loadGocraftWebSingle(http.MethodGet, fiveColon, gocraftWebHandler)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGoji_Param5(b *testing.B) {
	router := loadGojiSingle(http.MethodGet, fiveColon, httpHandlerFunc)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGojiv2_Param5(b *testing.B) {
	router := loadGojiv2Single(http.MethodGet, fiveColon, gojiv2Handler)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGoJsonRest_Param5(b *testing.B) {
	handler := loadGoJsonRestSingle(http.MethodGet, fiveColon, goJsonRestHandler)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, handler, r)
}
func BenchmarkGoRestful_Param5(b *testing.B) {
	router := loadGoRestfulSingle(http.MethodGet, fiveBrace, goRestfulHandler)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGorillaMux_Param5(b *testing.B) {
	router := loadGorillaMuxSingle(http.MethodGet, fiveBrace, httpHandlerFunc)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGowwwRouter_Param5(b *testing.B) {
	router := loadGowwwRouterSingle(http.MethodGet, fiveColon, http.HandlerFunc(httpHandlerFunc))

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_Param5(b *testing.B) {
	router := loadHttpRouterSingle(http.MethodGet, fiveColon, httpRouterHandle)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpTreeMux_Param5(b *testing.B) {
	router := loadHttpTreeMuxSingle(http.MethodGet, fiveColon, httpTreeMuxHandler)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkKocha_Param5(b *testing.B) {
	handler := new(kochaHandler)
	router := loadKochaSingle(
		http.MethodGet, fiveColon,
		handler, http.HandlerFunc(handler.Get),
	)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkLARS_Param5(b *testing.B) {
	router := loadLARSSingle(http.MethodGet, fiveColon, larsHandler)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkMacaron_Param5(b *testing.B) {
	router := loadMacaronSingle(http.MethodGet, fiveColon, macaronHandler)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkMartini_Param5(b *testing.B) {
	router := loadMartiniSingle(http.MethodGet, fiveColon, martiniHandler)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkPat_Param5(b *testing.B) {
	router := loadPatSingle(http.MethodGet, fiveColon, http.HandlerFunc(httpHandlerFunc))

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}

func BenchmarkR2router_Param5(b *testing.B) {
	router := loadR2routerSingle(http.MethodGet, fiveColon, r2routerHandler)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}

// func BenchmarkRevel_Param5(b *testing.B) {
// 	router := loadRevelSingle(http.MethodGet, fiveColon, "RevelController.Handle")

//		r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
//		benchRequest(b, router, r)
//	}
func BenchmarkRivet_Param5(b *testing.B) {
	router := loadRivetSingle(http.MethodGet, fiveColon, rivetHandler)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}

func BenchmarkTigerTonic_Param5(b *testing.B) {
	router := loadTigerTonicSingle(http.MethodGet, fiveBrace, httpHandlerFunc)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkTraffic_Param5(b *testing.B) {
	router := loadTrafficSingle(http.MethodGet, fiveColon, trafficHandler)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkVulcan_Param5(b *testing.B) {
	router := loadVulcanSingle(http.MethodGet, fiveColon, vulcanHandler)

	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
	benchRequest(b, router, r)
}

// func BenchmarkZeus_Param5(b *testing.B) {
// 	router := loadZeusSingle(http.MethodGet, fiveColon, http.HandlerFunc(httpHandlerFunc))

// 	r, _ := http.NewRequest(http.MethodGet, fiveRoute, nil)
// 	benchRequest(b, router, r)
// }

// Route with 20 Params (no write)
const twentyColon = "/:a/:b/:c/:d/:e/:f/:g/:h/:i/:j/:k/:l/:m/:n/:o/:p/:q/:r/:s/:t"
const twentyBrace = "/{a}/{b}/{c}/{d}/{e}/{f}/{g}/{h}/{i}/{j}/{k}/{l}/{m}/{n}/{o}/{p}/{q}/{r}/{s}/{t}"
const twentyRoute = "/a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t"

func BenchmarkAce_Param20(b *testing.B) {
	router := loadAceSingle(http.MethodGet, twentyColon, aceHandle)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}

func BenchmarkBear_Param20(b *testing.B) {
	router := loadBearSingle(http.MethodGet, twentyBrace, bearHandler)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkBeego_Param20(b *testing.B) {
	router := loadBeegoSingle(http.MethodGet, twentyColon, beegoHandler)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkBone_Param20(b *testing.B) {
	router := loadBoneSingle(http.MethodGet, twentyColon, http.HandlerFunc(httpHandlerFunc))

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkChi_Param20(b *testing.B) {
	router := loadChiSingle(http.MethodGet, twentyBrace, httpHandlerFunc)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}

func BenchmarkSuperhttp_Param20(b *testing.B) {
	router := loadSuperhttpSingle(http.MethodGet, twentyBrace, httpHandlerFunc)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}

func BenchmarkDenco_Param20(b *testing.B) {
	router := loadDencoSingle(http.MethodGet, twentyColon, dencoHandler)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkEcho_Param20(b *testing.B) {
	router := loadEchoSingle(http.MethodGet, twentyColon, echoHandler)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_Param20(b *testing.B) {
	router := loadGinSingle(http.MethodGet, twentyColon, ginHandle)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGocraftWeb_Param20(b *testing.B) {
	router := loadGocraftWebSingle(http.MethodGet, twentyColon, gocraftWebHandler)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGoji_Param20(b *testing.B) {
	router := loadGojiSingle(http.MethodGet, twentyColon, httpHandlerFunc)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGojiv2_Param20(b *testing.B) {
	router := loadGojiv2Single(http.MethodGet, twentyColon, gojiv2Handler)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGoJsonRest_Param20(b *testing.B) {
	handler := loadGoJsonRestSingle(http.MethodGet, twentyColon, goJsonRestHandler)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, handler, r)
}
func BenchmarkGoRestful_Param20(b *testing.B) {
	handler := loadGoRestfulSingle(http.MethodGet, twentyBrace, goRestfulHandler)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, handler, r)
}
func BenchmarkGorillaMux_Param20(b *testing.B) {
	router := loadGorillaMuxSingle(http.MethodGet, twentyBrace, httpHandlerFunc)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkGowwwRouter_Param20(b *testing.B) {
	router := loadGowwwRouterSingle(http.MethodGet, twentyColon, http.HandlerFunc(httpHandlerFunc))

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_Param20(b *testing.B) {
	router := loadHttpRouterSingle(http.MethodGet, twentyColon, httpRouterHandle)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpTreeMux_Param20(b *testing.B) {
	router := loadHttpTreeMuxSingle(http.MethodGet, twentyColon, httpTreeMuxHandler)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkKocha_Param20(b *testing.B) {
	handler := new(kochaHandler)
	router := loadKochaSingle(
		http.MethodGet, twentyColon,
		handler, http.HandlerFunc(handler.Get),
	)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkLARS_Param20(b *testing.B) {
	router := loadLARSSingle(http.MethodGet, twentyColon, larsHandler)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkMacaron_Param20(b *testing.B) {
	router := loadMacaronSingle(http.MethodGet, twentyColon, macaronHandler)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkMartini_Param20(b *testing.B) {
	router := loadMartiniSingle(http.MethodGet, twentyColon, martiniHandler)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkPat_Param20(b *testing.B) {
	router := loadPatSingle(http.MethodGet, twentyColon, http.HandlerFunc(httpHandlerFunc))

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}

func BenchmarkR2router_Param20(b *testing.B) {
	router := loadR2routerSingle(http.MethodGet, twentyColon, r2routerHandler)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}

// func BenchmarkRevel_Param20(b *testing.B) {
// 	router := loadRevelSingle(http.MethodGet, twentyColon, "RevelController.Handle")

//		r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
//		benchRequest(b, router, r)
//	}
func BenchmarkRivet_Param20(b *testing.B) {
	router := loadRivetSingle(http.MethodGet, twentyColon, rivetHandler)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}

func BenchmarkTigerTonic_Param20(b *testing.B) {
	router := loadTigerTonicSingle(http.MethodGet, twentyBrace, httpHandlerFunc)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkTraffic_Param20(b *testing.B) {
	router := loadTrafficSingle(http.MethodGet, twentyColon, trafficHandler)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}
func BenchmarkVulcan_Param20(b *testing.B) {
	router := loadVulcanSingle(http.MethodGet, twentyColon, vulcanHandler)

	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
	benchRequest(b, router, r)
}

// func BenchmarkZeus_Param20(b *testing.B) {
// 	router := loadZeusSingle(http.MethodGet, twentyColon, http.HandlerFunc(httpHandlerFunc))

// 	r, _ := http.NewRequest(http.MethodGet, twentyRoute, nil)
// 	benchRequest(b, router, r)
// }

// Route with Param and write
func BenchmarkAce_ParamWrite(b *testing.B) {
	router := loadAceSingle(http.MethodGet, "/user/:name", aceHandleWrite)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}

func BenchmarkBear_ParamWrite(b *testing.B) {
	router := loadBearSingle(http.MethodGet, "/user/{name}", bearHandlerWrite)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkBeego_ParamWrite(b *testing.B) {
	router := loadBeegoSingle(http.MethodGet, "/user/:name", beegoHandlerWrite)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkBone_ParamWrite(b *testing.B) {
	router := loadBoneSingle(http.MethodGet, "/user/:name", http.HandlerFunc(boneHandlerWrite))

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkChi_ParamWrite(b *testing.B) {
	router := loadChiSingle(http.MethodGet, "/user/{name}", chiHandleWrite)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}

func BenchmarkSuperhttp_ParamWrite(b *testing.B) {
	router := loadSuperhttpSingle(http.MethodGet, "/user/{name}", superhttpHandleWrite)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}

func BenchmarkDenco_ParamWrite(b *testing.B) {
	router := loadDencoSingle(http.MethodGet, "/user/:name", dencoHandlerWrite)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkEcho_ParamWrite(b *testing.B) {
	router := loadEchoSingle(http.MethodGet, "/user/:name", echoHandlerWrite)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_ParamWrite(b *testing.B) {
	router := loadGinSingle(http.MethodGet, "/user/:name", ginHandleWrite)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGocraftWeb_ParamWrite(b *testing.B) {
	router := loadGocraftWebSingle(http.MethodGet, "/user/:name", gocraftWebHandlerWrite)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGoji_ParamWrite(b *testing.B) {
	router := loadGojiSingle(http.MethodGet, "/user/:name", gojiFuncWrite)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGojiv2_ParamWrite(b *testing.B) {
	router := loadGojiv2Single(http.MethodGet, "/user/:name", gojiv2HandlerWrite)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGoJsonRest_ParamWrite(b *testing.B) {
	handler := loadGoJsonRestSingle(http.MethodGet, "/user/:name", goJsonRestHandlerWrite)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, handler, r)
}
func BenchmarkGoRestful_ParamWrite(b *testing.B) {
	handler := loadGoRestfulSingle(http.MethodGet, "/user/{name}", goRestfulHandlerWrite)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, handler, r)
}
func BenchmarkGorillaMux_ParamWrite(b *testing.B) {
	router := loadGorillaMuxSingle(http.MethodGet, "/user/{name}", gorillaHandlerWrite)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkGowwwRouter_ParamWrite(b *testing.B) {
	router := loadGowwwRouterSingle(http.MethodGet, "/user/:name", http.HandlerFunc(gowwwRouterHandleWrite))

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpRouter_ParamWrite(b *testing.B) {
	router := loadHttpRouterSingle(http.MethodGet, "/user/:name", httpRouterHandleWrite)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkHttpTreeMux_ParamWrite(b *testing.B) {
	router := loadHttpTreeMuxSingle(http.MethodGet, "/user/:name", httpTreeMuxHandlerWrite)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkKocha_ParamWrite(b *testing.B) {
	handler := new(kochaHandler)
	router := loadKochaSingle(
		http.MethodGet, "/user/:name",
		handler, http.HandlerFunc(handler.kochaHandlerWrite),
	)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkLARS_ParamWrite(b *testing.B) {
	router := loadLARSSingle(http.MethodGet, "/user/:name", larsHandlerWrite)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkMacaron_ParamWrite(b *testing.B) {
	router := loadMacaronSingle(http.MethodGet, "/user/:name", macaronHandlerWrite)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkMartini_ParamWrite(b *testing.B) {
	router := loadMartiniSingle(http.MethodGet, "/user/:name", martiniHandlerWrite)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkPat_ParamWrite(b *testing.B) {
	router := loadPatSingle(http.MethodGet, "/user/:name", http.HandlerFunc(patHandlerWrite))

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}

func BenchmarkR2router_ParamWrite(b *testing.B) {
	router := loadR2routerSingle(http.MethodGet, "/user/:name", r2routerHandleWrite)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}

// func BenchmarkRevel_ParamWrite(b *testing.B) {
// 	router := loadRevelSingle(http.MethodGet, "/user/:name", "RevelController.HandleWrite")

//		r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
//		benchRequest(b, router, r)
//	}
func BenchmarkRivet_ParamWrite(b *testing.B) {
	router := loadRivetSingle(http.MethodGet, "/user/:name", rivetHandlerWrite)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkTigerTonic_ParamWrite(b *testing.B) {
	router := loadTigerTonicSingle(
		http.MethodGet, "/user/{name}",
		http.HandlerFunc(tigerTonicHandlerWrite),
	)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkTraffic_ParamWrite(b *testing.B) {
	router := loadTrafficSingle(http.MethodGet, "/user/:name", trafficHandlerWrite)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}
func BenchmarkVulcan_ParamWrite(b *testing.B) {
	router := loadVulcanSingle(http.MethodGet, "/user/:name", vulcanHandlerWrite)

	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
	benchRequest(b, router, r)
}

// func BenchmarkZeus_ParamWrite(b *testing.B) {
// 	router := loadZeusSingle(http.MethodGet, "/user/:name", zeusHandlerWrite)

// 	r, _ := http.NewRequest(http.MethodGet, "/user/gordon", nil)
// 	benchRequest(b, router, r)
// }
