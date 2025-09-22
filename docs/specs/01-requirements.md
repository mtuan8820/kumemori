# 1. Project Context
Help users creating decks and adding cards that contain new term that they want to learn.

Apply Space repetition to optimize the memorization and review process.

Quiz Mode with multiple-choice questions provides an interactive way for users to test their knowledge.
# 2. Stakeholders
Product Owner / Developer / End User: Me

- Defines requirements.

- Implements and maintains the system.

- Uses the application for personal learning purposes.

# 3. User story

## 3.1 Deck Management

    - As a user, I want to create a new deck, so that I can group related terms together.
    - As a user, I want to edit or delete a deck, so that I can keep my collection organized.
    - As a user, I want to view a list of my decks, so that I can quickly access the one I need.

## 3.2 Card Management
    - As a user, I want to add a card to a deck, so that I can store a new term with its meaning.
    - As a user, I want to edit or delete a card, so that I can fix mistakes or remove unnecessary terms.
    - As a user, I want to see all cards in a deck, so that I can review them before studying.

## 3.3 Learning with SRS
    - As a user, I want to study cards using spaced repetition, so that I review them at the best time for memory retention.
    - As a user, I want the system to track my study progress (due date, repetitions, lapses, ease factor, etc.), so that my review schedule is optimized automatically.

## 3.4 Quiz Mode
    - As a user, I want to start a quiz from a deck, so that I can test my knowledge interactively.
    - As a user, I want to answer multiple-choice questions, so that I can practice recognition and recall.
    - As a user, I want to get immediate feedback (correct/incorrect), so that I know if I was right.
    - As a user, I want to see a summary of my quiz results (score, accuracy, time per question), so that I can evaluate my performance.

## 3.5 General
    - As a user, I want to see my study history (last studied, progress stats), so that I stay motivated.
    - As a user, I want the application to be fast and easy to use, so that studying feels smooth.

# 4. Functional Requirements
## 4.1 Deck Management

- The system shall allow users to create, view, edit, and delete decks.

- Each deck shall have a unique name and an optional description.

## 4.2 Card Management

- The system shall allow users to add, view, edit, and delete cards within a deck.

- Each card contain a front (question/term) and a back (answer/definition).

- The system shall associate each card with exactly one deck.

## 4.3 Spaced Repetition Study

- The system shall schedule cards for review using the spaced repetition algorithm.

- The system shall track review metadata for each card (e.g., repetitions, lapses, ease factor, interval, due date).

- The system shall update scheduling parameters based on user feedback after studying.

## 4.4 Quiz Mode

- The system shall allow the user to start a quiz from a selected deck.

- For each question, the system shall display the card’s front and provide multiple-choice answers (one correct, several distractors).

- The system shall allow the user to select one answer per question.

- The system shall provide immediate feedback (correct/incorrect) after each question.

- The system shall record quiz results (e.g., correct answers, incorrect answers, time per question).

- The system shall generate a quiz summary (score, accuracy, performance stats).

- The system may integrate quiz results into the spaced repetition algorithm to adjust scheduling. (optional)

# 5. Non-functional Requirements
The system shall load the deck list within 2 seconds under normal conditions.

The system shall support quizzes of up to 100 cards without noticeable lag.

The user interface shall be simple and intuitive, requiring no training for a first-time user.

The system shall provide responsive design for desktop.

The system shall not lose data in case of unexpected shutdowns.

The system shall include automated tests for core functionality (deck management, card management, spaced repetition, quiz mode).

# 6. Glossary

Deck: A collection of cards grouped by topic or purpose

Card: The basic study unit, containing two sides:

    Front: a term, question, or keyword.

    Back: The answer or explanation.

Spaced Repetition (SRS): A learning technique that schedules reviews of cards at increasing intervals to optimize long-term retention.