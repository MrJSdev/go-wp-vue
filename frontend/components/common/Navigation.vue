<template>
  <header class="w-full bg-white border md:mb-5 border-gray-100 rounded-md py-4">

    <div class="flex">
      <div class="flex-1">
        <!-- Logo -->
        <NuxtLink to="/">
          <h1 class="text-2xl cursor-pointer text-uppercase pl-5">Logo</h1>
        </NuxtLink>

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
    path: '/'
  },
  {
    title: 'Blog',
    path: '/blog'
  },
] 
</script>