<template>
  <NuxtLoadingIndicator />
  <Navigation />
  <div class="bg-blue-50">
    <div class="container mx-auto">
      <div class="flex">

        <div class="w-3/4">
          <slot />
        </div>
        <div class="w-1/4">
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