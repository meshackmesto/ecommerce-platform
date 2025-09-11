<script setup>
import { RouterView } from 'vue-router'
import Header from '@/components/common/Header.vue'
import Footer from '@/components/common/Footer.vue'
</script>

<template>
  <Header />
  <main class="container mx-auto p-4">
    <RouterView />
  </main>
  <Footer />
</template>