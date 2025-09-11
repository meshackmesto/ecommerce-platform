// File: frontend/src/components/common/Footer.vue
<template>
  <footer class="bg-gray-800 text-white py-16">
    <div class="container mx-auto px-4 grid grid-cols-1 md:grid-cols-3 gap-12 text-center md:text-left">
      <div>
        <h2 class="text-2xl font-bold">Gadget & Grove</h2>
        <p class="text-gray-400 mt-2">Your one stop tech store.</p>
      </div>
      <div>
        <h3 class="text-lg font-semibold mb-4">MAIN OFFICE</h3>
        <p class="text-gray-400">123 Anywhere St.</p>
        <p class="text-gray-400">Any City, State</p>
        <p class="text-gray-400">Any Country</p>
        <p class="text-gray-400 mt-2">(123) 456 7890</p>
      </div>
      <div>
        <h3 class="text-lg font-semibold mb-4">SOCIAL MEDIA</h3>
        <a href="#" class="block text-gray-400 hover:text-white transition">Facebook</a>
        <a href="#" class="block text-gray-400 hover:text-white transition mt-1">Instagram</a>
        <a href="#" class="block text-gray-400 hover:text-white transition mt-1">Twitter</a>
      </div>
    </div>
    <div class="container mx-auto px-4 mt-12 border-t border-gray-700 pt-8 text-center text-gray-500">
      <p>&copy; 2025 Gadget & Grove. All Rights Reserved.</p>
    </div>
  </footer>
</template>