# WakaTime-Readme

Go GitHub Action that adds Coding time statistics to your README

<!--WakaTime-Start-->
<pre>Svelte                  52 mins ███████████░░░░░░░░░░░░░░ 45.84 %</br>JavaScript              32 mins ███████░░░░░░░░░░░░░░░░░░ 28.91 %</br>Go                      15 mins ███░░░░░░░░░░░░░░░░░░░░░░ 13.60 %</br>JSON                    12 mins ██░░░░░░░░░░░░░░░░░░░░░░░ 10.99 %</br>Other                    0 secs ░░░░░░░░░░░░░░░░░░░░░░░░░  0.28 %</br></pre>
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
name: WakaTime Readme

on:
  schedule:
    - cron: "0 1 * * *"
  workflow_dispatch:

env:
  WAKATIME_API_KEY: ${{ secrets.WAKATIME_API_KEY }}
  GH_TOKEN: ${{ secrets.GH_TOKEN  }}

jobs:
  WakaTime-README:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: NickRTR/WakaTime-Readme@v0.4.0
```

6. Wait for the action to run automatically every night or run it manually for testing purposes.
7. Let the magic happen 🚀

### Themes
