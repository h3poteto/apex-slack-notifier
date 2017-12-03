# apex-slack-notifier
AWS lambda functions for my aws notification. These functions notify messages to slack.

## Install

```
$ git clone https://github.com/h3poteto/apex-slack-notifier.git
$ cd apex-slack-notifier
```

These functions require `function.json` file to set environment variables for each functions.

You should set slack information in this file, for example:

```json
{
  "SLACK_URL": "https://hooks.slack.com/services/hoge/fuga/hogehoge",
  "SLACK_CHANNEL": "#general"
}
```

## Deploy
### login
This function receive login notification from CloudWatch Event, but CloudWatch Event can catch console login event only us-east-1 region.
So you must deploy this function in us-east-1.

```
$ apex deploy login -r us-east-1
   • updating config           env= function=login
   • updating function         env= function=login
   • updated alias current     env= function=login version=16
   • function updated          env= function=login name=slack-notifier_login version=16
```

Of course you must create CloudWatch Event in us-east-1 region.

### alarm
This function receive CloudWatch Alarm through AWS SNS.

```
$ apex deploy alarm
```
