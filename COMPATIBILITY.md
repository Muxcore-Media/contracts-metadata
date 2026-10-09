# Compatibility

| Component | Requirement |
|-----------|-------------|
| Contract interface | `MetadataService` **v1** (logical name **MetadataProvider**) |
| Go module path | `github.com/Muxcore-Media/contracts-metadata` |
| Generated package | `github.com/Muxcore-Media/contracts-metadata/muxcore/metadata/v1` |
| MuxCore core | ≥ 0.4.0 (consumers) |
| Capability ID (contract repo) | `contracts.metadata` |
| Capability ID (implementers / callers) | `metadata` |

Implementers register capability **`metadata`** in `muxcore.json` and implement gRPC service **`MetadataService`**. The contract repository capability id **`contracts.metadata`** identifies this proto module in the umbrella catalog.

## Provider-neutral IDs

Detail, season, episode, collection, similar, and recommendation requests use **`id`** (provider-native integer) plus optional **`external_ids`** (`tmdb_id`, `tvdb_id`, `imdb_id`). Callers must not assume `id` is always a TMDB id — only that the active metadata provider understands it.

Implementers should honor `external_ids` when `id` is zero and the provider supports cross-ID lookup.

## Search localization

| Field | Rule |
|-------|------|
| `SearchRequest.include_adult` | Default **false** when unset; omit adult titles from results |
| `SearchRequest.region` | ISO 3166-1 region code for localized results (e.g. `US`) |

## TV airing status

`GetTVDetailsResponse.in_production` is a **bool** at field tag **30**. Field **18** is reserved (previously an incorrectly typed string). Callers must not parse string values for airing status.

## External IDs on results

`SearchResult` and `GetTVDetailsResponse` expose **`imdb_id`** and **`tvdb_id`** for Arr migrate, list-sync, subtitles, and playback correlation. Movie details already carried `imdb_id`.

## Episode stills

`Episode.still_url` is the fully qualified still image URL (parity with movie `poster_url`). Prefer `GetEpisodeDetails` over fetching an entire season when refreshing one episode.

## Content certification (ADR-0031 §2)

| Message | `certification` | `certification_country` |
|---------|-----------------|-------------------------|
| `GetMovieDetailsResponse` | tag **28** (string) | tag **29** (string) |
| `GetTVDetailsResponse` | tag **33** (string) | tag **34** (string) |

- `certification` is the provider's **raw** value for one configured ISO 3166-1 alpha-2 country (e.g. `PG-13`, `TV-MA`, `NR`, or a non-US token such as `15`). Implementers trim it, strip control characters and return it only if it is at most 16 bytes; they do **not** map it onto a rating ladder.
- Both fields are empty when the provider has no certification for that country. Empty is never an error and never means "unrestricted".
- Consumers (media-movies, media-tvshows) own the mapping: an `operator` rating always wins, and an empty or unknown token is *unavailable* (ADR-0031 §2.5).
- The tags are frozen in `contract_test.go`. metadata-tmdb also serves a deprecated copy (`metadata-tmdb/proto/metadatav1`) decoded by media-movies; both copies must keep identical numbers, names and types for every message the server returns.

## Discovery beyond trending/popular

`ListSimilar` and `ListRecommendations` accept media type, provider `id`, and paging. Implementers may return `Unimplemented` until supported; callers should degrade gracefully.

Breaking proto changes require a new major contract version and coordinated consumer updates.
