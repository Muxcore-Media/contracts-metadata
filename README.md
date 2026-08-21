# contracts-metadata

Protobuf/gRPC contract interfaces for metadata provider modules in MuxCore.

This repo is the canonical home for the **MetadataProvider** contract referenced by metadata modules (e.g. `metadata-tmdb` `muxcore.json`).

## Proto Service

- **`MetadataService`** (`proto/muxcore/metadata/v1/metadata.proto`)
  - `Search` — search movies and TV by query
  - `GetMovieDetails` / `GetTVDetails` / `GetSeasonDetails` — fetch TMDB-style metadata
  - `GetCollection` — fetch movie collection details
  - `GetConfiguration` — image base URLs and size presets
  - `ListTrending` / `ListPopular` — discovery lists
  - `FindByExternalID` — resolve external IDs (e.g. IMDb)
  - `GetAlternativeTitles` — localized/alternate titles

Logical contract name: **MetadataProvider** v0.1.0.

## Go import

Generated stubs (after `make proto`): `github.com/Muxcore-Media/contracts-metadata/muxcore/metadata/v1`

Proto source was migrated from `metadata-tmdb/proto/metadatav1/`. Generated Go is not checked in yet; run `make proto` when `protoc` and the Go plugins are available.

## Implementing Modules

Any module advertising capability `metadata` may implement this service.

### Known implementers

- [metadata-tmdb](https://github.com/Muxcore-Media/metadata-tmdb) — TMDB API provider
