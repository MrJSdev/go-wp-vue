<template>
  <div class="grid grid-cols-2">
    <PostItem v-for="post in posts" :post="post" :key="post.id" />
  </div>
</template>

<script lang="ts" setup>
import { Post } from '@/types/Post'
import PostItem from '@/components/post/PostItem.vue';
const { data: posts } = await useFetch<Post[]>('http://localhost:8081/api/posts');
</script>