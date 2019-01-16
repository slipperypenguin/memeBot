# memeBot
🌶️ spicy memes

## About
Serverless side project built with [IBM Cloud Functions](https://console.bluemix.net/openwhisk). Automated deployments through [IBM Cloud Toolchain](https://console.bluemix.net/devops/getting-started)
Pulls random memes and delivers them (on a schedule) to a Slack Channel. Sources from [r/programmerhumor](https://www.reddit.com/r/ProgrammerHumor) and [xkcd](https://xkcd.com).

## Schedule
| Source 🔗   |  Day  📅  |  Time (EST)  |  Time (UTC)  |
|----------|:------:|:----------:|:----------:|
| Reddit | M-F |    9am, 12pm, 4pm    |    14:00, 17:00, 21:00    |
| xkcd |   M-F |    10:30am, 2:30pm, 5:30pm    |    15:30, 19:30, 22:30    |


## References
- [manifest examples](https://github.com/apache/incubator-openwhisk-wskdeploy/blob/master/docs/programming_guide.md#wskdeploy-utility-by-example)
- [IBM Cloud Functions Documentation](https://console.bluemix.net/docs/openwhisk/index.html#index)
