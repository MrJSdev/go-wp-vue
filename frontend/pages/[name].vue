<template>
  <div>

    <PostSingle v-if="post" :post="post" />

    <PostRelated />

    <PostComments class="mt-5" :postID="post.id" :comments="post.group_comments" />
  </div>
</template>
<script lang="ts" setup>
import { Post } from '@/types/Post';
import PostComments from '@/components/post/PostComments.vue';
import PostSingle from '@/components/post/PostSingle.vue';
import PostRelated from '@/components/post/PostRelated.vue';
const route = useRoute()


const postName = route.params.name;
const { data: post } = await useFetch<Post>(`http://localhost:8081/api/posts/${postName}`);

useHead({
  title: post.value.post_title,
  meta: [
    {
      name: 'description',
      content: post.value.post_excerpt,
    },
  ],
})

</script>