# ads-service-module

`ads-service-module` is a Azure IoT Hub module that connects and sends events to the Azure IoT Hub.

## Example usage
Create a new deployment manifest that contains the module, e.g. `myapplication.yaml`.

```yaml
---
application: my-application
modules:
  - name: ads-node-module
    image: ads-node-module:latest
    createOptions: '{}'
    imagePullPolicy: on-create
    restartPolicy: always
    status: running
    startupOrder: 1
```

Deploy the manifest using the edgefarm-cli

```sh
$ edgefarm alm apply -f myapplication.yaml
```
