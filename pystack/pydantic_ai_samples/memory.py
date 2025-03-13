import asyncio

from pydantic import BaseModel, Field
from pydantic_ai import Agent, RunContext
from typing import List
"""
# Python Coding Test: Conversation Summary for an LLM

## Problem Statement

You are tasked with developing a Python program that simulates an interactive chat between a user and an agent. The key requirement is to implement a *summary memory* that continuously condenses the conversation history. This summary is intended to be used as contextual input for a large language model (LLM) during subsequent interactions, ensuring that the LLM can generate informed responses without needing the entire conversation history.

## Requirements

1. **Interactive Conversation**  
   - The program should continuously prompt the user for input (e.g., via the console).
   - The agent should generate responses using a simple, rule-based approach.
   - At each turn, the agent's response should incorporate or reference the current summary memory.

2. **Summary Memory Implementation**  
   - **Maintaining State:**  
     Create a mechanism (using functions, a class, or another structure) to hold the conversation history.
   - **Summarization:**  
     Instead of storing every message, build and update a concise summary that captures the main points of the conversation.  
     *Hint:* Consider a heuristic approach such as appending keywords, topics, or key phrases from each interaction.
   - **Usage with an LLM:**  
     Assume the summary will eventually be sent as context to an LLM. Your agent’s responses should reflect the information captured in the summary. For example:
     > "Based on our previous discussion where you mentioned _[summary]_, ..."

3. **Dynamic Updates and Edge Case Handling**  
   - The summary memory should be updated after each turn of the conversation.
   - Handle cases such as empty user input or very long conversations gracefully.
   - Aim for a balance: the summary should be detailed enough to be useful, yet concise enough to avoid overwhelming the LLM context window.

## Example Interaction

Below is an example of how the conversation might flow:

User: Hi there! Agent: Hello! How can I help you today?
[Summary Memory: "User greeted."]

User: I'm interested in learning Python. Agent: Great! Earlier you mentioned interest in learning Python. What would you like to start with: basics or advanced topics?
[Summary Memory: "User greeted; interested in learning Python."]

User: Let's start with the basics. Agent: Since you want to start with the basics, here are some resources to get you started. Would you also like tips on setting up your development environment?
[Summary Memory: "User greeted; interested in learning Python; wants to start with basics."]
"""


class Summary(BaseModel):
    conversation_history: List[str] = Field(default_factory=list, description="Complete conversation history")
    current_summary: str = Field(default="", description="Current summary of the conversation")


chat_agent = Agent(
    "google-gla:gemini-1.5-flash",
    deps_type=Summary,
    system_prompt="""You are a helpful assistant with perfect memory of our conversation.
    The complete conversation history is available in deps.conversation_history.
    The current summary of our conversation is in deps.current_summary.
    
    IMPORTANT: Before responding, ALWAYS review the conversation history and summary 
    to maintain context and provide accurate responses about previously discussed information."""
)


summary_agent = Agent(
    "google-gla:gemini-1.5-flash",
    system_prompt="""You are a summarization assistant.
    Create a concise but detailed summary of the conversation that includes:
    - All personal information shared (names, preferences, etc.)
    - Key discussion points
    - Important context"""
)


async def main():
    memory = Summary(conversation_history=[], current_summary="No conversation yet.")
    print("Chat started. Type 'exit' to end the conversation.")

    while True:
        user_input = input("\nYou: ")

        if user_input.lower() == 'exit':
            print("Chat ended.")
            break

        try:
            # Add user input to history
            memory.conversation_history.append(f"User: {user_input}")
            
            # Generate new summary
            full_conversation = "\n".join(memory.conversation_history)
            summary_result = await summary_agent.run(full_conversation)
            memory.current_summary = summary_result.data
            
            # Get AI response using both history and summary
            response = await chat_agent.run(
                f"""Context: {memory.current_summary}
                Question: {user_input}""",
                deps=memory
            )
            
            # Add AI response to history
            memory.conversation_history.append(f"Assistant: {response.data}")
            
            print(f"\nAgent: {response.data}")

        except Exception as e:
            print(f"\nError: {e}")


if __name__ == "__main__":
    asyncio.run(main())
