from pydantic_ai import Agent

agent = Agent(
    "google-gla:gemini-1.5-flash",
    system_prompt="Be concise and to the point.",
)

result = agent.run_sync("What is the meanest way to say hello?")
print(result.data)
