var CACHE_NAME = "sacer";
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

const stripPath = url => {
  return url.replace(/^.*\/\/[^\/]+/, "")
}

// Network falling back to the cache
self.addEventListener('fetch', function(event) {
  const path = stripPath(event.request.url);
  const ext = path.substr(-3);
  if (ext == "mp4" || ".ts") {
    return false;
  }
  event.respondWith(
    fetch(event.request).then(response => {
      if (path == "/" || path == "manifest.json") {
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

