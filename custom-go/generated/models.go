package generated

type Disabled_apiResponseData_todo_findManyTodo struct {
    Id int64 `json:"id"`
}
type Todo__DeleteOneTodoResponseData_data struct {
    Id int64 `json:"id"`
}
type Todo__GetManyTodoResponseData struct {
    Data []Todo__GetManyTodoResponseData_data `json:"data"`
}
type Todo__CreateOneTodoInternalInput struct {
    Title string `json:"title"`
}
type Todo__GetManyTodoResponseData_data struct {
    Completed bool `json:"completed"`
    CreatedAt string `json:"createdAt"`
    Id int64 `json:"id"`
    Title string `json:"title"`
}
type Todo__CreateOneTodoResponseData struct {
    Data Todo__CreateOneTodoResponseData_data `json:"data,omitempty"`
}
type RootInternalInput struct {
    Skip int64 `json:"skip,omitempty"`
    Take int64 `json:"take,omitempty"`
}
type Disabled_apiInput struct {
}
type Root_apiInput struct {
    Skip int64 `json:"skip,omitempty"`
    Take int64 `json:"take,omitempty"`
}
type RootResponseData struct {
    Todo_findManyTodo []RootResponseData_todo_findManyTodo `json:"todo_findManyTodo"`
}
type UpperResponseData struct {
    Todo_findFirstTodo UpperResponseData_todo_findFirstTodo `json:"todo_findFirstTodo,omitempty"`
}
type Todo__DeleteOneTodoInput struct {
    Id int64 `json:"id"`
}
type UpperInput struct {
    Equals int64 `json:"equals,omitempty"`
}
type Todo__UpdateOneTodoResponseData_data struct {
    Completed bool `json:"completed"`
    CreatedAt string `json:"createdAt"`
    Id int64 `json:"id"`
    Title string `json:"title"`
}
type UpperInternalInput struct {
    Equals int64 `json:"equals,omitempty"`
}
type Root_apiResponseData_todo_findManyTodo struct {
    Completed bool `json:"completed"`
    Id int64 `json:"id"`
    Title string `json:"title"`
}
type Todo__UpdateOneTodoInternalInput struct {
    Id int64 `json:"id"`
    Title string `json:"title"`
}
type RootInput struct {
    Skip int64 `json:"skip,omitempty"`
    Take int64 `json:"take,omitempty"`
}
type Disabled_apiResponseData struct {
    Todo_findManyTodo []Disabled_apiResponseData_todo_findManyTodo `json:"todo_findManyTodo"`
}
type Todo__CreateOneTodoInput struct {
    Title string `json:"title"`
}
type Todo__UpdateTodoCompletedInput struct {
    Completed bool `json:"completed"`
    Id int64 `json:"id"`
}
type Todo__UpdateTodoCompletedResponseData_data struct {
    Completed bool `json:"completed"`
    CreatedAt string `json:"createdAt"`
    Id int64 `json:"id"`
    Title string `json:"title"`
}
type RootResponseData_todo_findManyTodo struct {
    Completed bool `json:"completed"`
    Id int64 `json:"id"`
    Title string `json:"title"`
}
type Todo__UpdateOneTodoResponseData struct {
    Data Todo__UpdateOneTodoResponseData_data `json:"data,omitempty"`
}
type Todo__UpdateOneTodoInput struct {
    Id int64 `json:"id"`
    Title string `json:"title"`
}
type UpperResponseData_todo_findFirstTodo struct {
    Completed bool `json:"completed"`
    Id int64 `json:"id"`
    Title string `json:"title"`
}
type Todo__DeleteOneTodoInternalInput struct {
    Id int64 `json:"id"`
}
type Todo__UpdateTodoCompletedInternalInput struct {
    Completed bool `json:"completed"`
    Id int64 `json:"id"`
}
type Todo__GetManyTodoInternalInput struct {
}
type Todo__UpdateTodoCompletedResponseData struct {
    Data Todo__UpdateTodoCompletedResponseData_data `json:"data,omitempty"`
}
type Todo__CreateOneTodoResponseData_data struct {
    Completed bool `json:"completed"`
    CreatedAt string `json:"createdAt"`
    Id int64 `json:"id"`
    Title string `json:"title"`
}
type Todo__DeleteOneTodoResponseData struct {
    Data Todo__DeleteOneTodoResponseData_data `json:"data,omitempty"`
}
type Disabled_apiInternalInput struct {
}
type Todo__GetManyTodoInput struct {
}
type Root_apiResponseData struct {
    Todo_findManyTodo []Root_apiResponseData_todo_findManyTodo `json:"todo_findManyTodo"`
}
type Root_apiInternalInput struct {
    Skip int64 `json:"skip,omitempty"`
    Take int64 `json:"take,omitempty"`
}

