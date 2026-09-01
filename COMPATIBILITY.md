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

## Discovery beyond trending/popular

`ListSimilar` and `ListRecommendations` accept media type, provider `id`, and paging. Implementers may return `Unimplemented` until supported; callers should degrade gracefully.

Breaking proto changes require a new major contract version and coordinated consumer updates.
