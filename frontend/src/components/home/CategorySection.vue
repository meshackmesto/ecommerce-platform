// File: frontend/src/components/home/CategorySection.vue
<template>
  <div class="py-24 bg-white">
    <div class="container mx-auto px-4">
      <div class="text-center mb-12">
        <h2 class="text-4xl font-bold text-gray-800">Shop by category</h2>
        <p class="text-gray-600 mt-2">Find the best tech for your needs.</p>
      </div>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <div class="group relative overflow-hidden rounded-lg shadow-lg cursor-pointer">
          <img src="https://images.unsplash.com/photo-1605236453805-6ff5365673a2?q=80&w=1974" class="w-full h-96 object-cover transform group-hover:scale-110 transition duration-500" />
          <div class="absolute inset-0 bg-black bg-opacity-40 flex items-end">
            <h3 class="text-white text-2xl font-bold p-6">PHONES & ACCESSORIES</h3>
          </div>
        </div>
        <div class="group relative overflow-hidden rounded-lg shadow-lg cursor-pointer">
          <img src="https://images.unsplash.com/photo-1517336714731-489689fd1ca8?q=80&w=1926" class="w-full h-96 object-cover transform group-hover:scale-110 transition duration-500" />
          <div class="absolute inset-0 bg-black bg-opacity-40 flex items-end">
            <h3 class="text-white text-2xl font-bold p-6">COMPUTER EQUIPMENT</h3>
          </div>
        </div>
        <div class="group relative overflow-hidden rounded-lg shadow-lg cursor-pointer">
          <img src="https://images.unsplash.com/photo-1505740420928-5e560c06d30e?q=80&w=2070" class="w-full h-96 object-cover transform group-hover:scale-110 transition duration-500" />
          <div class="absolute inset-0 bg-black bg-opacity-40 flex items-end">
            <h3 class="text-white text-2xl font-bold p-6">MISCELLANEOUS ITEMS</h3>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>