// File: frontend/src/components/common/Loading.vue
<template>
    <div class="flex flex-col items-center justify-center py-10">
        <div class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-teal-500"></div>
        <p class="mt-4 text-gray-600">{{ message }}</p>
    </div>
</template>

<script setup>
defineProps({
    message: {
        type: String,
        default: 'Loading...'
    }
})
</script>