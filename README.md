<div align="center">
  <a href="https://speakeasyapi.dev/">
    <img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
  <a href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" style="width: auto; height: 28px;" /></a>
  <a href="https://github.com/speakeasy-api/grpc-rest-service/issues">
    <img src="https://img.shields.io/github/issues/ritza-co/grpc-rest-service?style=for-the-badge" style="width: auto; height: 28px;" /></a>
    <!-- TODO: Change the Repo URL after publishing -->
  <a href="https://vscode.dev/redirect?url=vscode://ms-vscode-remote.remote-containers/cloneInVolume?url=https://github.com/ritza-co/grpc-rest-service">
    <img src="https://img.shields.io/static/v1?label=Dev%20Containers&message=Open&color=blue&logo=visualstudiocode" style="width: auto; height: 28px;" /></a>
    <!-- TODO: Change the Repo URL after publishing -->
</div>

<br />
<br />

<h1 align="center">Gnostic-gRPC and gRPC-Gateway Example with Generated SDKs</h1>
<div align="center">
    Example project using Gnostic to generate gRPC and gRPC-Gateway code, and the generated SDKs for Go, Python, and Node.js.
</div>
</br>

</br>

## Requirements

- [Docker](https://www.docker.com/)
- [VS Code](https://code.visualstudio.com/)
- [Remote - Containers Extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

## Getting Started

1. Click on the "Open in Dev Containers" badge above to open the project in a Dev Container.
2. Copy `.env.template` to `.env` and update the values.
3. Rebuid the container by running the `Remote-Containers: Rebuild Container` command.
4. Open a terminal in VS Code and run the following command:

```bash
go run main.go
```

5. Open a new terminal and run the following command:

```bash
grpcurl -plaintext -d '{"drink_name_request": {"name": "Pan Galactic Gargle Blaster"}}' localhost:50051 bar.Bar/GetDrink | sed 's/%3A/:/g; s/%2F/\//g; s/%3D/=/g; s/%3F/?/g; s/%26/\&/g' | jq
```

6. You should see a JSON response with the drink details.

## How to Change the API

1. Edit `bar.yaml`.
2. Run the following command to generate the a protobuf binary:

```bash
gnostic --pb-out=. bar.yaml
```

3. Run the following command to generate the gRPC protobuf file with annotations:

```bash
gnostic-grpc -input bar.pb -output .
```

4. Run the following command to generate the gRPC and gRPC-Gateway code:

```bash
buf generate
```

5. Run the following command to generate the SDKs:

```bash
speakeasy generate sdk \
    --schema bar.yaml \
    --lang typescript \
    --out ./bar-sdk-typescript
```

6. Rebuild the container by running the `Remote-Containers: Rebuild Container` command.

7. Continue from step 4 in the [Getting Started](#getting-started) section.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.