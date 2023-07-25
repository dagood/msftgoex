# sum

The program prints an example response to the console when it starts. It also hosts a web service to return similar responses.

Go to <https://localhost:8080/sum?v=golang> to see something like this if GOEXPERIMENT is set to opensslcrypto:

```json
{
  "Request": "/sum?v=golang",
  "Version": "go1.20.6 X:opensslcrypto",
  "Type": "*openssl.sha256Hash",
  "Value": "golang",
  "Sum": "856d7285aaa4797447bb5cc1e5d266fd0682141947e4394b921ad6618bdddd6f",
  "Value2": "golang2023-07-25T00:08:53.137249726Z",
  "Sum2": "e24bbe318ebb88d99f2d911cc1ce300e06e19ed9bf25a17be968275a261b826b"
}
```

The `Type` is `*cng.shaXHash`, if GOEXPERIMENT is set to cngcrypto, and `*sha256.digest` if no backend is enabled.
