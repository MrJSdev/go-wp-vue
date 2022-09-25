<template>
  <PostCard title="Opinions">
  <PostCommentsItem v-for="comment in comments" :key="comment.comment_id" :comment="comment">

    <!-- Reply of comments -->
    <template v-if="comment.children">
      <PostCommentsItem v-for="childComment in comment.children" :comment="childComment"
        :key="childComment.comment_id" />

        
      
    </template>

  </PostCommentsItem>
  <!-- Comment form -->
  <PostCommentsForm :postID="postID" />
</PostCard>
</template>

<script lang="ts" setup>
import { PostComment } from '@/types/Post';
import PostCommentsItem from '@/components/post/PostCommentsItem.vue';
import PostCommentsForm from '@/components/post/PostCommentsForm.vue';
import PostCard from '@/components/common/PostCard.vue';

defineProps<{
  comments: PostComment[];
  postID: number;
}>()

</script>