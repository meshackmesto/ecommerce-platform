// File: frontend/src/main.js (Corrected Order)
import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import { setupInterceptors } from './services/api'
import './assets/css/main.css'

const app = createApp(App)

// 1. Install Pinia FIRST. This makes all the stores available.
app.use(createPinia())

// 2. Install the Router SECOND. Now the router's navigation guards can access the stores.
app.use(router)

setupInterceptors(router)

app.mount('#app')