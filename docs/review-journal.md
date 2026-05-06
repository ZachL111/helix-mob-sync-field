# Review Journal

I treated `helix-mob-sync-field` as a project where the smallest useful behavior should still be inspectable.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its mobile workflows focus without claiming live deployment or external usage.

## Cases

- `baseline`: `form pressure`, score 243, lane `ship`
- `stress`: `sync drift`, score 152, lane `ship`
- `edge`: `local state`, score 165, lane `ship`
- `recovery`: `conflict cost`, score 114, lane `watch`
- `stale`: `form pressure`, score 245, lane `ship`

## Note

This file is intentionally plain so the fixture remains the source of truth.
