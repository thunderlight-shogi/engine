import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import * as path from 'path';

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue()],
    resolve: {
        alias: {
          '@': path.resolve(__dirname, 'src'),
        }
      },
    css: {
        preprocessorOptions: {
            sass: {
                additionalData: `@import "@/assets/global.sass"`
            }
        }
    },
    optimizeDeps: {
        exclude: ['js-big-decimal']
    }
})
