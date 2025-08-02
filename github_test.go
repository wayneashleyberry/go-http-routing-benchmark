// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"net/http"
	"testing"
)

// http://developer.github.com/v3/
var githubAPI = []route{
	// OAuth Authorizations
	{http.MethodGet, "/authorizations"},
	{http.MethodGet, "/authorizations/:id"},
	{http.MethodPost, "/authorizations"},
	//{http.MethodPut, "/authorizations/clients/:client_id"},
	//{http.MethodPatch, "/authorizations/:id"},
	{http.MethodDelete, "/authorizations/:id"},
	{http.MethodGet, "/applications/:client_id/tokens/:access_token"},
	{http.MethodDelete, "/applications/:client_id/tokens"},
	{http.MethodDelete, "/applications/:client_id/tokens/:access_token"},

	// Activity
	{http.MethodGet, "/events"},
	{http.MethodGet, "/repos/:owner/:repo/events"},
	{http.MethodGet, "/networks/:owner/:repo/events"},
	{http.MethodGet, "/orgs/:org/events"},
	{http.MethodGet, "/users/:user/received_events"},
	{http.MethodGet, "/users/:user/received_events/public"},
	{http.MethodGet, "/users/:user/events"},
	{http.MethodGet, "/users/:user/events/public"},
	{http.MethodGet, "/users/:user/events/orgs/:org"},
	{http.MethodGet, "/feeds"},
	{http.MethodGet, "/notifications"},
	{http.MethodGet, "/repos/:owner/:repo/notifications"},
	{http.MethodPut, "/notifications"},
	{http.MethodPut, "/repos/:owner/:repo/notifications"},
	{http.MethodGet, "/notifications/threads/:id"},
	//{http.MethodPatch, "/notifications/threads/:id"},
	{http.MethodGet, "/notifications/threads/:id/subscription"},
	{http.MethodPut, "/notifications/threads/:id/subscription"},
	{http.MethodDelete, "/notifications/threads/:id/subscription"},
	{http.MethodGet, "/repos/:owner/:repo/stargazers"},
	{http.MethodGet, "/users/:user/starred"},
	{http.MethodGet, "/user/starred"},
	{http.MethodGet, "/user/starred/:owner/:repo"},
	{http.MethodPut, "/user/starred/:owner/:repo"},
	{http.MethodDelete, "/user/starred/:owner/:repo"},
	{http.MethodGet, "/repos/:owner/:repo/subscribers"},
	{http.MethodGet, "/users/:user/subscriptions"},
	{http.MethodGet, "/user/subscriptions"},
	{http.MethodGet, "/repos/:owner/:repo/subscription"},
	{http.MethodPut, "/repos/:owner/:repo/subscription"},
	{http.MethodDelete, "/repos/:owner/:repo/subscription"},
	{http.MethodGet, "/user/subscriptions/:owner/:repo"},
	{http.MethodPut, "/user/subscriptions/:owner/:repo"},
	{http.MethodDelete, "/user/subscriptions/:owner/:repo"},

	// Gists
	{http.MethodGet, "/users/:user/gists"},
	{http.MethodGet, "/gists"},
	//{http.MethodGet, "/gists/public"},
	//{http.MethodGet, "/gists/starred"},
	{http.MethodGet, "/gists/:id"},
	{http.MethodPost, "/gists"},
	//{http.MethodPatch, "/gists/:id"},
	{http.MethodPut, "/gists/:id/star"},
	{http.MethodDelete, "/gists/:id/star"},
	{http.MethodGet, "/gists/:id/star"},
	{http.MethodPost, "/gists/:id/forks"},
	{http.MethodDelete, "/gists/:id"},

	// Git Data
	{http.MethodGet, "/repos/:owner/:repo/git/blobs/:sha"},
	{http.MethodPost, "/repos/:owner/:repo/git/blobs"},
	{http.MethodGet, "/repos/:owner/:repo/git/commits/:sha"},
	{http.MethodPost, "/repos/:owner/:repo/git/commits"},
	//{http.MethodGet, "/repos/:owner/:repo/git/refs/*ref"},
	{http.MethodGet, "/repos/:owner/:repo/git/refs"},
	{http.MethodPost, "/repos/:owner/:repo/git/refs"},
	//{http.MethodPatch, "/repos/:owner/:repo/git/refs/*ref"},
	//{http.MethodDelete, "/repos/:owner/:repo/git/refs/*ref"},
	{http.MethodGet, "/repos/:owner/:repo/git/tags/:sha"},
	{http.MethodPost, "/repos/:owner/:repo/git/tags"},
	{http.MethodGet, "/repos/:owner/:repo/git/trees/:sha"},
	{http.MethodPost, "/repos/:owner/:repo/git/trees"},

	// Issues
	{http.MethodGet, "/issues"},
	{http.MethodGet, "/user/issues"},
	{http.MethodGet, "/orgs/:org/issues"},
	{http.MethodGet, "/repos/:owner/:repo/issues"},
	{http.MethodGet, "/repos/:owner/:repo/issues/:number"},
	{http.MethodPost, "/repos/:owner/:repo/issues"},
	//{http.MethodPatch, "/repos/:owner/:repo/issues/:number"},
	{http.MethodGet, "/repos/:owner/:repo/assignees"},
	{http.MethodGet, "/repos/:owner/:repo/assignees/:assignee"},
	{http.MethodGet, "/repos/:owner/:repo/issues/:number/comments"},
	//{http.MethodGet, "/repos/:owner/:repo/issues/comments"},
	//{http.MethodGet, "/repos/:owner/:repo/issues/comments/:id"},
	{http.MethodPost, "/repos/:owner/:repo/issues/:number/comments"},
	//{http.MethodPatch, "/repos/:owner/:repo/issues/comments/:id"},
	//{http.MethodDelete, "/repos/:owner/:repo/issues/comments/:id"},
	{http.MethodGet, "/repos/:owner/:repo/issues/:number/events"},
	//{http.MethodGet, "/repos/:owner/:repo/issues/events"},
	//{http.MethodGet, "/repos/:owner/:repo/issues/events/:id"},
	{http.MethodGet, "/repos/:owner/:repo/labels"},
	{http.MethodGet, "/repos/:owner/:repo/labels/:name"},
	{http.MethodPost, "/repos/:owner/:repo/labels"},
	//{http.MethodPatch, "/repos/:owner/:repo/labels/:name"},
	{http.MethodDelete, "/repos/:owner/:repo/labels/:name"},
	{http.MethodGet, "/repos/:owner/:repo/issues/:number/labels"},
	{http.MethodPost, "/repos/:owner/:repo/issues/:number/labels"},
	{http.MethodDelete, "/repos/:owner/:repo/issues/:number/labels/:name"},
	{http.MethodPut, "/repos/:owner/:repo/issues/:number/labels"},
	{http.MethodDelete, "/repos/:owner/:repo/issues/:number/labels"},
	{http.MethodGet, "/repos/:owner/:repo/milestones/:number/labels"},
	{http.MethodGet, "/repos/:owner/:repo/milestones"},
	{http.MethodGet, "/repos/:owner/:repo/milestones/:number"},
	{http.MethodPost, "/repos/:owner/:repo/milestones"},
	//{http.MethodPatch, "/repos/:owner/:repo/milestones/:number"},
	{http.MethodDelete, "/repos/:owner/:repo/milestones/:number"},

	// Miscellaneous
	{http.MethodGet, "/emojis"},
	{http.MethodGet, "/gitignore/templates"},
	{http.MethodGet, "/gitignore/templates/:name"},
	{http.MethodPost, "/markdown"},
	{http.MethodPost, "/markdown/raw"},
	{http.MethodGet, "/meta"},
	{http.MethodGet, "/rate_limit"},

	// Organizations
	{http.MethodGet, "/users/:user/orgs"},
	{http.MethodGet, "/user/orgs"},
	{http.MethodGet, "/orgs/:org"},
	//{http.MethodPatch, "/orgs/:org"},
	{http.MethodGet, "/orgs/:org/members"},
	{http.MethodGet, "/orgs/:org/members/:user"},
	{http.MethodDelete, "/orgs/:org/members/:user"},
	{http.MethodGet, "/orgs/:org/public_members"},
	{http.MethodGet, "/orgs/:org/public_members/:user"},
	{http.MethodPut, "/orgs/:org/public_members/:user"},
	{http.MethodDelete, "/orgs/:org/public_members/:user"},
	{http.MethodGet, "/orgs/:org/teams"},
	{http.MethodGet, "/teams/:id"},
	{http.MethodPost, "/orgs/:org/teams"},
	//{http.MethodPatch, "/teams/:id"},
	{http.MethodDelete, "/teams/:id"},
	{http.MethodGet, "/teams/:id/members"},
	{http.MethodGet, "/teams/:id/members/:user"},
	{http.MethodPut, "/teams/:id/members/:user"},
	{http.MethodDelete, "/teams/:id/members/:user"},
	{http.MethodGet, "/teams/:id/repos"},
	{http.MethodGet, "/teams/:id/repos/:owner/:repo"},
	{http.MethodPut, "/teams/:id/repos/:owner/:repo"},
	{http.MethodDelete, "/teams/:id/repos/:owner/:repo"},
	{http.MethodGet, "/user/teams"},

	// Pull Requests
	{http.MethodGet, "/repos/:owner/:repo/pulls"},
	{http.MethodGet, "/repos/:owner/:repo/pulls/:number"},
	{http.MethodPost, "/repos/:owner/:repo/pulls"},
	//{http.MethodPatch, "/repos/:owner/:repo/pulls/:number"},
	{http.MethodGet, "/repos/:owner/:repo/pulls/:number/commits"},
	{http.MethodGet, "/repos/:owner/:repo/pulls/:number/files"},
	{http.MethodGet, "/repos/:owner/:repo/pulls/:number/merge"},
	{http.MethodPut, "/repos/:owner/:repo/pulls/:number/merge"},
	{http.MethodGet, "/repos/:owner/:repo/pulls/:number/comments"},
	//{http.MethodGet, "/repos/:owner/:repo/pulls/comments"},
	//{http.MethodGet, "/repos/:owner/:repo/pulls/comments/:number"},
	{http.MethodPut, "/repos/:owner/:repo/pulls/:number/comments"},
	//{http.MethodPatch, "/repos/:owner/:repo/pulls/comments/:number"},
	//{http.MethodDelete, "/repos/:owner/:repo/pulls/comments/:number"},

	// Repositories
	{http.MethodGet, "/user/repos"},
	{http.MethodGet, "/users/:user/repos"},
	{http.MethodGet, "/orgs/:org/repos"},
	{http.MethodGet, "/repositories"},
	{http.MethodPost, "/user/repos"},
	{http.MethodPost, "/orgs/:org/repos"},
	{http.MethodGet, "/repos/:owner/:repo"},
	//{http.MethodPatch, "/repos/:owner/:repo"},
	{http.MethodGet, "/repos/:owner/:repo/contributors"},
	{http.MethodGet, "/repos/:owner/:repo/languages"},
	{http.MethodGet, "/repos/:owner/:repo/teams"},
	{http.MethodGet, "/repos/:owner/:repo/tags"},
	{http.MethodGet, "/repos/:owner/:repo/branches"},
	{http.MethodGet, "/repos/:owner/:repo/branches/:branch"},
	{http.MethodDelete, "/repos/:owner/:repo"},
	{http.MethodGet, "/repos/:owner/:repo/collaborators"},
	{http.MethodGet, "/repos/:owner/:repo/collaborators/:user"},
	{http.MethodPut, "/repos/:owner/:repo/collaborators/:user"},
	{http.MethodDelete, "/repos/:owner/:repo/collaborators/:user"},
	{http.MethodGet, "/repos/:owner/:repo/comments"},
	{http.MethodGet, "/repos/:owner/:repo/commits/:sha/comments"},
	{http.MethodPost, "/repos/:owner/:repo/commits/:sha/comments"},
	{http.MethodGet, "/repos/:owner/:repo/comments/:id"},
	//{http.MethodPatch, "/repos/:owner/:repo/comments/:id"},
	{http.MethodDelete, "/repos/:owner/:repo/comments/:id"},
	{http.MethodGet, "/repos/:owner/:repo/commits"},
	{http.MethodGet, "/repos/:owner/:repo/commits/:sha"},
	{http.MethodGet, "/repos/:owner/:repo/readme"},
	//{http.MethodGet, "/repos/:owner/:repo/contents/*path"},
	//{http.MethodPut, "/repos/:owner/:repo/contents/*path"},
	//{http.MethodDelete, "/repos/:owner/:repo/contents/*path"},
	//{http.MethodGet, "/repos/:owner/:repo/:archive_format/:ref"},
	{http.MethodGet, "/repos/:owner/:repo/keys"},
	{http.MethodGet, "/repos/:owner/:repo/keys/:id"},
	{http.MethodPost, "/repos/:owner/:repo/keys"},
	//{http.MethodPatch, "/repos/:owner/:repo/keys/:id"},
	{http.MethodDelete, "/repos/:owner/:repo/keys/:id"},
	{http.MethodGet, "/repos/:owner/:repo/downloads"},
	{http.MethodGet, "/repos/:owner/:repo/downloads/:id"},
	{http.MethodDelete, "/repos/:owner/:repo/downloads/:id"},
	{http.MethodGet, "/repos/:owner/:repo/forks"},
	{http.MethodPost, "/repos/:owner/:repo/forks"},
	{http.MethodGet, "/repos/:owner/:repo/hooks"},
	{http.MethodGet, "/repos/:owner/:repo/hooks/:id"},
	{http.MethodPost, "/repos/:owner/:repo/hooks"},
	//{http.MethodPatch, "/repos/:owner/:repo/hooks/:id"},
	{http.MethodPost, "/repos/:owner/:repo/hooks/:id/tests"},
	{http.MethodDelete, "/repos/:owner/:repo/hooks/:id"},
	{http.MethodPost, "/repos/:owner/:repo/merges"},
	{http.MethodGet, "/repos/:owner/:repo/releases"},
	{http.MethodGet, "/repos/:owner/:repo/releases/:id"},
	{http.MethodPost, "/repos/:owner/:repo/releases"},
	//{http.MethodPatch, "/repos/:owner/:repo/releases/:id"},
	{http.MethodDelete, "/repos/:owner/:repo/releases/:id"},
	{http.MethodGet, "/repos/:owner/:repo/releases/:id/assets"},
	{http.MethodGet, "/repos/:owner/:repo/stats/contributors"},
	{http.MethodGet, "/repos/:owner/:repo/stats/commit_activity"},
	{http.MethodGet, "/repos/:owner/:repo/stats/code_frequency"},
	{http.MethodGet, "/repos/:owner/:repo/stats/participation"},
	{http.MethodGet, "/repos/:owner/:repo/stats/punch_card"},
	{http.MethodGet, "/repos/:owner/:repo/statuses/:ref"},
	{http.MethodPost, "/repos/:owner/:repo/statuses/:ref"},

	// Search
	{http.MethodGet, "/search/repositories"},
	{http.MethodGet, "/search/code"},
	{http.MethodGet, "/search/issues"},
	{http.MethodGet, "/search/users"},
	{http.MethodGet, "/legacy/issues/search/:owner/:repository/:state/:keyword"},
	{http.MethodGet, "/legacy/repos/search/:keyword"},
	{http.MethodGet, "/legacy/user/search/:keyword"},
	{http.MethodGet, "/legacy/user/email/:email"},

	// Users
	{http.MethodGet, "/users/:user"},
	{http.MethodGet, "/user"},
	//{http.MethodPatch, "/user"},
	{http.MethodGet, "/users"},
	{http.MethodGet, "/user/emails"},
	{http.MethodPost, "/user/emails"},
	{http.MethodDelete, "/user/emails"},
	{http.MethodGet, "/users/:user/followers"},
	{http.MethodGet, "/user/followers"},
	{http.MethodGet, "/users/:user/following"},
	{http.MethodGet, "/user/following"},
	{http.MethodGet, "/user/following/:user"},
	{http.MethodGet, "/users/:user/following/:target_user"},
	{http.MethodPut, "/user/following/:user"},
	{http.MethodDelete, "/user/following/:user"},
	{http.MethodGet, "/users/:user/keys"},
	{http.MethodGet, "/user/keys"},
	{http.MethodGet, "/user/keys/:id"},
	{http.MethodPost, "/user/keys"},
	//{http.MethodPatch, "/user/keys/:id"},
	{http.MethodDelete, "/user/keys/:id"},
}

var (
	githubAce         http.Handler
	githubBear        http.Handler
	githubBeego       http.Handler
	githubBone        http.Handler
	githubChi         http.Handler
	githubSuperhttp   http.Handler
	githubDenco       http.Handler
	githubEcho        http.Handler
	githubGin         http.Handler
	githubGocraftWeb  http.Handler
	githubGoji        http.Handler
	githubGojiv2      http.Handler
	githubGoJsonRest  http.Handler
	githubGoRestful   http.Handler
	githubGorillaMux  http.Handler
	githubGowwwRouter http.Handler
	githubHttpRouter  http.Handler
	githubHttpTreeMux http.Handler
	githubKocha       http.Handler
	githubLARS        http.Handler
	githubMacaron     http.Handler
	githubMartini     http.Handler
	githubPat         http.Handler
	githubR2router    http.Handler
	githubRevel       http.Handler
	githubRivet       http.Handler
	githubTigerTonic  http.Handler
	githubTraffic     http.Handler
	githubVulcan      http.Handler
	// githubZeus        http.Handler
)

func init() {
	println("#GithubAPI Routes:", len(githubAPI))

	calcMem("Ace", func() {
		githubAce = loadAce(githubAPI)
	})
	calcMem("Bear", func() {
		githubBear = loadBear(githubAPI)
	})
	calcMem("Beego", func() {
		githubBeego = loadBeego(githubAPI)
	})
	calcMem("Bone", func() {
		githubBone = loadBone(githubAPI)
	})
	calcMem("Chi", func() {
		githubChi = loadChi(githubAPI)
	})
	calcMem("Superhttp", func() {
		githubSuperhttp = loadSuperhttp(githubAPI)
	})
	calcMem("Denco", func() {
		githubDenco = loadDenco(githubAPI)
	})
	calcMem("Echo", func() {
		githubEcho = loadEcho(githubAPI)
	})
	calcMem("Gin", func() {
		githubGin = loadGin(githubAPI)
	})
	calcMem("GocraftWeb", func() {
		githubGocraftWeb = loadGocraftWeb(githubAPI)
	})
	calcMem("Goji", func() {
		githubGoji = loadGoji(githubAPI)
	})
	calcMem("Gojiv2", func() {
		githubGojiv2 = loadGojiv2(githubAPI)
	})
	calcMem("GoJsonRest", func() {
		githubGoJsonRest = loadGoJsonRest(githubAPI)
	})
	calcMem("GoRestful", func() {
		githubGoRestful = loadGoRestful(githubAPI)
	})
	calcMem("GorillaMux", func() {
		githubGorillaMux = loadGorillaMux(githubAPI)
	})
	calcMem("GowwwRouter", func() {
		githubGowwwRouter = loadGowwwRouter(githubAPI)
	})
	calcMem("HttpRouter", func() {
		githubHttpRouter = loadHttpRouter(githubAPI)
	})
	calcMem("HttpTreeMux", func() {
		githubHttpTreeMux = loadHttpTreeMux(githubAPI)
	})
	calcMem("Kocha", func() {
		githubKocha = loadKocha(githubAPI)
	})
	calcMem("LARS", func() {
		githubLARS = loadLARS(githubAPI)
	})
	calcMem("Macaron", func() {
		githubMacaron = loadMacaron(githubAPI)
	})
	calcMem("Martini", func() {
		githubMartini = loadMartini(githubAPI)
	})
	calcMem("Pat", func() {
		githubPat = loadPat(githubAPI)
	})
	calcMem("R2router", func() {
		githubR2router = loadR2router(githubAPI)
	})
	// calcMem("Revel", func() {
	// 	githubRevel = loadRevel(githubAPI)
	// })
	calcMem("Rivet", func() {
		githubRivet = loadRivet(githubAPI)
	})
	calcMem("TigerTonic", func() {
		githubTigerTonic = loadTigerTonic(githubAPI)
	})
	calcMem("Traffic", func() {
		githubTraffic = loadTraffic(githubAPI)
	})
	calcMem("Vulcan", func() {
		githubVulcan = loadVulcan(githubAPI)
	})
	// calcMem("Zeus", func() {
	// 	githubZeus = loadZeus(githubAPI)
	// })

	println()
}

// Static
func BenchmarkAce_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubAce, req)
}

func BenchmarkBear_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubBear, req)
}
func BenchmarkBeego_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubBeego, req)
}
func BenchmarkBone_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubBone, req)
}

func BenchmarkChi_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubChi, req)
}
func BenchmarkSuperhttp_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubSuperhttp, req)
}
func BenchmarkDenco_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubDenco, req)
}
func BenchmarkEcho_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubEcho, req)
}
func BenchmarkGin_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubGin, req)
}
func BenchmarkGocraftWeb_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubGocraftWeb, req)
}
func BenchmarkGoji_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubGoji, req)
}
func BenchmarkGojiv2_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubGojiv2, req)
}
func BenchmarkGoRestful_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubGoRestful, req)
}
func BenchmarkGoJsonRest_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubGoJsonRest, req)
}
func BenchmarkGorillaMux_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubGorillaMux, req)
}
func BenchmarkGowwwRouter_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubGowwwRouter, req)
}
func BenchmarkHttpRouter_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubHttpRouter, req)
}
func BenchmarkHttpTreeMux_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubHttpTreeMux, req)
}
func BenchmarkKocha_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubKocha, req)
}
func BenchmarkLARS_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubLARS, req)
}
func BenchmarkMacaron_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubMacaron, req)
}
func BenchmarkMartini_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubMartini, req)
}
func BenchmarkPat_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubPat, req)
}
func BenchmarkR2router_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubR2router, req)
}

//	func BenchmarkRevel_GithubStatic(b *testing.B) {
//		req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
//		benchRequest(b, githubRevel, req)
//	}
func BenchmarkRivet_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubRivet, req)
}
func BenchmarkTigerTonic_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubTigerTonic, req)
}
func BenchmarkTraffic_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubTraffic, req)
}
func BenchmarkVulcan_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
	benchRequest(b, githubVulcan, req)
}

// func BenchmarkZeus_GithubStatic(b *testing.B) {
// 	req, _ := http.NewRequest(http.MethodGet, "/user/repos", nil)
// 	benchRequest(b, githubZeus, req)
// }

// Param
func BenchmarkAce_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubAce, req)
}

func BenchmarkBear_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubBear, req)
}
func BenchmarkBeego_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubBeego, req)
}
func BenchmarkBone_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubBone, req)
}
func BenchmarkChi_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubChi, req)
}
func BenchmarkSuperhttp_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubSuperhttp, req)
}

func BenchmarkDenco_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubDenco, req)
}
func BenchmarkEcho_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubEcho, req)
}
func BenchmarkGin_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGin, req)
}
func BenchmarkGocraftWeb_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGocraftWeb, req)
}
func BenchmarkGoji_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGoji, req)
}
func BenchmarkGojiv2_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGojiv2, req)
}
func BenchmarkGoJsonRest_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGoJsonRest, req)
}
func BenchmarkGoRestful_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGoRestful, req)
}
func BenchmarkGorillaMux_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGorillaMux, req)
}
func BenchmarkGowwwRouter_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGowwwRouter, req)
}
func BenchmarkHttpRouter_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubHttpRouter, req)
}
func BenchmarkHttpTreeMux_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubHttpTreeMux, req)
}
func BenchmarkKocha_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubKocha, req)
}
func BenchmarkLARS_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubLARS, req)
}
func BenchmarkMacaron_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubMacaron, req)
}
func BenchmarkMartini_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubMartini, req)
}
func BenchmarkPat_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubPat, req)
}
func BenchmarkR2router_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubR2router, req)
}

//	func BenchmarkRevel_GithubParam(b *testing.B) {
//		req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
//		benchRequest(b, githubRevel, req)
//	}
func BenchmarkRivet_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubRivet, req)
}

func BenchmarkTigerTonic_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubTigerTonic, req)
}
func BenchmarkTraffic_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubTraffic, req)
}
func BenchmarkVulcan_GithubParam(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubVulcan, req)
}

// func BenchmarkZeus_GithubParam(b *testing.B) {
// 	req, _ := http.NewRequest(http.MethodGet, "/repos/julienschmidt/httprouter/stargazers", nil)
// 	benchRequest(b, githubZeus, req)
// }

// All routes
func BenchmarkAce_GithubAll(b *testing.B) {
	benchRoutes(b, githubAce, githubAPI)
}

func BenchmarkBear_GithubAll(b *testing.B) {
	benchRoutes(b, githubBear, githubAPI)
}
func BenchmarkBeego_GithubAll(b *testing.B) {
	benchRoutes(b, githubBeego, githubAPI)
}
func BenchmarkBone_GithubAll(b *testing.B) {
	benchRoutes(b, githubBone, githubAPI)
}
func BenchmarkChi_GithubAll(b *testing.B) {
	benchRoutes(b, githubChi, githubAPI)
}
func BenchmarkSuperhttp_GithubAll(b *testing.B) {
	benchRoutes(b, githubSuperhttp, githubAPI)
}

func BenchmarkDenco_GithubAll(b *testing.B) {
	benchRoutes(b, githubDenco, githubAPI)
}
func BenchmarkEcho_GithubAll(b *testing.B) {
	benchRoutes(b, githubEcho, githubAPI)
}
func BenchmarkGin_GithubAll(b *testing.B) {
	benchRoutes(b, githubGin, githubAPI)
}
func BenchmarkGocraftWeb_GithubAll(b *testing.B) {
	benchRoutes(b, githubGocraftWeb, githubAPI)
}
func BenchmarkGoji_GithubAll(b *testing.B) {
	benchRoutes(b, githubGoji, githubAPI)
}
func BenchmarkGojiv2_GithubAll(b *testing.B) {
	benchRoutes(b, githubGojiv2, githubAPI)
}
func BenchmarkGoJsonRest_GithubAll(b *testing.B) {
	benchRoutes(b, githubGoJsonRest, githubAPI)
}
func BenchmarkGoRestful_GithubAll(b *testing.B) {
	benchRoutes(b, githubGoRestful, githubAPI)
}
func BenchmarkGorillaMux_GithubAll(b *testing.B) {
	benchRoutes(b, githubGorillaMux, githubAPI)
}
func BenchmarkGowwwRouter_GithubAll(b *testing.B) {
	benchRoutes(b, githubGowwwRouter, githubAPI)
}
func BenchmarkHttpRouter_GithubAll(b *testing.B) {
	benchRoutes(b, githubHttpRouter, githubAPI)
}
func BenchmarkHttpTreeMux_GithubAll(b *testing.B) {
	benchRoutes(b, githubHttpTreeMux, githubAPI)
}
func BenchmarkKocha_GithubAll(b *testing.B) {
	benchRoutes(b, githubKocha, githubAPI)
}
func BenchmarkLARS_GithubAll(b *testing.B) {
	benchRoutes(b, githubLARS, githubAPI)
}
func BenchmarkMacaron_GithubAll(b *testing.B) {
	benchRoutes(b, githubMacaron, githubAPI)
}
func BenchmarkMartini_GithubAll(b *testing.B) {
	benchRoutes(b, githubMartini, githubAPI)
}
func BenchmarkPat_GithubAll(b *testing.B) {
	benchRoutes(b, githubPat, githubAPI)
}
func BenchmarkR2router_GithubAll(b *testing.B) {
	benchRoutes(b, githubR2router, githubAPI)
}

//	func BenchmarkRevel_GithubAll(b *testing.B) {
//		benchRoutes(b, githubRevel, githubAPI)
//	}
func BenchmarkRivet_GithubAll(b *testing.B) {
	benchRoutes(b, githubRivet, githubAPI)
}

func BenchmarkTigerTonic_GithubAll(b *testing.B) {
	benchRoutes(b, githubTigerTonic, githubAPI)
}
func BenchmarkTraffic_GithubAll(b *testing.B) {
	benchRoutes(b, githubTraffic, githubAPI)
}
func BenchmarkVulcan_GithubAll(b *testing.B) {
	benchRoutes(b, githubVulcan, githubAPI)
}

// func BenchmarkZeus_GithubAll(b *testing.B) {
// 	benchRoutes(b, githubZeus, githubAPI)
// }
