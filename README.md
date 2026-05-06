# helix-mob-sync-field

`helix-mob-sync-field` is a Go project in mobile workflows. Its focus is to create a Go reference implementation for sync workflows, centered on policy evaluation, deny and allow fixtures, and explainable decision traces.

## Why I Keep It Small

This is intentionally local and self-contained so it can be inspected without credentials, services, or seeded history.

## Helix Mob Sync Field Review Notes

The first comparison I would make is `form pressure` against `conflict cost` because it shows where the rule is most opinionated.

## Included Behavior

- `fixtures/domain_review.csv` adds cases for form pressure and sync drift.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/helix-mob-sync-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `form pressure` and `conflict cost`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## Internal Model

The core code exposes a scoring path and the added review layer uses `signal`, `slack`, `drag`, and `confidence`. The domain terms are `form pressure`, `sync drift`, `local state`, and `conflict cost`.

The Go code keeps the review rule close to the tests.

## Try It Locally

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Validation

The verifier is intentionally local. It should fail if the fixture score math, lane assignment, or language-specific test drifts.

## Scope

This remains a local project with deterministic fixtures. It does not depend on credentials, hosted services, or live data. Future work should add richer malformed inputs before widening the public API.
