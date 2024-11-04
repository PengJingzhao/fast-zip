import { defineConfig } from "@rsbuild/core";
import { pluginVue } from "@rsbuild/plugin-vue";

export default defineConfig({
    plugins: [pluginVue()],
    html: {
        template: './index.html',
    },
    source: {
        entry: {
            index: './src/main.js',
        },
    },
    server: {
        port: 1420,
    }
});