import { defineConfig } from 'vite';
import { vite as vidstack } from 'vidstack/plugins';

// https://vitejs.dev/config/
export default defineConfig({
  //build: { cssMinify: false },
  "base": "/static/js/video",
  build: {
    watch: {

    },
    cssMinify: false,
    assetsDir: "",
    manifest: true,
    rollupOptions: {
      // overwrite default .html entry
      input: 'src/main.ts',
    },
  },
  plugins: [vidstack()],
});
