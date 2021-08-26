# Basic

This code doesn't do anything particularly special.
This is an example of how to not write GO code to be able to unit test the code.
As we make no effort to encapsulate external packages.
- We can't force errors on the `http.Get` or `ioutil.ReadAll` function calls, to check our error handling is appropriate.
- `http.Get` makes an external HTTP call to a URL which we don't want in unit tests as these are unpredictable, slow and won't allow us to test the code consistently by providing the same outcome every time we test.
