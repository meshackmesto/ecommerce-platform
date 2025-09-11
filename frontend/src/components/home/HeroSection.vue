// File: frontend/src/components/home/HeroSection.vue
<template>
  <div class="relative overflow-hidden pt-16 pb-24">
    <div class="container mx-auto px-4">
      <div class="flex flex-col md:flex-row items-center">
        <div class="md:w-1/2 text-center md:text-left">
          <h1 class="text-4xl md:text-6xl font-bold text-gray-800 leading-tight">
            The one stop<br />
            <span class="text-teal-500">TECH STORE</span>
          </h1>
          <p class="mt-6 text-gray-600 max-w-md mx-auto md:mx-0">
            High class, high quality gadgets at affordable prices. All your desktop or laptop needs made available.
          </p>
          <div class="mt-8 flex justify-center md:justify-start space-x-4">
            <router-link to="/products" class="bg-gray-800 text-white font-semibold py-3 px-8 rounded-lg hover:bg-gray-900 transition duration-300">
              Buy Now
            </router-link>
            <router-link to="/products" class="bg-white text-gray-800 font-semibold py-3 px-8 rounded-lg border border-gray-300 hover:bg-gray-100 transition duration-300">
              Shop All
            </router-link>
          </div>
        </div>
        <div class="md:w-1/2 mt-12 md:mt-0 flex justify-center">
          <img src="https://images.unsplash.com/photo-1542496658-e33a6d0d50f6?q=80&w=2070" alt="Tech Store Products" class="rounded-lg shadow-2xl w-4/5 transform md:rotate-6 transition duration-500 hover:rotate-0 hover:scale-105" />
        </div>
      </div>
    </div>
  </div>
</template>