# Six
### _The most trusted personal agent._
![N|Solid](https://beta.getsix.co/logo.png)

Six is an AI native, privacy and trust focused agent that acts on behalf of the user - making texts-to-action awesome.

✅ 🥕 Groceries \
✅ 🍔  Food delivery \
🔜 ✈️ Air ticket \
🔜 🛌 Hotel booking \
🔜 🚕 Ride hailing \
🔜 🍽️ Restaurant Reservation \
🔜 🧾 Bill payment

# Overview

This exercise will provide us with a working piece of code that can serve as neutral ground for technical discussion and on-site pair programming session. 

We expect the exercise to take no more than a day and recommend that you do not spend any more than that, regardless of completion. Please read this entire document before beginning.

Feel free to use whatever resources you need. Just like in the real world, Google, ChatGPT, etc. are tools of the trade. We prefer Go, but will accept solutions in other languages.

If you have questions, please reach out to us on Discord:

[![Join Six on Discord](https://img.shields.io/badge/discord-join-5865F2?logo=discord)](https://discord.gg/xBvmhz3k)

# The Problem - Make recommendation with AI

Your goal is to build a web service from which users can get real time recommendations on food deliveries. Users can give an ambiguous demand like "get me some Mexican food" or a clear instruction like "get me a grande latte from Starbucks". 

Your service should make up to three recommendations and all of them can be referenced to the original source. To take things a step further, you can store and cache the recommendations with a persistent database and caching database to speed up similar requests in the future.

A typical setup looks like this:
- A web server framework (e.g. Gin)
- A relational database (e.g. PostgreSQL)
- An LLM service (e.g. OpenAI)
- A third party aggregator / platform (e.g. UberEats, you might find something you need by simply press ```F12``` in their web app)
- A caching service (e.g. Redis)

Your client facing interface should look something like:

`GET /recommendations`

Return a JSON document with the following:
```
[
    {
        "recommendationId": "<id>",
        "store": {
            "name": "<name>",
            "id": "<id>",
            "menuItem": [
                {"name": "<name>", "id": "<id>", "price": "<price>"},
                {"name": "<name>", "id": "<id>", "price": "<price>"}
            ]
        }
    },
    {
        "recommendationId": "<id>",
        "store": {
            "name": "<name>",
            "id": "<id>",
            "menuItem": [
                {"name": "<name>", "id": "<id>", "price": "<price>"},
                {"name": "<name>", "id": "<id>", "price": "<price>"}
            ]
        }
    }
]

```

# Your solution

Prepare code to solve the problem. We encourage you to be mindful of time and start with an MVP solution and iterate if time permits (hint: you may not need persistence in your MVP).

Include a README with at least the following:

- Instructions for environment setup and running your code and tests
- Any additional context on your solution and approach, including any assumptions made
- What are the shortcomings of your solution?
- If you had additional time to work on this problem, what would you add or refine?

# Submitting results

Create a folder named ```solution``` and put your solution in it and submit a pull request on GitHub.


Good luck!
