// Template from here: https://developers.google.com/web/fundamentals/primers/service-workers

var CACHE_NAME = "stferal";
var urlsToCache = [
  "/",
  "/manifest.json"
];

self.addEventListener("install", function(event) {
  event.waitUntil(
    caches.open(CACHE_NAME)
      .then(function(cache) {
        return cache.addAll(urlsToCache);
      })
  );
});

const path = url => {
  return url.replace(/^.*\/\/[^\/]+/, "")
}

let fetchedOption = false;

self.addEventListener("fetch", function(event) {
  const req = path(event.request.url);

  //  see if option was requested
  if (req.substr(0, 5) == "/opt/") {
    fetchedOptions = true;
  }

  event.respondWith(
    caches.match(event.request)
      .then(function(response) {

        // dont serve index from cache if option was requested
        if (response && req != "/" && !fetchedOption) {
          return response;
        }
        return fetch(event.request).then(
          response => {

            if(!response || response.status !== 200 || response.type !== 'basic') {
              return response;
            }

            // cache new index after option was requested
            if (req == "/" && fetchedOption) {
              let responseToCache = response.clone();
              caches.open(CACHE_NAME)
                .then(cache => {
                  cache.put(event.request, responseToCache);
                  fetchedOption = false;
                });
            }

            return response;
          }
        );
      })
    );
});
