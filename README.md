# contracts-metadata

Protobuf/gRPC contract interfaces for metadata provider modules in MuxCore.

This repo is the canonical home for the **MetadataProvider** contract referenced by metadata modules (e.g. `metadata-tmdb` `muxcore.json`).

## Proto Service

- **`MetadataService`** (`proto/muxcore/metadata/v1/metadata.proto`) — logical interface name **`MetadataProvider`**
  - `Search` — search movies and TV by query
  - `GetMovieDetails` / `GetTVDetails` / `GetSeasonDetails` / `GetEpisodeDetails` — fetch metadata
  - `GetCollection` — fetch movie collection details
  - `GetConfiguration` — image base URLs and size presets
  - `ListTrending` / `ListPopular` / `ListSimilar` / `ListRecommendations` — discovery lists
  - `FindByExternalID` — resolve external IDs (e.g. IMDb, TVDB)
  - `GetAlternativeTitles` — localized/alternate titles

Contract repo capability: **`contracts.metadata`**. Implementers advertise capability **`metadata`**.

### Request notes

- Detail/season/episode/collection requests use provider-neutral **`id`** plus optional **`external_ids`** (`tmdb_id`, `tvdb_id`, `imdb_id`).
- `SearchRequest.include_adult` defaults false when unset; `region` carries ISO 3166-1 localization.
- `GetTVDetailsResponse.in_production` is a **bool** (field 30); field 18 is reserved.

## Version

**v0.2.0** — see [CHANGELOG.md](CHANGELOG.md) and [COMPATIBILITY.md](COMPATIBILITY.md).

## Go import

Generated stubs: `github.com/Muxcore-Media/contracts-metadata/muxcore/metadata/v1`

Stubs are checked in under `muxcore/metadata/v1/`. Regenerate after proto edits only:

```bash
make proto
go test ./...
```

CI runs `make proto` and fails if generated stubs drift.

## Implementing Modules

Any module advertising capability **`metadata`** may implement this service.

### Known implementers

- [metadata-tmdb](https://github.com/Muxcore-Media/metadata-tmdb) — TMDB API provider

### Known consumers

- [metadata-tmdb](https://github.com/Muxcore-Media/metadata-tmdb) — primary implementer
- [media-movies](https://github.com/Muxcore-Media/media-movies) — movie library metadata refresh
- [media-tvshows](https://github.com/Muxcore-Media/media-tvshows) — TV library metadata refresh
- [request-media](https://github.com/Muxcore-Media/request-media) — request/discover flows
- [media-list-sync](https://github.com/Muxcore-Media/media-list-sync) — list import/migration
- `_mvp/cmd/mediauiprox` — consumer UI discover proxy
