# IMPORTANT NOTES

## Package names

 - Because of the way that Go package management works, I've been
   developing this code in my own repository, using the package name
   `github.com/ian-ross/EdgeAuth/golang`.

 - This code (which is on branch `go-package` in this repository)
   works, and can be tested using the example in the `README`, but to
   release the code under the `PhenixRTS` namespace, the paths in the
   code need to be changed to refer to
   `github.com/PhenixRTS/EdgeAuth/golang`.

 - Code with the correct `PhenixRTS` paths is available on the
   `go-package-release` branch of this repository. The code on the
   `go-package` and `go-package-release` branches is identical except
   for these differences (and the presence of this file of notes).

 - (There are tricks to handle working in a fork like this once the
   "real" package is published, but they don't really work if there's
   not yet any Go package published at all, so it seemed simpler to do
   things this way.)


# Review, release and versioning

 - I suggest that you review the code on this `go-package` branch
   first, since it's possible to install and use it, and we can iron
   out any issues you have with it before release.

 - Releasing will involve:
 
    1. Merging the code from the `go-package-release` branch into the
       `github.com/PhenixRTS/EdgeAuth` repository.
       
    2. Tagging that repository with a tag of the form
       `golang/v0.0.1`.
       
    3. Telling the Go Module Index to pick up the new module by doing:
    
```shell script
GOPROXY=proxy.golang.org go list -m github.com/PhenixRTS/EdgeAuth/golang@v0.0.1
```


# Code issues for review

 - The Go JSON encoder in the standard library doesn't give any
   guarantees about key ordering in the encoding of JSON objects, in
   the way that Java's `JsonObjectBuilder` does, so I've used a
   [dependency](https://github.com/iancoleman/orderedmap) that
   implements an "ordered map" that does give those guarantees. If
   it's not OK to rely on an external dependency like that, let me
   know and I'll implement the functionality myself. It's not
   completely trivial, since it means writing a JSON decoder, but it's
   easy enough to base that on the decoder in the Go standard library.

 - Looking at the other language implementations, there seems to be a
   little ambiguity in how to canonicalise capabilities and tags lists
   for tokens. The Python implementation uses a set to remove
   duplicates in the capabilities list and sorts the resulting list of
   unique capabilities, while the Java and Javascript implementations
   just append capabilities to an array. That means that adding
   capabilities "b" and "a" in that order would result in different
   token digests between the Java(script) and Python library versions.
   Which is the correct behavior here?


# Behavior-Driven Development with `godog`

## Install `godog`

```
go install github.com/cucumber/godog/cmd/godog@v0.12.0
```

## Run `godog`

```
godog run ../features
```
