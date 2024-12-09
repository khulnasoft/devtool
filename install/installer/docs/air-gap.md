# Installing Devtool in an air-gap network with the Devtool Installer

## Mirror Devtool Images

You need a registry that is reachable in your network. Add the domain of your registry to the Devtool config `devtool.config.yaml` like this:
```yaml
repository: your-registry.example.com
```

The command `devtool-installer mirror list` gives you a list of all images needed by Devtool. You can run the following code to pull the needed images and push them to your registry:

```
for row in $(devtool-installer mirror list --config devtool.config.yaml | jq -c '.[]'); do
    original=$(echo $row | jq -r '.original')
    target=$(echo $row | jq -r '.target')

    docker pull $original
    docker tag $original $target
    docker push $target
done
```

## Install Devtool in Air-Gap Mode

To install Devtool in an air-gap network, you need to configure the repository of the images needed by Devtool (see previous step). Add this to your Devtool config:

```yaml
repository: your-registry.example.com
```

That's it. Run the following commands as usual and Devtool fetches the images from your registry and does not need internet access to operate:

```
devtool-installer render --config devtool.config.yaml > devtool.yaml
kubectl apply -f devtool.yaml
```
