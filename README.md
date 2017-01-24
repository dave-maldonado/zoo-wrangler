# Zoo Wrangler

## Description

Zoo Wrangler is meant to be a replacement for [exhibitor](https://github.com/soabase/exhibitor).  

Here is a list of some features we want to have in our first release.

* Zookeeper configuration sync over S3
* Process manager for zookeeper (start/stop/restart)
* API to Zookeeper related information
* Back up & Recovery
* Visualization


## Documentation

To view generated docs for zoo-wrangler, run the below command and point your browser to http://127.0.0.1:6060/pkg/

    godoc -http=:6060 2>/dev/null &

 ## Run Tests

    buffalo test

 ## Run in dev

    buffalo dev

[Powered by Buffalo](http://gobuffalo.io)

