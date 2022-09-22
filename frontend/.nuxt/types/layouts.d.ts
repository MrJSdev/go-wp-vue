import { ComputedRef, Ref } from 'vue'
export type LayoutKey = "default"
declare module "/Users/riyaz/go/src/go-wp-vue/frontend/node_modules/nuxt/dist/pages/runtime/composables" {
  interface PageMeta {
    layout?: false | LayoutKey | Ref<LayoutKey> | ComputedRef<LayoutKey>
  }
}