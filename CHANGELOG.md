# Changelog

## [Unreleased]

### Added
- `certification` and `certification_country` on `GetMovieDetailsResponse` (fields **28**, **29**) and `GetTVDetailsResponse` (fields **33**, **34**): the provider's raw certification for one configured ISO 3166-1 alpha-2 country, empty when none (ADR-0031 §2, roadmap T-M4-01 S4a).
- `contract_test.go` freezes the new field numbers and types.

### Compatibility
- Additive and wire compatible: old readers ignore the fields, old implementers leave them empty.
- metadata-tmdb's own `proto/metadatav1` copy must carry the identical numbers, names and types (media-movies decodes that copy from the same server).

## [0.2.0] - 2026-10-05

### Changed
- CI runs on GitHub-hosted runners from the umbrella template; retired-origin workflows removed.
- Dependencies resolve from published GitHub tags (no filesystem `replace`); requires core v0.6.0.

## [0.2.0] — 2026-08-31

### Added
- `ExternalIds` message and provider-neutral `id` on detail/collection/alternative-title requests (replaces TMDB-only `tmdb_id`).
- `imdb_id` and `tvdb_id` on `SearchResult` and `GetTVDetailsResponse`.
- `GetEpisodeDetails` RPC and `still_url` on `Episode`.
- `ListSimilar` and `ListRecommendations` RPCs (media type + id + page).
- `SearchRequest.include_adult` and `SearchRequest.region`.
- CI `proto-check` job.
- `contract_test.go` freezing RPC names and field numbers.

### Changed
- `GetTVDetailsResponse.in_production`: reserved field 18 (was string); new `bool in_production` at tag 30.
- `Makefile` `clean` clears Go cache only — does not delete published stubs.

### Compatibility
- **Breaking** request-field rename (`tmdb_id` → `id`) and `in_production` wire-type change require coordinated consumer/implementer bumps.
- Additive **v1** fields otherwise; implementers may adopt incrementally.

## [0.1.0] — 2026-08-08

### Added
- Initial `MetadataService` contract (`Search`, `GetMovieDetails`, `GetTVDetails`, `GetSeasonDetails`, `GetCollection`, `GetConfiguration`, `ListTrending`, `ListPopular`, `FindByExternalID`, `GetAlternativeTitles`).
- Generated Go stubs under `muxcore/metadata/v1`.

### Compatibility
- Contract interface version **v1**; logical name **MetadataProvider**; implementers advertise capability **`metadata`**.
- Consumers: `metadata-tmdb`, `media-movies`, `media-tvshows`, `request-media`, `media-list-sync`, `_mvp/cmd/mediauiprox`.
