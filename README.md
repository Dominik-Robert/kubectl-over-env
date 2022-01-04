# kubectl-over-env
This repository was created for github actions or any other CI/CD-Pipelining tool which cannot push to a custom local kubernetes cluster. So inside this container is a kubectl binary in the version of the tag (currently only v1.22.2) and you can manage the config for your cluster with environment variables or with a volume mount.

# Usage 

## with environment variables
```bash
docker run -it -e CERTIFICATE_AUTHORITY_DATA=DATA -e CLIENT_CERTIFICATE_DATA=DATA -e CLIENT_KEY_DATA=DATA -e SERVER=DATA ghcr.io/dominik-robert/kubectl-over-env:v1.22.2 apply -f Deployment.yaml

# or with local environment variables

docker run -it -e $CERTIFICATE_AUTHORITY_DATA -e $CLIENT_CERTIFICATE_DATA -e $CLIENT_KEY_DATA -e $SERVER ghcr.io/dominik-robert/kubectl-over-env:v1.22.2 apply -f Deployment.yaml
```
## with volume mount
docker run -it -v PATH_TO_YOUR_CONFIG:/config ghcr.io/dominik-robert/kubectl-over-env:v1.22.2 apply -f Deployment.yaml

# Examples
Here are some examples of the usage of the container
## Github Actions
You can easily push to your local kubernetes cluster inside github actions with this snippet

```yaml
- name: Deploy Kubernetes
  id: deploy
  uses: docker://ghcr.io/dominik-robert/kubectl-over-env:v1.22.2
  with:
    args: '"apply -f Deployment.yaml"'
  env:
    CERTIFICATE_AUTHORITY_DATA: ${{ secrets.CERTIFICATE_AUTHORITY_DATA }}
    CLIENT_CERTIFICATE_DATA: ${{ secrets.CLIENT_CERTIFICATE_DATA }}
    CLIENT_KEY_DATA: ${{ secrets.CLIENT_KEY_DATA }}
    SERVER: ${{ secrets.SERVER }}
```