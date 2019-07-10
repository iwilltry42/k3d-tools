# k3d-tools

k3d-tools support the usage of [https://github.com/rancher/k3d](https://github.com/rancher/k3d)

## Tests

Test saving images from the local docker daemon to a volume e.g. via `docker run --rm -v /var/run/docker.sock:/var/run/docker.sock -v $(pwd)/images:/images --privileged iwilltry42/k3d-tools:v0.0.1 save-image nginx:local nginx:latest`
