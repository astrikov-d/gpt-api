package main

// getGPTResponse schema
// prompt: user input for GPT prompt.
// context: user input for GPT context.
type getGPTResponseRequest struct {
	Prompt  string
	Context string
}
