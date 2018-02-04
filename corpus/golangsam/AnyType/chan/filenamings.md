## Source file names

TODO: This is outdated!

Note: currently, there is exactly one file named "chan.go" (and some related chan*_test.go)
in "runtime": src/runtime/chan.go. The implementation of channels.

Thus: We feel free to choose and use "chan" as prefix for all our source files.
Even if we 'inject' such into existing packages in order to enhance them in place.

More specifically:
* "chan.go"
  for manual contributions

* "chan\_generic.go"
  for generated source files ( if _Type_ == "" )
* "chan\_ _Type_ \_generic.go"
  for generated source files ( if _Type_ != "" )
