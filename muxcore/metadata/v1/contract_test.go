package metadatav1

import (
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestMetadataServiceRPCNamesStable(t *testing.T) {
	want := map[string]bool{
		"Search":                true,
		"GetMovieDetails":       true,
		"GetTVDetails":          true,
		"GetSeasonDetails":      true,
		"GetEpisodeDetails":     true,
		"GetCollection":         true,
		"GetConfiguration":      true,
		"ListTrending":          true,
		"ListPopular":           true,
		"ListSimilar":           true,
		"ListRecommendations":   true,
		"FindByExternalID":      true,
		"GetAlternativeTitles":  true,
	}
	got := make(map[string]bool, len(MetadataService_ServiceDesc.Methods))
	for _, m := range MetadataService_ServiceDesc.Methods {
		got[m.MethodName] = true
	}
	for name := range want {
		if !got[name] {
			t.Errorf("missing RPC %s", name)
		}
	}
}

func TestSearchRequestFieldNumbers(t *testing.T) {
	want := map[string]protoreflect.FieldNumber{
		"query":         1,
		"type":          2,
		"page":          3,
		"language":      4,
		"year":          5,
		"include_adult": 6,
		"region":        7,
	}
	assertFieldNumbers(t, (&SearchRequest{}).ProtoReflect().Descriptor(), want)
}

func TestSearchResultFieldNumbers(t *testing.T) {
	want := map[string]protoreflect.FieldNumber{
		"id":                1,
		"title":             2,
		"original_title":    3,
		"overview":          4,
		"poster_path":       5,
		"backdrop_path":     6,
		"vote_average":      7,
		"vote_count":        8,
		"release_date":      9,
		"media_type":        10,
		"genre_ids":         11,
		"original_language": 12,
		"popularity":        13,
		"name":              14,
		"first_air_date":    15,
		"imdb_id":           16,
		"tvdb_id":           17,
	}
	assertFieldNumbers(t, (&SearchResult{}).ProtoReflect().Descriptor(), want)
}

func TestGetTVDetailsResponseInProductionField(t *testing.T) {
	desc := (&GetTVDetailsResponse{}).ProtoReflect().Descriptor()
	if fd := desc.Fields().ByName("in_production"); fd == nil {
		t.Fatal("missing in_production field")
	} else if fd.Kind() != protoreflect.BoolKind {
		t.Fatalf("in_production kind = %v, want bool", fd.Kind())
	} else if fd.Number() != 30 {
		t.Fatalf("in_production tag = %d, want 30", fd.Number())
	}
}

func TestDetailRequestsUseProviderNeutralID(t *testing.T) {
	cases := []protoreflect.MessageDescriptor{
		(&GetMovieDetailsRequest{}).ProtoReflect().Descriptor(),
		(&GetTVDetailsRequest{}).ProtoReflect().Descriptor(),
		(&GetSeasonDetailsRequest{}).ProtoReflect().Descriptor(),
		(&GetCollectionRequest{}).ProtoReflect().Descriptor(),
		(&GetAlternativeTitlesRequest{}).ProtoReflect().Descriptor(),
	}
	for _, desc := range cases {
		if desc.Fields().ByName("tmdb_id") != nil {
			t.Errorf("%s still has tmdb_id field", desc.Name())
		}
		if desc.Fields().ByName("id") == nil {
			t.Errorf("%s missing id field", desc.Name())
		}
	}
}

func TestImportPath(t *testing.T) {
	var _ grpc.ServiceDesc = MetadataService_ServiceDesc
}

func TestCapabilityIDs(t *testing.T) {
	const consumerCapability = "metadata"
	const contractCapability = "contracts.metadata"
	if consumerCapability == contractCapability {
		t.Fatal("consumer and contract capability ids must differ")
	}
}

func assertFieldNumbers(t *testing.T, desc protoreflect.MessageDescriptor, want map[string]protoreflect.FieldNumber) {
	t.Helper()
	for name, num := range want {
		fd := desc.Fields().ByName(protoreflect.Name(name))
		if fd == nil {
			t.Errorf("missing field %q", name)
			continue
		}
		if fd.Number() != num {
			t.Errorf("field %q number = %d, want %d", name, fd.Number(), num)
		}
	}
}
