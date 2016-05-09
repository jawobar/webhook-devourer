# Webhook Devourer [![wercker status](https://app.wercker.com/status/523a0924fb03585bd2b5141526ee28b4/s/master "wercker status")](https://app.wercker.com/project/bykey/523a0924fb03585bd2b5141526ee28b4) [![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE.md) #

This is a simple and customisable HTTP(S) server ready to take action upon receiving any kind of webhook call it understands. 



### How does it work? ###

When a weebhook call is made, it's body is parsed by a specific handler mapped to the URL the request was sent to. Currently there are two kinds of handlers supported:

* Bitbucket
* DockerHub

Each handler can have assigned one or more so called runners. Runner is a piece of code doing the job you want as a reaction to webhook calls. Just before being executed it gets a context, provided by handler and containing some of the significant values extracted from the webhook's body. Currently there is only one runner supported:

* Bash

All this logic can be secured by using SSL (encryption) and/or API keys (to prevent unauthorized access).   

### How to use it? ###

Start by looking at included example of a [CONFIG](config.yml) file. Define handlers and map them to unique urls. Apart from type, you can specify the `auth` flag, which turns on the api key verification, as well as `log` flag, which prints the webhook's body on the console.

Next, add as many runners to each handler as you want. Their parameters can contain placeholders like `$REPO$`, which are replaced with values from context passed by the handler on each webhook call.

In order to use encryption, provide your key and certificate files in the `tls` section.

Add api keys to the `apikeys` section. Only request with specified `apikey` request parameter will be allowed for handlers with `auth` flag set to true.

### How to build it? ###

Webhook Devourer was developed within a dockerized environment. Hava a look at the [WERCKER](wercker.yml) file for dev and build pipelines.
 
### License ###

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).