# Test Social

Un projet d'application sociale avec un backend Go et un frontend Vue.js.

## Structure du projet

```
backend/          # Backend en Go
├── main.go      # Point d'entrée principal
├── go.mod       # Dépendances Go
├── go.sum       # Versions des dépendances
└── pkg/         # Packages du projet
    ├── db/      # Couche base de données
    └── handlers/ # Gestionnaires HTTP

frontend/         # Frontend Vue.js
├── src/         # Code source Vue
├── public/      # Fichiers publics
├── package.json # Dépendances Node.js
└── vite.config.js # Configuration Vite
```

## Installation

### Backend
```bash
cd backend
go mod download
go run main.go
```

### Frontend
```bash
cd frontend
npm install
npm run dev
```

## Développement

- Backend : http://localhost:8080
- Frontend : http://localhost:5173
