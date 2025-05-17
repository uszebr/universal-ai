package base

type Role string

const (
	SystemRole    Role = "system"
	UserRole      Role = "user"
	AssistantRole Role = "assistant"
	ToolRole      Role = "tool"
)
const (
	FunctionToolType ToolCallType = "function"
)

type FullAIRequest struct {
	Model string `json:"model"`
	AIRequest
}

type AIRequest struct {
	Messages    []Message   `json:"messages"`
	N           *int        `json:"n,omitempty"`
	Temperature *float64    `json:"temperature,omitempty"`
	TopP        *float64    `json:"top_p,omitempty"`
	Tools       []Tool      `json:"tools,omitempty"`
	ToolChoice  *ToolChoice `json:"tool_choice,omitempty"`
}

type Tool struct {
	Type            string      `json:"type"`               //ex: "function"
	RequestFunction interface{} `json:"function,omitempty"` //actual code of function in string
}

type Message struct {
	Index      int        `json:"index,omitempty"`
	Role       Role       `json:"role"`
	Content    *string    `json:"content,omitempty"`      // Must be pointer to allow null for tool_calls
	ToolCallID *string    `json:"tool_call_id,omitempty"` // Only used when role is "tool"
	Name       *string    `json:"name,omitempty"`         // Optional: for system/tool functions
	ToolCalls  []ToolCall `json:"tool_calls,omitempty"`   // Present when assistant makes tool call
}

// for making function required(only one for now) in Request
// ex:  "tool_choice": {"type": "function", "function": {"name": "get_weather" }}
type ToolChoice struct {
	FunctionToolType   ToolCallType     `json:"type"` //ex: "function"
	ToolChoiceFunction ToolCallFunction `json:"function"`
}

type AIResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Choices []Choice `json:"choices"`
}
type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type ToolCall struct {
	ID               string           `json:"id"`
	ToolCallType     ToolCallType     `json:"type"`
	ToolCallFunction ToolCallFunction `json:"function"`
}

type ToolCallType string

type ToolCallFunction struct {
	Name      string  `json:"name"`
	Arguments *string `json:"arguments,omitempty"` //it is stored in string and need to be parsed(to Map); ex: "{\"location\":\"New York, NY\"}"
}
