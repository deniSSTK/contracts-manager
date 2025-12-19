import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import * as path from "node:path";

// https://vite.dev/config/
export default defineConfig({
    plugins: [vue()],
    server: {
        host: true,
        port: 5173,
        watch: {
            usePolling: true,
        },
    },
    resolve: {
        alias: {
            '@model': path.resolve(__dirname, 'src/models'),
            '@store': path.resolve(__dirname, 'src/stores'),
            '@usecase': path.resolve(__dirname, 'src/usecases'),
            '@component': path.resolve(__dirname, 'src/components'),
            '@view': path.resolve(__dirname, 'src/views'),
            '@repository': path.resolve(__dirname, 'src/infrastructure/api/repositories'),
            '@utils': path.resolve(__dirname, 'src/utils'),
        }
    }
})
