# Ginkgo & Gomega Talk 2020

## About the author
I'm Alexander Egurnov, Senior Software Developer at PubNative GmbH. 
I've been coding in Go since March 2015.

## About this repo
This is an introduction and a live demo of my favorite features of [Ginkgo](https://onsi.github.io/ginkgo/)
testing framework and it's accompanying assertion library [Gomega](http://onsi.github.io/gomega/).

### Talking points
* Gomega
    * Standalone
    * `Expect` and `Ω` notations
    * Annotations
    * Matchers
        * Error checking and panics
        * Equivalence
        * Numbers
        * Other types
        * Collections
    * Combining matchers
    * Async assertions
    * Transformations and custom matchers
* Ginkgo
    * CLI
        * `ginkgo bootstrap` and `ginkgo generate`
        * `go test` interoperability
        * `GinkgoWriter` and `-v` option
        * `ginkgo -p`
        * `ginkgo -r`
        * `--flakeAttempts=ATTEMPTS`
        * `-untilItFails`
        * `ginkgo watch -notify -r`
    * Connecting Gomega
    * `ginkgo blur/unfocus`
    * Setup and teardown
    * Async functions
    * Measurements  
* General advice
    * Keeping test readable 