<template>
  <header class="container mx-auto py-4">

    <div class="flex">
      <div class="flex-1">
        <!-- Logo -->
        <h1 class="text-2xl text-uppercase">Logo</h1>
      </div>
      <!-- Menu list -->
      <nav class="flex-1 text-right">
        <ul class="justify-end">
          <NuxtLink class="px-3 py-2 bg-gray-100 text-blue-600 rounded-md mr-3 " v-for="navItem in navigationList"
            :key="navItem.path" :to="navItem.path">
            {{ navItem.title}}
          </NuxtLink>
        </ul>
      </nav>
    </div>
  </header>
</template>

<script lang="ts" setup>
import { NavigationItem } from '@/types/Navigation'

const navigationList: NavigationItem[] = [
  {
    title: 'Home',
    path: '/home'
  },
  {
    title: 'Blog',
    path: '/blog'
  },
  {
    title: 'Post item',
    path: '/blog/2'
  }
] 
</script>