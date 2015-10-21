A simple jspm + Babel + React repository to test performance of loading many JS files in Chrome on Mac OS X. Times are measured with Date.now()-[PerformanceTiming.fetchStart](https://developer.mozilla.org/en-US/docs/Web/API/PerformanceTiming).

### Summary
JSPM with DevTools open causes ~20ms extra load time per JS file. A simple JS XHR performance test saw a 1 to 2.5ms increase with DevTools open, implying that there is something else SystemJS is doing that is triggering a significant performance hit in DevTools.

Feedback and replications are welcome.

### jspm test

I've tested numerous times now and consistently getting the below figures. Interestingly, **the "Disable cache" option in DevTools Network tab** does not have a noticeable impact on load times. Cache is verified by the "(from cache)" in the Size column.

Load time ranges from the simple jspm app (in the jspm folder):

**DevTools not open**
- http: 650-720ms
- h2: 689-719ms

**DevTools open to Elements tab**
- http: 3591-3850ms
- h2: 3594-3838ms

**DevTools open to Network tab**
- http: 3970-4196ms
- h2: 4070-4206ms

Caddy was tested and the numbers fell within the above ranges.
Browsersync was tested and was 30-200ms slower than the above ranges.

Web server commands (run from the public dir):
- serve.go (simple Go server with HTTP and H2): `go run serve.go`
- [Wago](https://github.com/JonahBraun/wago):  `wago -http :8420 -h2 :8421`
- [Caddy](https://caddyserver.com/): `caddy`
- [Browsersync](http://www.browsersync.io/): `browser-sync start --server`

### XHR test

In response to issue #1, I added an XHR test. This makes 165 iterative (not parallel) XHR requests for a 5KB JS payload. Nothing is done with the JS file. Average times:

- DevTools closed: 520ms
- DevTools open to Elements tab: 710ms (1.15ms/file increase)
- DevTools open to Network tab: 861ms (2.07ms/file increase)


### Conclusion

Just having DevTools open imposes a significant performance hit when loading many JS files with SystemJS. Initially, I thought this was the XHR requests (in this case 165), but the XHR test points to the issue being elsewhere.

HTTP2 actually resulted in slightly slower times. I hypothesize that the overhead of TLS costs more than the benefit of using a single connection when talking to localhost. This test would need to be replicated over the Internet to verify this.
