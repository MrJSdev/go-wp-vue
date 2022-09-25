<template>
  <form class="shadow-sm bg-gray-100 rounded px-8 pt-6 pb-8 mb-4">
    <h3 class=" text-2xl font-semibold">How was it?</h3>
    <p class=" text-gray-800 my-2">Your email address will not be published. Required fields are marked *</p>
    <div class="mb-4 w">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="comment">
        Comment
      </label>
      <textarea v-model="commentForm.comment_content"
        class="shadow appearance-none border rounded h-20 py-3 px-4 w-full text-gray-700 leading-tight focus:outline-none focus:shadow-outline" placeholder="comment" />
    </div>
    <div class="mb-4">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="name">
        Name
      </label>
      <input v-model="commentForm.comment_author" class="shadow appearance-none border rounded w-full py-3 px-4 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" type="text" placeholder="Name">
    </div>
    <div class="mb-4">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="name">
        Email*
      </label>
      <input v-model="commentForm.comment_author_email" class="shadow appearance-none border rounded w-full py-3 px-4 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" type="text" placeholder="Email">
    </div>
    
    <div class="flex justify-end">
      <button  class="bg-primary hover:bg-green-600 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline" type="button" @click="submitComment">
        {{ isLoading ? 'Loading...' : 'Submit' }}
      </button>
    </div>
  </form>

</template>

<script lang="ts" setup>
import { PostComment } from '@/types/Post';

const props = defineProps<{
  parentID?: number;
  postID?: number;
}>();

const commentForm = ref({} as PostComment)
const isLoading = ref(false)


const submitComment = () => {
  isLoading.value = true;
  try{
    const { data } = useFetch(`http://localhost:8081/api/comments/${props.postID}`, {
      method: 'POST',
      body: JSON.stringify({...commentForm.value, comment_parent: props.parentID, comment_post_id: props.postID })
    })

  } catch(err) {
    console.log(err);
  } finally {
    isLoading.value = false;
  }
}
</script>