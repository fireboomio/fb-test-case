package generated

import "custom-go/pkg/base"

type (

    Todo__CreateOneTodoBody = *base.OperationBody[Todo__CreateOneTodoInternalInput, Todo__CreateOneTodoResponseData]
    
    Todo__DeleteOneTodoBody = *base.OperationBody[Todo__DeleteOneTodoInternalInput, Todo__DeleteOneTodoResponseData]
    
    Todo__GetManyTodoBody = *base.OperationBody[Todo__GetManyTodoInternalInput, Todo__GetManyTodoResponseData]
    
    Todo__UpdateOneTodoBody = *base.OperationBody[Todo__UpdateOneTodoInternalInput, Todo__UpdateOneTodoResponseData]
    
    Todo__UpdateTodoCompletedBody = *base.OperationBody[Todo__UpdateTodoCompletedInternalInput, Todo__UpdateTodoCompletedResponseData]
    
    UpperBody = *base.OperationBody[UpperInternalInput, UpperResponseData]
    
    Disabled_apiBody = *base.OperationBody[Disabled_apiInternalInput, Disabled_apiResponseData]
    
    RootBody = *base.OperationBody[RootInternalInput, RootResponseData]
    
    Root_apiBody = *base.OperationBody[Root_apiInternalInput, Root_apiResponseData]
    )

const (
    Todo__GetManyTodo base.OperationQueryPath = "Todo/GetManyTodo"
    Upper base.OperationQueryPath = "Upper"
    Disabled_api base.OperationQueryPath = "disabled_api"
    Root base.OperationQueryPath = "root"
    Root_api base.OperationQueryPath = "root_api"
)

const (
    Todo__CreateOneTodo base.OperationMutationPath = "Todo/CreateOneTodo"
    Todo__DeleteOneTodo base.OperationMutationPath = "Todo/DeleteOneTodo"
    Todo__UpdateOneTodo base.OperationMutationPath = "Todo/UpdateOneTodo"
    Todo__UpdateTodoCompleted base.OperationMutationPath = "Todo/UpdateTodoCompleted"
)
