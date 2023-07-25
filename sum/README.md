# sum

Returns a JSON object with a checksum and other information. Just an example.

## Build

### Main Dockerfile

Useful build + run for the main `Dockerfile`:
Uses a multi-stage build to keep the final image small, per standard practice.

```sh
docker build . -t sum && docker run -p 8080:8080 -it --rm sum
```

### Dockerfile.OneStage

This Dockerfile runs on the build machine.
Not good for deployment, but easier to use to prod at results because it includes Go, and a distro.

```sh
docker build . -t sum -f Dockerfile.OneStage && docker run -p 8080:8080 -it --rm sum
```

To poke around, override the entrypoint:

```sh
docker build . -t sum -f Dockerfile.OneStage && docker run -p 8080:8080 -it --rm --entrypoint bash sum
```

## Results

Go to <https://localhost:8080/sum?v=golang>, see something like this:

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
