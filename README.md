# Nutri-Meal Application

A web application for managing nutritional goals, meal planning, and community engagement.

## Table of Contents

1. [Project Structure](#project-structure)
2. [Features](#features)
3. [Tech Stack](#tech-stack)
4. [Setup Instructions](#setup-instructions)
5. [Environment Variables](#environment-variables)
6. [API Endpoints](#api-endpoints)
7. [Frontend Routes](#frontend-routes)
8. [Contributing](#contributing)
9. [License](#license)

---

## Project Structure

```
nutri-meal/
├── backend/
│   ├── handlers/
│   │   └── ...
│   ├── models/
│   └── database/
└── frontend/
```

## Features

- **Health Goals**: Set and track personal health objectives.
- **Meal Planning**: Generate daily or weekly meal plans with recipes.
- **Community**: Share posts, like, and interact with other users.
- **Food Logs**: Track daily food intake and nutritional values.
- **Progress Tracking**: Visualize health metrics over time.

## Tech Stack

- **Frontend**: React, TailwindCSS, Vite
- **Backend**: Go (Fiber framework)
- **Database**: MongoDB
- **Authentication**: Google OAuth, Facebook Login
- **Image Hosting**: Imgur API

## Setup Instructions

### Backend

1. Clone the repository:
   ```bash
   cd nutri-meal
   ```

2. Install dependencies:
   ```bash
   # Backend (Go)
   cd backend
   go mod download
   ```

## Environment Variables

```env
MONGODB_URI=mongodb_connection_string
PORT=8080
```

## API Endpoints

### Health Goals

- `POST /health-goals` - Create a new health goal
- `GET /health-goals/:user_id` - Retrieve a user's health goal
- `PUT /health-goals/:user_id` - Update a user's health goal
- `GET /health-goals/user/:userId` - Get health goals by user ID

## Frontend Routes

----



