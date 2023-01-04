# 🪝 **Goseek:** A simple edge function workflow

One way to run a fully automated workflow on the edge for free is to use a combination of github actions and cloudflare workers.

You can specify an action in whatever language you desire and then create a CI pipeline that triggers based on a new created/modified issue. This will be the starting point and this pipeline will trigger the execution of an edge function on cloudflare.

The beauty of cloudflare workers is that it can run web assembly as well as JavaScript. This means that you can conserve your initial programming language without learning JavaScript and its vast complexities.

Let's say that we want a golang script that can extract elements from a webpage using custom selectors.

First of, you will write a simple go program that sends a request triggering the cloudflare worker. Then, the golang code will need to execute a custom business logic in order to fetch, parse, and extract from the page we want. As a last touch, we can write the results as comments in the issue and close it when the function exits.
