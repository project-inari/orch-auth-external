# orch-auth-external

GO codebase repository for Inari backend system microservice `orch-auth-external`.

[Kustomize Repo](https://github.com/project-inari/orch-auth-external-cd-config)<br>
[ArgoCD](https://argocd-alpha.inari-th.com/applications/argocd/orch-auth-external?view=tree&resource=)<br>
[Confluence Docs](https://inari.atlassian.net/wiki/spaces/KB/pages/60784641/orch-auth-external)<br>

## Development
Prepare the development environment:
```bash
make prepare
```

Run in local development environment:
```bash
make dev.up
```

Run in alpha development environment:
```bash
make dev.up.alpha
```
