# Zoo Wrangler

## Description

Zoo Wrangler is meant to be a replacement for [exhibitor](https://github.com/soabase/exhibitor).  

Here is a list of some features we want to have in our first release.

* Zookeeper configuration sync over S3
* Process manager for zookeeper (start/stop/restart)
* API to Zookeeper related information
* Back up & Recovery
* Visualization


## Status

Currently in the beginning of development nothing concrete pull requests welcome!

## Contributing

1. Fork it
2. Create a ticket on the main repo https://github.com/amedeiros/zoo-wrangler/issues.
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Commit your changes (`git commit -am 'Add some feature'`)
5. buffalo test
6. Push to the branch (`git push origin my-new-feature`)
7. Create new Pull Request
8. Wait for review. Peer review is important please resolve any issues found in this step by repeating steps 3-5.

## Documentation

To view generated docs for zoo-wrangler, run the below command and point your browser to http://127.0.0.1:6060/pkg/

    godoc -http=:6060 2>/dev/null &

 ## Run Tests

    buffalo test

 ## Run in dev

    buffalo dev

[Powered by Buffalo](http://gobuffalo.io)

