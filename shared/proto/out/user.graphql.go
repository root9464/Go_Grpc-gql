// Code generated by proroc-gen-graphql, DO NOT EDIT.
package out

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	gql__type_HelloResponse        *graphql.Object      // message HelloResponse in user.proto
	gql__type_HelloRequest         *graphql.Object      // message HelloRequest in user.proto
	gql__type_GetUserInfoResponse  *graphql.Object      // message GetUserInfoResponse in user.proto
	gql__type_GetUserInfoRequest   *graphql.Object      // message GetUserInfoRequest in user.proto
	gql__type_CreateUserResponse   *graphql.Object      // message CreateUserResponse in user.proto
	gql__type_CreateUserRequest    *graphql.Object      // message CreateUserRequest in user.proto
	gql__type_ChangeUserResponse   *graphql.Object      // message ChangeUserResponse in user.proto
	gql__type_ChangeUserRequest    *graphql.Object      // message ChangeUserRequest in user.proto
	gql__input_HelloResponse       *graphql.InputObject // message HelloResponse in user.proto
	gql__input_HelloRequest        *graphql.InputObject // message HelloRequest in user.proto
	gql__input_GetUserInfoResponse *graphql.InputObject // message GetUserInfoResponse in user.proto
	gql__input_GetUserInfoRequest  *graphql.InputObject // message GetUserInfoRequest in user.proto
	gql__input_CreateUserResponse  *graphql.InputObject // message CreateUserResponse in user.proto
	gql__input_CreateUserRequest   *graphql.InputObject // message CreateUserRequest in user.proto
	gql__input_ChangeUserResponse  *graphql.InputObject // message ChangeUserResponse in user.proto
	gql__input_ChangeUserRequest   *graphql.InputObject // message ChangeUserRequest in user.proto
)

func Gql__type_HelloResponse() *graphql.Object {
	if gql__type_HelloResponse == nil {
		gql__type_HelloResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Out_Type_HelloResponse",
			Fields: graphql.Fields{
				"message": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_HelloResponse
}

func Gql__type_HelloRequest() *graphql.Object {
	if gql__type_HelloRequest == nil {
		gql__type_HelloRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Out_Type_HelloRequest",
			Fields: graphql.Fields{
				"name": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		})
	}
	return gql__type_HelloRequest
}

func Gql__type_GetUserInfoResponse() *graphql.Object {
	if gql__type_GetUserInfoResponse == nil {
		gql__type_GetUserInfoResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Out_Type_GetUserInfoResponse",
			Fields: graphql.Fields{
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"password": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__type_GetUserInfoResponse
}

func Gql__type_GetUserInfoRequest() *graphql.Object {
	if gql__type_GetUserInfoRequest == nil {
		gql__type_GetUserInfoRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Out_Type_GetUserInfoRequest",
			Fields: graphql.Fields{
				"name": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		})
	}
	return gql__type_GetUserInfoRequest
}

func Gql__type_CreateUserResponse() *graphql.Object {
	if gql__type_CreateUserResponse == nil {
		gql__type_CreateUserResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Out_Type_CreateUserResponse",
			Fields: graphql.Fields{
				"status": &graphql.Field{
					Type: graphql.Boolean,
				},
			},
		})
	}
	return gql__type_CreateUserResponse
}

func Gql__type_CreateUserRequest() *graphql.Object {
	if gql__type_CreateUserRequest == nil {
		gql__type_CreateUserRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Out_Type_CreateUserRequest",
			Fields: graphql.Fields{
				"name": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"email": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		})
	}
	return gql__type_CreateUserRequest
}

func Gql__type_ChangeUserResponse() *graphql.Object {
	if gql__type_ChangeUserResponse == nil {
		gql__type_ChangeUserResponse = graphql.NewObject(graphql.ObjectConfig{
			Name: "Out_Type_ChangeUserResponse",
			Fields: graphql.Fields{
				"status": &graphql.Field{
					Type: graphql.Boolean,
				},
			},
		})
	}
	return gql__type_ChangeUserResponse
}

func Gql__type_ChangeUserRequest() *graphql.Object {
	if gql__type_ChangeUserRequest == nil {
		gql__type_ChangeUserRequest = graphql.NewObject(graphql.ObjectConfig{
			Name: "Out_Type_ChangeUserRequest",
			Fields: graphql.Fields{
				"name": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"email": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		})
	}
	return gql__type_ChangeUserRequest
}

func Gql__input_HelloResponse() *graphql.InputObject {
	if gql__input_HelloResponse == nil {
		gql__input_HelloResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Out_Input_HelloResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"message": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_HelloResponse
}

func Gql__input_HelloRequest() *graphql.InputObject {
	if gql__input_HelloRequest == nil {
		gql__input_HelloRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Out_Input_HelloRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"name": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		})
	}
	return gql__input_HelloRequest
}

func Gql__input_GetUserInfoResponse() *graphql.InputObject {
	if gql__input_GetUserInfoResponse == nil {
		gql__input_GetUserInfoResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Out_Input_GetUserInfoResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"name": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"password": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"email": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		})
	}
	return gql__input_GetUserInfoResponse
}

func Gql__input_GetUserInfoRequest() *graphql.InputObject {
	if gql__input_GetUserInfoRequest == nil {
		gql__input_GetUserInfoRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Out_Input_GetUserInfoRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"name": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		})
	}
	return gql__input_GetUserInfoRequest
}

func Gql__input_CreateUserResponse() *graphql.InputObject {
	if gql__input_CreateUserResponse == nil {
		gql__input_CreateUserResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Out_Input_CreateUserResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"status": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
			},
		})
	}
	return gql__input_CreateUserResponse
}

func Gql__input_CreateUserRequest() *graphql.InputObject {
	if gql__input_CreateUserRequest == nil {
		gql__input_CreateUserRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Out_Input_CreateUserRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"name": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"email": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		})
	}
	return gql__input_CreateUserRequest
}

func Gql__input_ChangeUserResponse() *graphql.InputObject {
	if gql__input_ChangeUserResponse == nil {
		gql__input_ChangeUserResponse = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Out_Input_ChangeUserResponse",
			Fields: graphql.InputObjectConfigFieldMap{
				"status": &graphql.InputObjectFieldConfig{
					Type: graphql.Boolean,
				},
			},
		})
	}
	return gql__input_ChangeUserResponse
}

func Gql__input_ChangeUserRequest() *graphql.InputObject {
	if gql__input_ChangeUserRequest == nil {
		gql__input_ChangeUserRequest = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Out_Input_ChangeUserRequest",
			Fields: graphql.InputObjectConfigFieldMap{
				"name": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"password": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"email": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		})
	}
	return gql__input_ChangeUserRequest
}

// graphql__resolver_AuthService is a struct for making query, mutation and resolve fields.
// This struct must be implemented runtime.SchemaBuilder interface.
type graphql__resolver_AuthService struct {

	// Automatic connection host
	host string

	// grpc dial options
	dialOptions []grpc.DialOption

	// grpc client connection.
	// this connection may be provided by user
	conn *grpc.ClientConn
}

// new_graphql_resolver_AuthService creates pointer of service struct
func new_graphql_resolver_AuthService(conn *grpc.ClientConn) *graphql__resolver_AuthService {
	return &graphql__resolver_AuthService{
		conn: conn,
		host: "localhost:50051",
		dialOptions: []grpc.DialOption{
			grpc.WithInsecure(),
		},
	}
}

// CreateConnection() returns grpc connection which user specified or newly connected and closing function
func (x *graphql__resolver_AuthService) CreateConnection(ctx context.Context) (*grpc.ClientConn, func(), error) {
	// If x.conn is not nil, user injected their own connection
	if x.conn != nil {
		return x.conn, func() {}, nil
	}

	// Otherwise, this handler opens connection with specified host
	conn, err := grpc.DialContext(ctx, x.host, x.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	return conn, func() { conn.Close() }, nil
}

// GetQueries returns acceptable graphql.Fields for Query.
func (x *graphql__resolver_AuthService) GetQueries(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{
		"hello": &graphql.Field{
			Type: Gql__type_HelloResponse(),
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type:         graphql.NewNonNull(graphql.String),
					DefaultValue: "",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req HelloRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for hello")
				}
				client := NewAuthServiceClient(conn)
				resp, err := client.SayHello(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC SayHello")
				}
				return resp, nil
			},
		},
		"user": &graphql.Field{
			Type: Gql__type_GetUserInfoResponse(),
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type:         graphql.NewNonNull(graphql.String),
					DefaultValue: "",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req GetUserInfoRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for user")
				}
				client := NewAuthServiceClient(conn)
				resp, err := client.GetUserInfo(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC GetUserInfo")
				}
				return resp, nil
			},
		},
	}
}

// GetMutations returns acceptable graphql.Fields for Mutation.
func (x *graphql__resolver_AuthService) GetMutations(conn *grpc.ClientConn) graphql.Fields {
	return graphql.Fields{
		"login": &graphql.Field{
			Type: Gql__type_CreateUserResponse(),
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type:         graphql.NewNonNull(graphql.String),
					DefaultValue: "",
				},
				"password": &graphql.ArgumentConfig{
					Type:         graphql.NewNonNull(graphql.String),
					DefaultValue: "",
				},
				"email": &graphql.ArgumentConfig{
					Type:         graphql.NewNonNull(graphql.String),
					DefaultValue: "",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req CreateUserRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for login")
				}
				client := NewAuthServiceClient(conn)
				resp, err := client.CreateUser(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC CreateUser")
				}
				return resp, nil
			},
		},

		"change": &graphql.Field{
			Type: Gql__type_ChangeUserResponse(),
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type:         graphql.NewNonNull(graphql.String),
					DefaultValue: "",
				},
				"password": &graphql.ArgumentConfig{
					Type:         graphql.NewNonNull(graphql.String),
					DefaultValue: "",
				},
				"email": &graphql.ArgumentConfig{
					Type:         graphql.NewNonNull(graphql.String),
					DefaultValue: "",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var req ChangeUserRequest
				if err := runtime.MarshalRequest(p.Args, &req, false); err != nil {
					return nil, errors.Wrap(err, "Failed to marshal request for change")
				}
				client := NewAuthServiceClient(conn)
				resp, err := client.ChangeUser(p.Context, &req)
				if err != nil {
					return nil, errors.Wrap(err, "Failed to call RPC ChangeUser")
				}
				return resp, nil
			},
		},
	}
}

// Register package divided graphql handler "without" *grpc.ClientConn,
// therefore gRPC connection will be opened and closed automatically.
// Occasionally you may worry about open/close performance for each handling graphql request,
// then you can call RegisterAuthServiceGraphqlHandler with *grpc.ClientConn manually.
func RegisterAuthServiceGraphql(mux *runtime.ServeMux) error {
	return RegisterAuthServiceGraphqlHandler(mux, nil)
}

// Register package divided graphql handler "with" *grpc.ClientConn.
// this function accepts your defined grpc connection, so that we reuse that and never close connection inside.
// You need to close it maunally when application will terminate.
// Otherwise, you can specify automatic opening connection with ServiceOption directive:
//
//	service AuthService {
//	   option (graphql.service) = {
//	       host: "host:port"
//	       insecure: true or false
//	   };
//
//	   ...with RPC definitions
//	}
func RegisterAuthServiceGraphqlHandler(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return mux.AddHandler(new_graphql_resolver_AuthService(conn))
}
