<template>
  <div class="container mx-auto">
    <h2 class="text-orange-500 text-bold">Blog page</h2>
    <p>This is my blog page</p>
    <div class="grid grid-cols-2">
      <div v-for="post in posts" class="px-5 py-5 rounded-md m-2 shadow-md">
      <span v-html="post.post_content"></span>
    </div>
    </div>
  </div>

</template>

<script lang="ts" setup>
import { Post } from '@/types/Post'

const { data: posts } = await useFetch<Post[]>('http://localhost:8081/api/posts');
</script>