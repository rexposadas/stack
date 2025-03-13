import os
from typing import cast

import logfire
from pydantic import BaseModel

from pydantic_ai import Agent
from pydantic_ai.models import KnownModelName

# 'if-token-present' means nothing will be sent (and the example will work) if you don't have logfire configured
logfire.configure()


class MyModel(BaseModel):
    city: str
    country: str


model = cast(KnownModelName, os.getenv("PYDANTIC_AI_MODEL", "openai:gpt-4o"))
print(f"Using model: {model}")
agent = Agent("google-gla:gemini-1.5-flash", result_type=MyModel, instrument=True)

if __name__ == "__main__":
    result = agent.run_sync("The windy city in the US of A.")
    print(result.data)
    print(result.usage())
