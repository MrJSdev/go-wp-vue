<template>
  <div class="bg-gray-100 h-full w-full">
    <NuxtLoadingIndicator />

    <div class="layout-container pt-5 mx-auto">
    <Navigation />

      <div class="flex gap-2">
        <div class="flex-grow">
          <slot />
        </div>

        <div class="mx-w-80">
          <PostSidebar />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>

import Navigation from '@/components/common/Navigation.vue';
import PostSidebar from '@/components/post/PostSidebar.vue';

</script> 
<style scoped>
.layout-container{
  max-width: 71.25rem;
}
</style>