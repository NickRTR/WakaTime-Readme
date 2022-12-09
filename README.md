# WakaTime-Readme

Go GitHub Action that adds Coding time statistics to your README

<!--WakaTime-Start-->
<pre>Go               13 hrs 30 mins ██████████████████████░░░ 89.13 %</br>YAML                    53 mins █░░░░░░░░░░░░░░░░░░░░░░░░  5.86 %</br>Other                   15 mins ░░░░░░░░░░░░░░░░░░░░░░░░░  1.65 %</br>JSON                    12 mins ░░░░░░░░░░░░░░░░░░░░░░░░░  1.39 %</br>Svelte                   6 mins ░░░░░░░░░░░░░░░░░░░░░░░░░  0.66 %</br></pre>
<!--WakaTime-End-->

## Documentation

1. Update the markdown file with 2 comments:
   `<!--WakaTime-Start-->` and `<!--WakaTime-End-->`
2. Get your WakaTime API Key (https://wakatime.com/settings/account)
3. Generate a GitHub API Token with repo and user scope (https://github.com/settings/tokens)
4. Store the WakaTime API Key and GitHub access token in your repository's action secrets as following:
   - WakaTime API Key: `WAKATIME_API_KEY=<your key>`
   - GitHub Token: `GH_TOKEN=<your Github access token>`
5. Create a new GitHub Action in the Repository, you want the Readme stats in:

```yml
name: Go

on:
  schedule:
    - cron: "0 1 * * *"
  workflow_dispatch:

env:
  WAKATIME_API_KEY: ${{ secrets.WAKATIME_API_KEY }}
  GH_TOKEN: ${{ secrets.GH_TOKEN  }}

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - name: Set up Go
        uses: actions/setup-go@v3
        with:
          go-version: 1.18

      - name: Go WakaTime Readme
        uses: NickRTR/WakaTime-Readme@v0.3.0
```

6. Wait for the action to run automatically every night or run it manually for testing purposes.
7. Let the magic happen 🚀

### Themes
