
# GoGPT

A simple CLI implementation with some colored output using OpenAI's GPT-3.5-turbo and Go. You can choose your role as a writer, programmer or teacher. It's just a first test hacked together in an hour or so.
**BIG Thanks** to @sashabaranov and @fatih (took me some time to figure out how to make colorized output in Go haha ...)
![image](https://user-images.githubusercontent.com/33793809/226075172-999091ba-586d-4bbf-960f-ab808752eb51.png)
## Dependencies (only if you want to fuck around but you need Go to build this or run it obviously)
- OpenAI API key - for lazy people like me, here is a direct link where to find/create on their [website](https://platform.openai.com/account/api-keys) OpenAI
- Go 1.17+ 
- github.com/fatih/color
- github.com/sashabaranov/go-openai

## Setup

1. Install Go and set up your Go environment. Tested with v1.19.6
2. Just clone the stuff:

    ```
git clone https://github.com/gyrusdentatus/GoGPT.git GoGPT
    ```

3. Replace `"your token"` with your actual OpenAI API key.

## Running the chatbot

To run the chatbot, simply execute the following command:

```

go run main.go

```

or build the whole thingy with `build` instead of run. You know the drill ... ;] 

## Note:
This is WIP and obviously I need to implement some memory recall as the regular chatgpt client, but surpringly it took just a few lines of code.
Happy for any contributions, PRs and comments or whatever shitposts. Feel free to cringe if you check the first commit. I always check for these things but in this case it does not matter because I deleted the *** already :) 
Otherwise I would have never pushed it anyway. 

