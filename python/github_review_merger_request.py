#  pip install fastapi uvicorn requests openai python-dotenv

from fastapi import FastAPI, Request
import requests
import openai
import os
from dotenv import load_dotenv

load_dotenv()

app = FastAPI()

GITHUB_TOKEN = os.getenv("GITHUB_TOKEN")
OPENAI_API_KEY = os.getenv("OPENAI_API_KEY")
openai.api_key = OPENAI_API_KEY

REVIEW_PROMPT_TEMPLATE = """
You are a highly experienced Senior Software Engineer tasked with conducting a thorough Git Merge Request (MR) review. Your goal is to provide constructive, actionable feedback that helps improve code quality, maintainability, and robustness.

Please meticulously review the following code changes. Analyze them and provide feedback structured as follows. If the provided code is in Go, pay special attention to adherence to the Google Go Style Guide (https://google.github.io/styleguide/go/) in relevant sections, particularly Naming Conventions and Best Practices.

Overall Summary & Key Changes:

Briefly state the main purpose or goal of this MR (e.g., "This MR introduces a new feature for X," or "This MR refactors Y to improve performance/clarity").

Summarize the most significant additions, modifications, or removals in terms of modules, functions, or core logic.

Code Clarity & Readability:

Is the code's intent immediately obvious? Is the logical flow easy to follow?

Are there any sections that are overly dense, convoluted, or require significant mental effort to decipher?

Suggest specific improvements for clarity (e.g., breaking down complex functions into smaller, single-responsibility units, simplifying conditional expressions, improving overall code structure or layout).

Are variable and function scopes as limited as possible?

Naming Conventions & Comments:

Are function, variable, class, struct, interface, and module names clear, descriptive, unambiguous, and consistent with common conventions for the language/project? (For Go, refer to the Google Go Style Guide: https://google.github.io/styleguide/go/)

Are comments used effectively to explain why something is done (the intent, trade-offs, or complex logic), rather than just what is being done (which the code should ideally convey)?

Are there any misleading, outdated, or redundant comments? Are there places where comments are missing and would be beneficial (e.g., for non-obvious business logic or complex algorithms)?

Is package naming and organization logical (if applicable)?

Simplicity & Complexity Reduction (KISS Principle):

Are there any overly complex conditional statements (e.g., deeply nested if/else), convoluted loops, or unnecessarily complicated algorithms?

Can any part of the code be refactored for greater simplicity or to reduce cognitive load without sacrificing functionality or clarity? Provide specific suggestions (e.g., using helper functions, better data structures, or design patterns).

Is there any "magic" (unexplained constants, arbitrary-seeming decisions) that could be clarified or made more explicit?

Potential Bugs & Edge Cases:

Do you foresee any potential bugs, race conditions, resource leaks (e.g., unclosed files/connections, goroutine leaks in Go), or off-by-one errors?

Are edge cases (e.g., null/nil inputs, empty collections, zero values, invalid data formats) and error conditions handled gracefully, comprehensively, and explicitly? Are errors propagated or handled appropriately?

(For Go: Are error values consistently checked and handled? Is panic used appropriately, i.e., rarely and for truly unrecoverable situations?)

Are there any unsafe operations or potential security vulnerabilities (e.g., SQL injection, XSS, improper input validation, data exposure, insecure use of cryptography)?

Best Practices & Design Principles:

Does the code adhere to relevant language-specific idioms and best practices? (For Go, this includes aspects like effective error handling, interface usage, concurrency patterns, and package design from the Google Go Style Guide.)

Does it follow established software design principles (e.g., DRY - Don't Repeat Yourself, SOLID, KISS - Keep It Simple, Stupid)?

Are there considerations for performance, scalability, or resource consumption that should be addressed? (e.g., inefficient algorithms, unnecessary allocations, blocking I/O in critical paths).

Is the code testable? Are dependencies managed in a way that facilitates testing (e.g., dependency injection)?

Are there any anti-patterns present? If so, suggest alternatives.

Is logging sufficient and appropriately leveled?

Testing (If Applicable/Inferable):

If tests are included or modified, are they clear, comprehensive, and effective in verifying the changes?

Do tests cover success paths, failure paths, and relevant edge cases for the new/modified code?

Are there any obvious missing test cases for the changes introduced?

Are tests isolated and repeatable? Do they clean up after themselves?

Recommendation & Rationale:

Based on your review, provide a clear recommendation:

Approve: All looks good, or only trivial nits that can be addressed post-merge or are stylistic preferences not violating guidelines.

Approve with suggestions: Minor changes are recommended (e.g., for clarity, minor refactoring) but are not blockers to merging. The author can consider them.

Request changes: Significant issues (e.g., potential bugs, major design flaws, missing critical tests, violations of core principles) need addressing before merging.

Clearly explain your reasoning, prioritizing critical issues over minor suggestions. If requesting changes, list the specific blocking issues that must be resolved.

{{diff}}
Respond only with a detailed review following the 8 steps above. Be specific, constructive, and provide examples from the code where appropriate to illustrate your points.
"""

@app.post("/webhook")
async def github_webhook(request: Request):
    event = request.headers.get('X-GitHub-Event')
    payload = await request.json()

    if event == "pull_request" and payload["action"] in ["opened", "synchronize"]:
        pr = payload["pull_request"]
        repo = payload["repository"]
        owner = repo["owner"]["login"]
        repo_name = repo["name"]
        pr_number = pr["number"]

        # Get the PR diff
        diff_url = pr["diff_url"]
        headers = {"Authorization": f"token {GITHUB_TOKEN}"}
        diff = requests.get(diff_url, headers=headers).text

        # Send to OpenAI
        prompt = REVIEW_PROMPT_TEMPLATE.format(diff=diff[:10000])  # truncate if too long
        response = openai.ChatCompletion.create(
            model="gpt-4-turbo",
            messages=[{"role": "user", "content": prompt}]
        )
        review_text = response["choices"][0]["message"]["content"]

        # Post review to GitHub
        comment_url = f"https://api.github.com/repos/{owner}/{repo_name}/issues/{pr_number}/comments"
        requests.post(comment_url,
                      headers={"Authorization": f"token {GITHUB_TOKEN}"},
                      json={"body": review_text})

        return {"message": "Review posted"}

    return {"message": "Event ignored"}
