<template>
  <PostCard >
  <h2 class="text-2xl font-bold mb-5">
    {{ post.post_title }}
  </h2>

  <div class="content">
    <span v-html="post.post_content"></span>
  </div>
</PostCard>

</template>

<script lang="ts" setup>
import PostCard from '@/components/common/PostCard.vue';
import { Post } from '@/types/Post';

defineProps<{
  post: Post;
}>()
</script>