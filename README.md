🛒 #Full-Stack E-Commerce Platform

#A production-ready e-commerce MVP built with Go (Fiber), Vue.js, and PostgreSQL, developed in 4 days with clean architecture, JWT authentication, and simulated M-Pesa payments.

🚀 Features
🔑 Authentication & Authorization

JWT-based authentication

Role-based access (Admin / Customer)

Protected routes & middleware

📦 Product Management

CRUD operations for products

Stock & inventory management

Image upload handling

Admin-only product controls

🛍️ E-Commerce Flow

Product browsing & filtering

Shopping cart functionality

Checkout with order validation

Order history & tracking

💳 Payment Integration

Mock M-Pesa STK Push API

Webhook handling & transaction storage

Real-time payment status updates via WebSockets

🖥️ Frontend (Vue.js + Tailwind)

Responsive, mobile-first UI

Separate Admin & Customer interfaces

Real-time updates (payments/orders)

Modern UI with reusable components

⚡ Backend (Go + Fiber)

REST API with clean architecture

Repository & service pattern

PostgreSQL with migrations

Secure middleware (CORS, Auth, Admin)

📂 Project Structure
ecommerce-platform/
├── backend/       # Go + Fiber API
├── frontend/      # Vue.js + Tailwind frontend
├── database/      # docker-compose for PostgreSQL & pgAdmin
├── docs/          # API docs & setup guides
└── README.md

🛠️ Tech Stack

Backend

Go (Fiber)
 – fast, minimal web framework

PostgreSQL
 – relational database

JWT & bcrypt for authentication/security

Frontend

Vue 3
 – reactive frontend framework

Pinia
 – state management

Tailwind CSS
 – modern styling

Axios for API requests, WebSocket for real-time updates

Other

Docker
 – containerization

pgAdmin – database UI

Mock M-Pesa payment API

⚙️ Setup & Installation
1. Clone the Repository
git clone https://github.com/yourusername/ecommerce-platform.git
cd ecommerce-platform

2. Backend Setup
cd backend
cp .env.example .env   # configure DB + JWT_SECRET
go mod tidy
go run cmd/server/main.go

3. Database with Docker
cd database
docker-compose up -d


This starts PostgreSQL (5432) and pgAdmin (5050).

pgAdmin login → admin@pgadmin.com / admin123

4. Frontend Setup
cd frontend
npm install
npm run dev


Frontend runs at → http://localhost:5173

🔑 Default Accounts (Seeded)
Admin:
  Email: admin@ecommerce.com
  Password: admin123

Customer:
  Email: customer@ecommerce.com
  Password: customer123

📖 API Documentation

See docs/API.md
 for full endpoint reference.

🎯 Roadmap

 Add Stripe/PayPal integration

 Unit & integration testing

 Deploy on Render/Heroku with CI/CD

 Analytics dashboard with charts

📸 Screenshots

(Add screenshots of your UI: login, products, admin panel, checkout, etc.)

🤝 Contributing

Pull requests are welcome. For major changes, open an issue first to discuss.

📜 License

MIT License © 2025 [Your Name]
