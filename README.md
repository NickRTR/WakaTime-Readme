# WakaTime-Readme

Go GitHub Action that adds Coding time statistics to your README

<!--WakaTime-Start-->
<pre><h2>Last 7 Days</h2>Svelte            5 hrs 48 mins 🟨🟨🟨🟨🟨🟨🟨🟨🟨🟨🟨🟨⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜  49.64 %</br>TypeScript        2 hrs 38 mins 🟨🟨🟨🟨🟨⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜  22.58 %</br>Go                 1 hr 44 mins 🟨🟨🟨⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜  14.94 %</br>JavaScript              55 mins 🟨🟨⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜   7.84 %</br>JSON                    11 mins ⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜   1.68 %</br><h2>All Time</h2><strong>Total Time Coded:   </strong>543 hrs 42 mins</br><strong>Timespan:           </strong>407 days</br><strong>Daily average:      </strong>1 hr(s) 19 min(s)</pre>
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
            - uses: NickRTR/WakaTime-Readme@main
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
| ⚪🟡   | `circle-yellow` |
| ⚪🔴   | `circle-red`    |
| ⚪🟣   | `circle-purple` |
| ⚪🟠   | `circle-orange` |
| ⚪🔵   | `circle-blue`   |
| ⚪⚫   | `circle-black`  |

Specify a theme by adding `THEME: <theme>` to the env variables of your workflow yaml file.

If no theme is specified, the default is selected.
