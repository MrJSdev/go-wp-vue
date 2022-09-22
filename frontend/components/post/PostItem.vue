<template>
<div  class="px-5 py-5 bg-white rounded-md m-2 shadow-md">
      <h2 class="text-2xl font-bold mb-5">
        {{ post.post_title }}
      </h2>

      <span v-html="post.post_content"></span>
    </div>

</template>

<script lang="ts" setup>
import { Post } from '@/types/Post'

defineProps<{
  post: Post;
}>();

</script>