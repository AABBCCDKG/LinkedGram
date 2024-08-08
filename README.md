# Chat with your email - gmaillm

## How to use:
A:
`pip install -r requirements.txt`

B: gmail part
1. get [gmail api key](https://developers.google.com/gmail/api/guides?hl=zh-cn)
2. download JSON file and rename it as  `credentials.json`
3. put `credentials.json` in the project directory

C: LLM part
1. get [openai api key](https://platform.openai.com/api-keys)
2. store "OPENAI_API_KEY=<your openai api key>" to the file `.env.local` in the project directory
