A simple jspm + React repository to test performance of loading many JS files.

I'm using [Wago](https://github.com/JonahBraun/wago) to serve files over HTTP and HTTP2 and measuring from [PerformanceTiming.fetchStart](https://developer.mozilla.org/en-US/docs/Web/API/PerformanceTiming).

Wago command (run from repository root): `wago -dir public -http :8420 -h2 :8421`

In Chrome on MacOSX, the following numbers were obtained by performing a hard refresh (cmd-shift-r) many times in a row and reporting the typical range.

**Developer Tools not open**
- http: 650-720ms
- h2: 689-719ms

**Developer Tools open to Elements tab**
- http: 3591-3850ms
- h2: 3594-3838ms

**Developer Tools open to Network tab**
- http: 3970-4196ms
- h2: 4070-4206ms

### Conclusions

Just having Tools open imposes a significant performance hit when making many XHR requests (in this case 165).

HTTP2 actually resulted in slightly slower times. I hypothesize that the overhead of TLS costs more than the benefit of using a single connection when talking to localhost. This test would need to be replicated over the Internet to verify this.

Feel free to open issues with questions / comments / reproduced results. Thanks!
