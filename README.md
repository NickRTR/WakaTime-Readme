# WakaTime-Readme

Go GitHub Action that adds Coding time statistics to your README

<!--WakaTime-Start-->
<pre>Svelte            4 hrs 53 mins ████████████░░░░░░░░░░░░░ 50.62 %</br>JavaScript         1 hr 50 mins ████░░░░░░░░░░░░░░░░░░░░░ 19.11 %</br>Markdown           1 hr 48 mins ████░░░░░░░░░░░░░░░░░░░░░ 18.66 %</br>JSON                    31 mins █░░░░░░░░░░░░░░░░░░░░░░░░  5.43 %</br>Bash                    14 mins ░░░░░░░░░░░░░░░░░░░░░░░░░  2.44 %</br></pre>
<!--WakaTime-End-->

## Documentation

1. Update the markdown file with 2 comments:
   `<!--WakaTime-Start-->` and `<!--WakaTime-End-->`
2. Get your WakaTime API Key (https://wakatime.com/settings/account)
3. Generate a GitHub API Token with repo and user scope (https://github.com/settings/tokens)
4. Store the WakaTime API Key and GitHub access token in your repository's action secrets as following:
   - WakaTime API Key: `WAKATIME_API_KEY=<your key>`
   - GitHub Token: `GH_TOKEN=<your GitHub access token>`
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
  THEME: "default"

jobs:
  WakaTime-README:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: NickRTR/WakaTime-Readme@v1.0.0
```

6. Wait for the action to run automatically every night or run it manually for testing purposes.
7. Let the magic happen 🚀

### Themes

There are a number of different themes, you can choose from.

| Theme  | env variable    |
| ------ | --------------- |
| ░█     | `default`       | 
| ------ | --------------  |
| ⬜🟩   | `block-green`   |
| ⬜🟨   | `block-yellow`  |
| ⬜🟥   | `block-red`     |
| ⬜🟪   | `block-purple`  |
| ⬜🟧   | `block-orange`  |
| ⬜🟦   | `block-blue`    |
| ⬜⬛   | `block-black`   |
| ------ | --------------  |
| ⚪🟢   | `circle-green`  |
| ⬜🟡   | `circle-yellow` |
| ⬜🔴   | `circle-red`    |
| ⬜🟣   | `circle-purple` |
| ⬜🟠   | `circle-orange` |
| ⬜🔵   | `circle-blue`   |
| ⬜⚫   | `circle-black`  |

Specify a theme by adding `THEME: <theme>` to the env variables of your workflow yaml file.

If no theme is specified, the default is selected.
