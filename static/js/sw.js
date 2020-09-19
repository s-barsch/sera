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

// Network falling back to the cache
self.addEventListener('fetch', function(event) {
  const req = path(event.request.url);
  event.respondWith(
    fetch(event.request).then(response => {
      if (req == "/" || req == "manifest.json") {
        let responseToCache = response.clone();
        caches.open(CACHE_NAME)
          .then(cache => {
            cache.put(event.request, responseToCache);
          });
      }
      return response;
    }).catch(function() {
      return caches.match(event.request);
    })
  );
});

