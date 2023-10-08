package server

import (
    "github.com/joho/godotenv"

    "custom-go/global"

    "custom-go/generated"
        
    operation_Root "custom-go/operation/root"

    operation_Root_api "custom-go/operation/root_api"
                    
    "custom-go/pkg/base"
    "custom-go/pkg/plugins"
    "custom-go/pkg/types"
)

const nodeEnvFilepath = "../.env"

func init() {
    _ = godotenv.Overload(nodeEnvFilepath)

    types.WdgHooksAndServerConfig = types.WunderGraphHooksAndServerConfig{
        Hooks: types.HooksConfiguration{
            Global: plugins.GlobalConfiguration{
                HttpTransport: plugins.HttpTransportHooks{
                    BeforeOriginRequest: global.BeforeOriginRequest,
                    OnOriginRequest: global.OnOriginRequest,
                    OnOriginResponse: global.OnOriginResponse,
                },
        
            },
            Authentication: plugins.AuthenticationConfiguration{
            },
            Queries: base.OperationHooks{
                "root": {
                    PreResolve: plugins.ConvertBodyFunc[generated.RootInternalInput, generated.RootResponseData](operation_Root.PreResolve),
                    MutatingPostResolve: plugins.ConvertBodyFunc[generated.RootInternalInput, generated.RootResponseData](operation_Root.MutatingPostResolve),
                    MutatingPreResolve: plugins.ConvertBodyFunc[generated.RootInternalInput, generated.RootResponseData](operation_Root.MutatingPreResolve),
                    PostResolve: plugins.ConvertBodyFunc[generated.RootInternalInput, generated.RootResponseData](operation_Root.PostResolve),
                },
                "root_api": {
                    CustomResolve: plugins.ConvertBodyFunc[generated.Root_apiInternalInput, generated.Root_apiResponseData](operation_Root_api.CustomResolve),
                },
            },
            Mutations: base.OperationHooks{
            },
        },
    }
}