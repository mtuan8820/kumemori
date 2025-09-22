# 1. Project name: kumemori

# 2. Purpose
- To develope coding skill, enrich my CV
- I'm currently use Quizlet (free version) for learning Japanese but it does not have the "real" spaced repetition feature. Anki on the other hands, provide that feature but does not provide Quiz feature like quizlet. -> i want to create ones for me
- Who it’s for: Me, and those who try to learn new language (or new knowledege)

# 3. Goals
Efficient flashcard review with spaced repetition.

Fun, engaging quiz mode for revision.

Simple desktop app (no server dependency).

# 4. Non-goal
A web version (not in scope for now)

Mobile app (not in scope for now).

Cloud sync / multi-device support.

# 5. High-level Features
- Card listing & filtering
- Space repetition
- Quiz generated from deck
- Configurable learning settings (daily goal, review batch size, quiz options).

# 6. Architecture at a Glance

The application is built using Wails, which bundles a Go backend with a Vue.js frontend into a desktop application.

- Frontend (Vue.js):

  - Provides the user interface and handles user interactions.

  - Communicates with the backend via Wails’ runtime bridge instead of traditional REST API calls.

- Backend (Go):

  - Lives inside the /internal folder.

  - Implements business logic, validation, and application services.

  - Exposes methods callable from the Vue frontend through Wails.

- Database (SQLite):

  - Stores all persistent data locally.

  - Accessed directly by the Go backend using an ORM or SQL queries.

  - No direct access from the frontend.

# 7. Tech Stack

Frontend: Vue.js + TailwindCSS (bundled by Wails)

Backend: Go (inside `/internal`)

Database: SQLite (lightweight local persistence)

Framework / Runtime: Wails (integrates Go backend with Vue frontend, builds desktop app)

Version Control: Git + GitHub 

Testing: Go `testing` package


# 8. Roadmap / Milestones
- Phase 1 – Foundation

  - Initialize repo structure (frontend/, internal/, build/, docs/).

  - Set up Wails project (Go backend + Vue frontend).

  - Configure TailwindCSS for styling.

- Initialize SQLite database with schema for decks, cards, and reviews.

- Phase 2 – Core Flashcard System

  - Create basic Deck & Card Management (CRUD).

  - Implement Card Listing & Filtering.

  - Establish data flow between Vue frontend ↔ Go backend ↔ SQLite.

- Phase 3 – Learning Modes

  - Implement Spaced Repetition Engine (review scheduler).

  - Create Quiz Mode (auto-generate quizzes from decks).

  - Add Custom Settings (e.g., review batch size, daily goal).

- Phase 4 – User Experience

  - Improve UI/UX with Tailwind components.

  - Error handling, validation, and loading states.

  - Add keyboard shortcuts for faster review.

- Phase 5 – Testing & Release

  - Unit tests (Go backend + Vue components).

  - Integration testing for card review flow.

  - Package desktop builds for Windows, macOS, Linux.

  - Polish documentation in /docs.