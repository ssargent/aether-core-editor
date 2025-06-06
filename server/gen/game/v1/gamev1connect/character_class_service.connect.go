// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: game/v1/character_class_service.proto

package gamev1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/ssargent/aether-core-editor/gen/game/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// CharacterClassServiceName is the fully-qualified name of the CharacterClassService service.
	CharacterClassServiceName = "game.v1.CharacterClassService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CharacterClassServiceCreateCharacterClassProcedure is the fully-qualified name of the
	// CharacterClassService's CreateCharacterClass RPC.
	CharacterClassServiceCreateCharacterClassProcedure = "/game.v1.CharacterClassService/CreateCharacterClass"
	// CharacterClassServiceGetCharacterClassProcedure is the fully-qualified name of the
	// CharacterClassService's GetCharacterClass RPC.
	CharacterClassServiceGetCharacterClassProcedure = "/game.v1.CharacterClassService/GetCharacterClass"
	// CharacterClassServiceUpdateCharacterClassProcedure is the fully-qualified name of the
	// CharacterClassService's UpdateCharacterClass RPC.
	CharacterClassServiceUpdateCharacterClassProcedure = "/game.v1.CharacterClassService/UpdateCharacterClass"
	// CharacterClassServiceDeleteCharacterClassProcedure is the fully-qualified name of the
	// CharacterClassService's DeleteCharacterClass RPC.
	CharacterClassServiceDeleteCharacterClassProcedure = "/game.v1.CharacterClassService/DeleteCharacterClass"
	// CharacterClassServiceListCharacterClassesProcedure is the fully-qualified name of the
	// CharacterClassService's ListCharacterClasses RPC.
	CharacterClassServiceListCharacterClassesProcedure = "/game.v1.CharacterClassService/ListCharacterClasses"
	// CharacterClassServiceCreateCharacterClassFeatureProcedure is the fully-qualified name of the
	// CharacterClassService's CreateCharacterClassFeature RPC.
	CharacterClassServiceCreateCharacterClassFeatureProcedure = "/game.v1.CharacterClassService/CreateCharacterClassFeature"
	// CharacterClassServiceGetCharacterClassFeatureProcedure is the fully-qualified name of the
	// CharacterClassService's GetCharacterClassFeature RPC.
	CharacterClassServiceGetCharacterClassFeatureProcedure = "/game.v1.CharacterClassService/GetCharacterClassFeature"
	// CharacterClassServiceUpdateCharacterClassFeatureProcedure is the fully-qualified name of the
	// CharacterClassService's UpdateCharacterClassFeature RPC.
	CharacterClassServiceUpdateCharacterClassFeatureProcedure = "/game.v1.CharacterClassService/UpdateCharacterClassFeature"
	// CharacterClassServiceDeleteCharacterClassFeatureProcedure is the fully-qualified name of the
	// CharacterClassService's DeleteCharacterClassFeature RPC.
	CharacterClassServiceDeleteCharacterClassFeatureProcedure = "/game.v1.CharacterClassService/DeleteCharacterClassFeature"
	// CharacterClassServiceListCharacterClassFeaturesProcedure is the fully-qualified name of the
	// CharacterClassService's ListCharacterClassFeatures RPC.
	CharacterClassServiceListCharacterClassFeaturesProcedure = "/game.v1.CharacterClassService/ListCharacterClassFeatures"
)

// CharacterClassServiceClient is a client for the game.v1.CharacterClassService service.
type CharacterClassServiceClient interface {
	CreateCharacterClass(context.Context, *connect.Request[v1.CreateCharacterClassRequest]) (*connect.Response[v1.CharacterClassResponse], error)
	GetCharacterClass(context.Context, *connect.Request[v1.GetCharacterClassRequest]) (*connect.Response[v1.CharacterClassResponse], error)
	UpdateCharacterClass(context.Context, *connect.Request[v1.UpdateCharacterClassRequest]) (*connect.Response[v1.CharacterClassResponse], error)
	DeleteCharacterClass(context.Context, *connect.Request[v1.DeleteCharacterClassRequest]) (*connect.Response[v1.DeleteResponse], error)
	ListCharacterClasses(context.Context, *connect.Request[v1.ListCharacterClassesRequest]) (*connect.Response[v1.ListCharacterClassesResponse], error)
	CreateCharacterClassFeature(context.Context, *connect.Request[v1.CreateCharacterClassFeatureRequest]) (*connect.Response[v1.CharacterClassFeatureResponse], error)
	GetCharacterClassFeature(context.Context, *connect.Request[v1.GetCharacterClassFeatureRequest]) (*connect.Response[v1.CharacterClassFeatureResponse], error)
	UpdateCharacterClassFeature(context.Context, *connect.Request[v1.UpdateCharacterClassFeatureRequest]) (*connect.Response[v1.CharacterClassFeatureResponse], error)
	DeleteCharacterClassFeature(context.Context, *connect.Request[v1.DeleteCharacterClassFeatureRequest]) (*connect.Response[v1.DeleteResponse], error)
	ListCharacterClassFeatures(context.Context, *connect.Request[v1.ListCharacterClassFeaturesRequest]) (*connect.Response[v1.ListCharacterClassFeaturesResponse], error)
}

// NewCharacterClassServiceClient constructs a client for the game.v1.CharacterClassService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCharacterClassServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) CharacterClassServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	characterClassServiceMethods := v1.File_game_v1_character_class_service_proto.Services().ByName("CharacterClassService").Methods()
	return &characterClassServiceClient{
		createCharacterClass: connect.NewClient[v1.CreateCharacterClassRequest, v1.CharacterClassResponse](
			httpClient,
			baseURL+CharacterClassServiceCreateCharacterClassProcedure,
			connect.WithSchema(characterClassServiceMethods.ByName("CreateCharacterClass")),
			connect.WithClientOptions(opts...),
		),
		getCharacterClass: connect.NewClient[v1.GetCharacterClassRequest, v1.CharacterClassResponse](
			httpClient,
			baseURL+CharacterClassServiceGetCharacterClassProcedure,
			connect.WithSchema(characterClassServiceMethods.ByName("GetCharacterClass")),
			connect.WithClientOptions(opts...),
		),
		updateCharacterClass: connect.NewClient[v1.UpdateCharacterClassRequest, v1.CharacterClassResponse](
			httpClient,
			baseURL+CharacterClassServiceUpdateCharacterClassProcedure,
			connect.WithSchema(characterClassServiceMethods.ByName("UpdateCharacterClass")),
			connect.WithClientOptions(opts...),
		),
		deleteCharacterClass: connect.NewClient[v1.DeleteCharacterClassRequest, v1.DeleteResponse](
			httpClient,
			baseURL+CharacterClassServiceDeleteCharacterClassProcedure,
			connect.WithSchema(characterClassServiceMethods.ByName("DeleteCharacterClass")),
			connect.WithClientOptions(opts...),
		),
		listCharacterClasses: connect.NewClient[v1.ListCharacterClassesRequest, v1.ListCharacterClassesResponse](
			httpClient,
			baseURL+CharacterClassServiceListCharacterClassesProcedure,
			connect.WithSchema(characterClassServiceMethods.ByName("ListCharacterClasses")),
			connect.WithClientOptions(opts...),
		),
		createCharacterClassFeature: connect.NewClient[v1.CreateCharacterClassFeatureRequest, v1.CharacterClassFeatureResponse](
			httpClient,
			baseURL+CharacterClassServiceCreateCharacterClassFeatureProcedure,
			connect.WithSchema(characterClassServiceMethods.ByName("CreateCharacterClassFeature")),
			connect.WithClientOptions(opts...),
		),
		getCharacterClassFeature: connect.NewClient[v1.GetCharacterClassFeatureRequest, v1.CharacterClassFeatureResponse](
			httpClient,
			baseURL+CharacterClassServiceGetCharacterClassFeatureProcedure,
			connect.WithSchema(characterClassServiceMethods.ByName("GetCharacterClassFeature")),
			connect.WithClientOptions(opts...),
		),
		updateCharacterClassFeature: connect.NewClient[v1.UpdateCharacterClassFeatureRequest, v1.CharacterClassFeatureResponse](
			httpClient,
			baseURL+CharacterClassServiceUpdateCharacterClassFeatureProcedure,
			connect.WithSchema(characterClassServiceMethods.ByName("UpdateCharacterClassFeature")),
			connect.WithClientOptions(opts...),
		),
		deleteCharacterClassFeature: connect.NewClient[v1.DeleteCharacterClassFeatureRequest, v1.DeleteResponse](
			httpClient,
			baseURL+CharacterClassServiceDeleteCharacterClassFeatureProcedure,
			connect.WithSchema(characterClassServiceMethods.ByName("DeleteCharacterClassFeature")),
			connect.WithClientOptions(opts...),
		),
		listCharacterClassFeatures: connect.NewClient[v1.ListCharacterClassFeaturesRequest, v1.ListCharacterClassFeaturesResponse](
			httpClient,
			baseURL+CharacterClassServiceListCharacterClassFeaturesProcedure,
			connect.WithSchema(characterClassServiceMethods.ByName("ListCharacterClassFeatures")),
			connect.WithClientOptions(opts...),
		),
	}
}

// characterClassServiceClient implements CharacterClassServiceClient.
type characterClassServiceClient struct {
	createCharacterClass        *connect.Client[v1.CreateCharacterClassRequest, v1.CharacterClassResponse]
	getCharacterClass           *connect.Client[v1.GetCharacterClassRequest, v1.CharacterClassResponse]
	updateCharacterClass        *connect.Client[v1.UpdateCharacterClassRequest, v1.CharacterClassResponse]
	deleteCharacterClass        *connect.Client[v1.DeleteCharacterClassRequest, v1.DeleteResponse]
	listCharacterClasses        *connect.Client[v1.ListCharacterClassesRequest, v1.ListCharacterClassesResponse]
	createCharacterClassFeature *connect.Client[v1.CreateCharacterClassFeatureRequest, v1.CharacterClassFeatureResponse]
	getCharacterClassFeature    *connect.Client[v1.GetCharacterClassFeatureRequest, v1.CharacterClassFeatureResponse]
	updateCharacterClassFeature *connect.Client[v1.UpdateCharacterClassFeatureRequest, v1.CharacterClassFeatureResponse]
	deleteCharacterClassFeature *connect.Client[v1.DeleteCharacterClassFeatureRequest, v1.DeleteResponse]
	listCharacterClassFeatures  *connect.Client[v1.ListCharacterClassFeaturesRequest, v1.ListCharacterClassFeaturesResponse]
}

// CreateCharacterClass calls game.v1.CharacterClassService.CreateCharacterClass.
func (c *characterClassServiceClient) CreateCharacterClass(ctx context.Context, req *connect.Request[v1.CreateCharacterClassRequest]) (*connect.Response[v1.CharacterClassResponse], error) {
	return c.createCharacterClass.CallUnary(ctx, req)
}

// GetCharacterClass calls game.v1.CharacterClassService.GetCharacterClass.
func (c *characterClassServiceClient) GetCharacterClass(ctx context.Context, req *connect.Request[v1.GetCharacterClassRequest]) (*connect.Response[v1.CharacterClassResponse], error) {
	return c.getCharacterClass.CallUnary(ctx, req)
}

// UpdateCharacterClass calls game.v1.CharacterClassService.UpdateCharacterClass.
func (c *characterClassServiceClient) UpdateCharacterClass(ctx context.Context, req *connect.Request[v1.UpdateCharacterClassRequest]) (*connect.Response[v1.CharacterClassResponse], error) {
	return c.updateCharacterClass.CallUnary(ctx, req)
}

// DeleteCharacterClass calls game.v1.CharacterClassService.DeleteCharacterClass.
func (c *characterClassServiceClient) DeleteCharacterClass(ctx context.Context, req *connect.Request[v1.DeleteCharacterClassRequest]) (*connect.Response[v1.DeleteResponse], error) {
	return c.deleteCharacterClass.CallUnary(ctx, req)
}

// ListCharacterClasses calls game.v1.CharacterClassService.ListCharacterClasses.
func (c *characterClassServiceClient) ListCharacterClasses(ctx context.Context, req *connect.Request[v1.ListCharacterClassesRequest]) (*connect.Response[v1.ListCharacterClassesResponse], error) {
	return c.listCharacterClasses.CallUnary(ctx, req)
}

// CreateCharacterClassFeature calls game.v1.CharacterClassService.CreateCharacterClassFeature.
func (c *characterClassServiceClient) CreateCharacterClassFeature(ctx context.Context, req *connect.Request[v1.CreateCharacterClassFeatureRequest]) (*connect.Response[v1.CharacterClassFeatureResponse], error) {
	return c.createCharacterClassFeature.CallUnary(ctx, req)
}

// GetCharacterClassFeature calls game.v1.CharacterClassService.GetCharacterClassFeature.
func (c *characterClassServiceClient) GetCharacterClassFeature(ctx context.Context, req *connect.Request[v1.GetCharacterClassFeatureRequest]) (*connect.Response[v1.CharacterClassFeatureResponse], error) {
	return c.getCharacterClassFeature.CallUnary(ctx, req)
}

// UpdateCharacterClassFeature calls game.v1.CharacterClassService.UpdateCharacterClassFeature.
func (c *characterClassServiceClient) UpdateCharacterClassFeature(ctx context.Context, req *connect.Request[v1.UpdateCharacterClassFeatureRequest]) (*connect.Response[v1.CharacterClassFeatureResponse], error) {
	return c.updateCharacterClassFeature.CallUnary(ctx, req)
}

// DeleteCharacterClassFeature calls game.v1.CharacterClassService.DeleteCharacterClassFeature.
func (c *characterClassServiceClient) DeleteCharacterClassFeature(ctx context.Context, req *connect.Request[v1.DeleteCharacterClassFeatureRequest]) (*connect.Response[v1.DeleteResponse], error) {
	return c.deleteCharacterClassFeature.CallUnary(ctx, req)
}

// ListCharacterClassFeatures calls game.v1.CharacterClassService.ListCharacterClassFeatures.
func (c *characterClassServiceClient) ListCharacterClassFeatures(ctx context.Context, req *connect.Request[v1.ListCharacterClassFeaturesRequest]) (*connect.Response[v1.ListCharacterClassFeaturesResponse], error) {
	return c.listCharacterClassFeatures.CallUnary(ctx, req)
}

// CharacterClassServiceHandler is an implementation of the game.v1.CharacterClassService service.
type CharacterClassServiceHandler interface {
	CreateCharacterClass(context.Context, *connect.Request[v1.CreateCharacterClassRequest]) (*connect.Response[v1.CharacterClassResponse], error)
	GetCharacterClass(context.Context, *connect.Request[v1.GetCharacterClassRequest]) (*connect.Response[v1.CharacterClassResponse], error)
	UpdateCharacterClass(context.Context, *connect.Request[v1.UpdateCharacterClassRequest]) (*connect.Response[v1.CharacterClassResponse], error)
	DeleteCharacterClass(context.Context, *connect.Request[v1.DeleteCharacterClassRequest]) (*connect.Response[v1.DeleteResponse], error)
	ListCharacterClasses(context.Context, *connect.Request[v1.ListCharacterClassesRequest]) (*connect.Response[v1.ListCharacterClassesResponse], error)
	CreateCharacterClassFeature(context.Context, *connect.Request[v1.CreateCharacterClassFeatureRequest]) (*connect.Response[v1.CharacterClassFeatureResponse], error)
	GetCharacterClassFeature(context.Context, *connect.Request[v1.GetCharacterClassFeatureRequest]) (*connect.Response[v1.CharacterClassFeatureResponse], error)
	UpdateCharacterClassFeature(context.Context, *connect.Request[v1.UpdateCharacterClassFeatureRequest]) (*connect.Response[v1.CharacterClassFeatureResponse], error)
	DeleteCharacterClassFeature(context.Context, *connect.Request[v1.DeleteCharacterClassFeatureRequest]) (*connect.Response[v1.DeleteResponse], error)
	ListCharacterClassFeatures(context.Context, *connect.Request[v1.ListCharacterClassFeaturesRequest]) (*connect.Response[v1.ListCharacterClassFeaturesResponse], error)
}

// NewCharacterClassServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCharacterClassServiceHandler(svc CharacterClassServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	characterClassServiceMethods := v1.File_game_v1_character_class_service_proto.Services().ByName("CharacterClassService").Methods()
	characterClassServiceCreateCharacterClassHandler := connect.NewUnaryHandler(
		CharacterClassServiceCreateCharacterClassProcedure,
		svc.CreateCharacterClass,
		connect.WithSchema(characterClassServiceMethods.ByName("CreateCharacterClass")),
		connect.WithHandlerOptions(opts...),
	)
	characterClassServiceGetCharacterClassHandler := connect.NewUnaryHandler(
		CharacterClassServiceGetCharacterClassProcedure,
		svc.GetCharacterClass,
		connect.WithSchema(characterClassServiceMethods.ByName("GetCharacterClass")),
		connect.WithHandlerOptions(opts...),
	)
	characterClassServiceUpdateCharacterClassHandler := connect.NewUnaryHandler(
		CharacterClassServiceUpdateCharacterClassProcedure,
		svc.UpdateCharacterClass,
		connect.WithSchema(characterClassServiceMethods.ByName("UpdateCharacterClass")),
		connect.WithHandlerOptions(opts...),
	)
	characterClassServiceDeleteCharacterClassHandler := connect.NewUnaryHandler(
		CharacterClassServiceDeleteCharacterClassProcedure,
		svc.DeleteCharacterClass,
		connect.WithSchema(characterClassServiceMethods.ByName("DeleteCharacterClass")),
		connect.WithHandlerOptions(opts...),
	)
	characterClassServiceListCharacterClassesHandler := connect.NewUnaryHandler(
		CharacterClassServiceListCharacterClassesProcedure,
		svc.ListCharacterClasses,
		connect.WithSchema(characterClassServiceMethods.ByName("ListCharacterClasses")),
		connect.WithHandlerOptions(opts...),
	)
	characterClassServiceCreateCharacterClassFeatureHandler := connect.NewUnaryHandler(
		CharacterClassServiceCreateCharacterClassFeatureProcedure,
		svc.CreateCharacterClassFeature,
		connect.WithSchema(characterClassServiceMethods.ByName("CreateCharacterClassFeature")),
		connect.WithHandlerOptions(opts...),
	)
	characterClassServiceGetCharacterClassFeatureHandler := connect.NewUnaryHandler(
		CharacterClassServiceGetCharacterClassFeatureProcedure,
		svc.GetCharacterClassFeature,
		connect.WithSchema(characterClassServiceMethods.ByName("GetCharacterClassFeature")),
		connect.WithHandlerOptions(opts...),
	)
	characterClassServiceUpdateCharacterClassFeatureHandler := connect.NewUnaryHandler(
		CharacterClassServiceUpdateCharacterClassFeatureProcedure,
		svc.UpdateCharacterClassFeature,
		connect.WithSchema(characterClassServiceMethods.ByName("UpdateCharacterClassFeature")),
		connect.WithHandlerOptions(opts...),
	)
	characterClassServiceDeleteCharacterClassFeatureHandler := connect.NewUnaryHandler(
		CharacterClassServiceDeleteCharacterClassFeatureProcedure,
		svc.DeleteCharacterClassFeature,
		connect.WithSchema(characterClassServiceMethods.ByName("DeleteCharacterClassFeature")),
		connect.WithHandlerOptions(opts...),
	)
	characterClassServiceListCharacterClassFeaturesHandler := connect.NewUnaryHandler(
		CharacterClassServiceListCharacterClassFeaturesProcedure,
		svc.ListCharacterClassFeatures,
		connect.WithSchema(characterClassServiceMethods.ByName("ListCharacterClassFeatures")),
		connect.WithHandlerOptions(opts...),
	)
	return "/game.v1.CharacterClassService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case CharacterClassServiceCreateCharacterClassProcedure:
			characterClassServiceCreateCharacterClassHandler.ServeHTTP(w, r)
		case CharacterClassServiceGetCharacterClassProcedure:
			characterClassServiceGetCharacterClassHandler.ServeHTTP(w, r)
		case CharacterClassServiceUpdateCharacterClassProcedure:
			characterClassServiceUpdateCharacterClassHandler.ServeHTTP(w, r)
		case CharacterClassServiceDeleteCharacterClassProcedure:
			characterClassServiceDeleteCharacterClassHandler.ServeHTTP(w, r)
		case CharacterClassServiceListCharacterClassesProcedure:
			characterClassServiceListCharacterClassesHandler.ServeHTTP(w, r)
		case CharacterClassServiceCreateCharacterClassFeatureProcedure:
			characterClassServiceCreateCharacterClassFeatureHandler.ServeHTTP(w, r)
		case CharacterClassServiceGetCharacterClassFeatureProcedure:
			characterClassServiceGetCharacterClassFeatureHandler.ServeHTTP(w, r)
		case CharacterClassServiceUpdateCharacterClassFeatureProcedure:
			characterClassServiceUpdateCharacterClassFeatureHandler.ServeHTTP(w, r)
		case CharacterClassServiceDeleteCharacterClassFeatureProcedure:
			characterClassServiceDeleteCharacterClassFeatureHandler.ServeHTTP(w, r)
		case CharacterClassServiceListCharacterClassFeaturesProcedure:
			characterClassServiceListCharacterClassFeaturesHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedCharacterClassServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCharacterClassServiceHandler struct{}

func (UnimplementedCharacterClassServiceHandler) CreateCharacterClass(context.Context, *connect.Request[v1.CreateCharacterClassRequest]) (*connect.Response[v1.CharacterClassResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("game.v1.CharacterClassService.CreateCharacterClass is not implemented"))
}

func (UnimplementedCharacterClassServiceHandler) GetCharacterClass(context.Context, *connect.Request[v1.GetCharacterClassRequest]) (*connect.Response[v1.CharacterClassResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("game.v1.CharacterClassService.GetCharacterClass is not implemented"))
}

func (UnimplementedCharacterClassServiceHandler) UpdateCharacterClass(context.Context, *connect.Request[v1.UpdateCharacterClassRequest]) (*connect.Response[v1.CharacterClassResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("game.v1.CharacterClassService.UpdateCharacterClass is not implemented"))
}

func (UnimplementedCharacterClassServiceHandler) DeleteCharacterClass(context.Context, *connect.Request[v1.DeleteCharacterClassRequest]) (*connect.Response[v1.DeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("game.v1.CharacterClassService.DeleteCharacterClass is not implemented"))
}

func (UnimplementedCharacterClassServiceHandler) ListCharacterClasses(context.Context, *connect.Request[v1.ListCharacterClassesRequest]) (*connect.Response[v1.ListCharacterClassesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("game.v1.CharacterClassService.ListCharacterClasses is not implemented"))
}

func (UnimplementedCharacterClassServiceHandler) CreateCharacterClassFeature(context.Context, *connect.Request[v1.CreateCharacterClassFeatureRequest]) (*connect.Response[v1.CharacterClassFeatureResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("game.v1.CharacterClassService.CreateCharacterClassFeature is not implemented"))
}

func (UnimplementedCharacterClassServiceHandler) GetCharacterClassFeature(context.Context, *connect.Request[v1.GetCharacterClassFeatureRequest]) (*connect.Response[v1.CharacterClassFeatureResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("game.v1.CharacterClassService.GetCharacterClassFeature is not implemented"))
}

func (UnimplementedCharacterClassServiceHandler) UpdateCharacterClassFeature(context.Context, *connect.Request[v1.UpdateCharacterClassFeatureRequest]) (*connect.Response[v1.CharacterClassFeatureResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("game.v1.CharacterClassService.UpdateCharacterClassFeature is not implemented"))
}

func (UnimplementedCharacterClassServiceHandler) DeleteCharacterClassFeature(context.Context, *connect.Request[v1.DeleteCharacterClassFeatureRequest]) (*connect.Response[v1.DeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("game.v1.CharacterClassService.DeleteCharacterClassFeature is not implemented"))
}

func (UnimplementedCharacterClassServiceHandler) ListCharacterClassFeatures(context.Context, *connect.Request[v1.ListCharacterClassFeaturesRequest]) (*connect.Response[v1.ListCharacterClassFeaturesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("game.v1.CharacterClassService.ListCharacterClassFeatures is not implemented"))
}
