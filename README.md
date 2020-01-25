# memeBot
🌶️ spicy memes

## About
Serverless side project built with [IBM Cloud Functions](https://cloud.ibm.com/functions/), [IBM Cloud Toolchain](https://cloud.ibm.com/devops/getting-started), and [GitHub Actions](https://help.github.com/en/actions).
Pulls random memes and delivers them (on a schedule) to a Slack Channel. Sources from [r/programmerhumor](https://www.reddit.com/r/ProgrammerHumor) and [xkcd](https://xkcd.com).

## Schedule
| Source 🔗   |  Day  📅  |  Time (EST)  |  Time (UTC)  |
|----------|:------:|:----------:|:----------:|
| Reddit | M-F |    9am, 12pm, 4pm    |    14:00, 17:00, 21:00    |
| xkcd |   M-F |    2:30pm    |    19:30   |


## References
- [manifest examples](https://github.com/apache/incubator-openwhisk-wskdeploy/blob/master/docs/programming_guide.md#wskdeploy-utility-by-example)
- [IBM Cloud Functions Documentation](https://console.bluemix.net/docs/openwhisk/index.html#index)
