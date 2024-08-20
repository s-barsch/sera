var CACHE_NAME = "sacer";
var urlsToCache = [
  "/static/offline.html",
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
  if (ext == "mp4") {
    return false;
  }
  event.respondWith(
    fetch(event.request)
      .then(response => {
        if (path == "/static/offline.html") {
          caches.open(CACHE_NAME)
            .then(cache => {
              cache.put(event.request, response.clone());
            });
        }
        return response;
      }).catch(function() {
        console.log("catched fetch");
        if (path === "/") {
          return caches.match("/static/offline.html");
        }
        return caches.match(event.request);
      })
  );
});

