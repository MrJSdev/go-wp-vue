<template>
  <div class="p-6 bg-gray-100 w-full mb-5 rounded-xl border-gray-100 border flex ">
    <div class="shrink-0 mr-5">
      <img class="h-12 w-12" alt="ChitChat Logo">
    </div>
    <div>
      <div class="text-xl font-medium text-black">{{ comment.comment_author }}</div>
      <p class="text-slate-500">{{ comment.comment_content }}

        <!-- Temporary comment  -->
      Lorem ipsum dolor sit amet consectetur adipisicing elit. Aspernatur enim, modi nisi eius, quam saepe aut repellendus tempore odit molestias vero minima perspiciatis dicta laboriosam cupiditate facilis illum. Saepe, labore.
      </p>


      <slot />

      <button class="text-primary font-medium py-2 focus:outline-none focus:shadow-outline" type="button" @click="isCommentFormVisible = !isCommentFormVisible">
        {{ isCommentFormVisible ? 'Cancel' : 'Reply' }}
      </button>

      <PostCommentsForm v-if="isCommentFormVisible" :parent-id="comment.comment_parent" />

    </div>
  </div>
</template>

<script lang="ts" setup>
import { PostComment } from '@/types/Post';
import PostCommentsForm from '@/components/post/PostCommentsForm.vue';

defineProps<{
  comment: PostComment;
}>()

const isCommentFormVisible = ref(false);

</script>