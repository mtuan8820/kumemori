# ADR: Transaction Management for Deck Create/Update

## Context
The `Deck` aggregate consists of a deck entity and its related cards.  
Typical operations:
- Create Deck: Insert into `deck` table, then insert multiple related `card` rows.
- Update Deck: Update deck info, update/add/remove/reorder related cards.

These operations touch multiple tables/entities. Without transactional guarantees, partial failures (e.g., deck updated but cards not updated) would lead to inconsistent state.

## Decision
Wrap **deck create and update operations in a database transaction**.

- Use the *Unit of Work pattern* to gather all changes (deck and related cards) and commit in one transaction.
- On any failure during the process, the transaction will rollback, leaving the system unchanged.
- For simple single-entity updates (e.g., renaming deck title only), a transaction is optional since the database ensures atomicity of single statements.

## Alternatives Considered
- **No transaction**:  
  - Simpler, but risks leaving orphan records or inconsistent data if one step fails.  
- **Saga / Event Sourcing**:  
  - Suitable for distributed or multi-service systems, but adds complexity.  
  - Not required in our current monolithic + single-database context.

## Consequences
- ✅ Ensures data consistency for complex create/update flows.  
- ✅ Easier reasoning about system state (all-or-nothing).  
- ⚠️ Slight overhead from managing transactions.  
- ⚠️ Must ensure repository and service layer code respect transaction boundaries.

## Status
Will be applied to all create/update deck use cases involving multiple entities.

## References
