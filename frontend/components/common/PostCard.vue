<template>
<div class="border mb-5 border-gray-200 bg-white p-5 rounded">
  <h2 class="text-2xl font-bold mb-5" v-if="title">
    {{ title }}
  </h2>
  <slot />
</div>
  
</template>

<script lang="ts" setup>
defineProps<{
  title?: string;
}>()
  </script>